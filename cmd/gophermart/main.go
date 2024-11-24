package main

import (
	"context"
	"crypto/rand"
	orderrepo "github.com/aridae/gophermart-diploma/internal/repos/order-repo"
	userbalancerepo "github.com/aridae/gophermart-diploma/internal/repos/user-balance-repo"
	userrepo "github.com/aridae/gophermart-diploma/internal/repos/user-repo"
	withdrawallogrepo "github.com/aridae/gophermart-diploma/internal/repos/withdrawal-log-repo"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/aridae/gophermart-diploma/internal/auth/authmw"
	"github.com/aridae/gophermart-diploma/internal/config"
	"github.com/aridae/gophermart-diploma/internal/database"
	"github.com/aridae/gophermart-diploma/internal/downstream/accrual"
	"github.com/aridae/gophermart-diploma/internal/jwt"
	"github.com/aridae/gophermart-diploma/internal/logger"
	orderaccrualsync "github.com/aridae/gophermart-diploma/internal/order-accrual-sync"
	httpserver "github.com/aridae/gophermart-diploma/internal/transport/http"
	httpapi "github.com/aridae/gophermart-diploma/internal/transport/http/http-api"
	oapispec "github.com/aridae/gophermart-diploma/internal/transport/http/http-api/oapi-spec"
	getbalance "github.com/aridae/gophermart-diploma/internal/usecases/get-balance"
	getorders "github.com/aridae/gophermart-diploma/internal/usecases/get-orders"
	getwithdrawalshistory "github.com/aridae/gophermart-diploma/internal/usecases/get-withdrawals-history"
	loginuser "github.com/aridae/gophermart-diploma/internal/usecases/login-user"
	registeruser "github.com/aridae/gophermart-diploma/internal/usecases/register-user"
	requestwithdrawal "github.com/aridae/gophermart-diploma/internal/usecases/request-withdrawal"
	submitorder "github.com/aridae/gophermart-diploma/internal/usecases/submit-order"
	"github.com/aridae/gophermart-diploma/pkg/postgres"
	trmpgx "github.com/avito-tech/go-transaction-manager/drivers/pgxv5/v2"
	trman "github.com/avito-tech/go-transaction-manager/trm/v2/manager"
	"github.com/flowchartsman/swaggerui"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	oapimw "github.com/oapi-codegen/nethttp-middleware"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go watchTerminationSignals(cancel, os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGHUP)

	cnf := config.Obtain()

	pgClient := mustInitPostgresClient(ctx, cnf)

	err := database.PrepareSchema(ctx, pgClient)
	if err != nil {
		logger.Fatalf("failed to prepare database schema: %v", err)
	}

	pgTxManager := trman.Must(trmpgx.NewDefaultFactory(pgClient))

	userRepository := userrepo.New(pgClient, trmpgx.DefaultCtxGetter)

	userBalanceRepository := userbalancerepo.New(pgClient, trmpgx.DefaultCtxGetter)

	ordersRepository := orderrepo.New(pgClient, trmpgx.DefaultCtxGetter)

	withdrawalsLogsRepository := withdrawallogrepo.New(pgClient, trmpgx.DefaultCtxGetter)

	jwtService := mustInitJWTService(ctx, cnf)

	getBalanceHandler := getbalance.NewHandler(userBalanceRepository)
	getOrdersHandler := getorders.NewHandler(ordersRepository)
	getWithdrawalsHistoryHandler := getwithdrawalshistory.NewHandler(withdrawalsLogsRepository)
	loginUserHandler := loginuser.NewHandler(userRepository, jwtService)
	registerUserHandler := registeruser.NewHandler(userRepository, jwtService)
	requestWithdrawalHandler := requestwithdrawal.NewHandler(pgTxManager, ordersRepository, withdrawalsLogsRepository)
	submitOrderHandler := submitorder.NewHandler(pgTxManager, ordersRepository)

	orderAccrualService := accrual.NewClient(cnf.AccuralSystemAddress)
	orderAccrualSyncer := orderaccrualsync.New(orderAccrualService, ordersRepository, cnf.AccrualSyncInterval, cnf.AccrualSyncWorkersPoolSize)
	go orderAccrualSyncer.Run(ctx)

	apiService := httpapi.NewAPIService(
		getBalanceHandler,
		getOrdersHandler,
		getWithdrawalsHistoryHandler,
		loginUserHandler,
		registerUserHandler,
		requestWithdrawalHandler,
		submitOrderHandler,
	)

	swagger, err := oapispec.GetSwagger()
	if err != nil {
		logger.Fatalf("failed to obtain swagger: %v", err)
	}

	httpPublicServer := httpserver.NewServer(
		cnf.Address,
		oapispec.HandlerFromMuxWithBaseURL(apiService, http.NewServeMux(), "/api"),
		oapimw.OapiRequestValidatorWithOptions(swagger, &oapimw.Options{
			Options: openapi3filter.Options{
				AuthenticationFunc: openapi3filter.NoopAuthenticationFunc,
			},
		}),
		authmw.AuthenticateMiddleware(jwtService, []string{
			`/api/user/login`,
			`/api/user/register`,
		}),
	)

	httpAdminServer := mustInitAdminServer(ctx, swagger, cnf)
	go func() {
		if err = httpAdminServer.Run(ctx); err != nil {
			logger.Fatalf("failed to run http admin server: %v", err)
		}
	}()

	if err = httpPublicServer.Run(ctx); err != nil {
		logger.Fatalf("failed to run http public server: %v", err)
	}
}

func mustInitPostgresClient(ctx context.Context, cnf *config.Config) *postgres.Client {
	client, err := postgres.NewClient(ctx, cnf.DatabaseDsn,
		postgres.WithInitialReconnectBackoffOnFail(time.Second),
	)
	if err != nil {
		logger.Fatalf("failed to init postgres client: %v", err)
	}

	return client
}

func mustInitJWTService(_ context.Context, cnf *config.Config) *jwt.Service {
	key := cnf.JWTKey
	if key == "" {
		randomFixedLenKey := make([]byte, 64)

		_, err := rand.Read(randomFixedLenKey)
		if err != nil {
			logger.Fatalf("failed to generate JWT key: %v", err)
		}

		key = string(randomFixedLenKey)
	}

	keyProvider := func(ctx context.Context) []byte {
		return []byte(key)
	}

	return jwt.NewService(keyProvider)
}

func mustInitAdminServer(_ context.Context, swagger *openapi3.T, cnf *config.Config) *httpserver.Server {
	swaggerRawSpec, err := swagger.MarshalJSON()
	if err != nil {
		logger.Fatalf("failed to obtain swagger bytes spec: %v", err)
	}

	adminRouter := http.NewServeMux()
	adminRouter.Handle("/swagger/", http.StripPrefix("/swagger", swaggerui.Handler(swaggerRawSpec)))

	adminServer := httpserver.NewServer(
		cnf.AdminAddress,
		adminRouter,
	)

	return adminServer
}

func watchTerminationSignals(cancel func(), signals ...os.Signal) {
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, signals...)

	<-signalCh

	logger.Infof("Got signal, shutting down...")

	// If you fail to cancel the context, the goroutine that WithCancel or WithTimeout created
	// will be retained in memory indefinitely (until the program shuts down), causing a memory leak.
	cancel()
}
