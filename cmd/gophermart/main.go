package main

import (
	"context"
	"crypto/rand"
	"github.com/aridae/gophermart-diploma/internal/auth/authmw"
	"github.com/aridae/gophermart-diploma/internal/config"
	"github.com/aridae/gophermart-diploma/internal/jwt"
	"github.com/aridae/gophermart-diploma/internal/logger"
	userdb "github.com/aridae/gophermart-diploma/internal/repo/user-db"
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
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		signalCh := make(chan os.Signal, 1)
		signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGHUP)

		<-signalCh

		logger.Infof("Got signal, shutting down...")

		// If you fail to cancel the context, the goroutine that WithCancel or WithTimeout created
		// will be retained in memory indefinitely (until the program shuts down), causing a memory leak.
		cancel()
	}()

	cnf := config.Obtain()

	pgClient := mustInitPostgresClient(ctx, cnf)
	pgTxManager := trman.Must(trmpgx.NewDefaultFactory(pgClient))
	_ = pgTxManager // TODO add handler-level transactions

	userRepository, err := userdb.NewRepo(ctx, pgClient, trmpgx.DefaultCtxGetter)
	if err != nil {
		logger.Fatalf("failed to initialize user repository: %v", err)
	}

	jwtService := mustInitJWTService(ctx, cnf)

	getBalanceHandler := getbalance.NewHandler()
	getOrdersHandler := getorders.NewHandler()
	getWithdrawalsHistoryHandler := getwithdrawalshistory.NewHandler()
	loginUserHandler := loginuser.NewHandler(userRepository, jwtService)
	registerUserHandler := registeruser.NewHandler(userRepository, jwtService)
	requestWithdrawalHandler := requestwithdrawal.NewHandler()
	submitOrderHandler := submitorder.NewHandler()

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
		oapispec.HandlerFromMux(apiService, http.NewServeMux()),
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
