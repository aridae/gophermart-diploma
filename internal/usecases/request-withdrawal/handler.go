package requestwithdrawal

import (
	"context"
	"fmt"
	"time"

	"github.com/aridae/gophermart-diploma/internal/auth/authctx"
	"github.com/aridae/gophermart-diploma/internal/model"
	domainerrors "github.com/aridae/gophermart-diploma/internal/model/domain-errors"
	orderrepo "github.com/aridae/gophermart-diploma/internal/repos/order-repo"
	"github.com/aridae/gophermart-diploma/pkg/pointer"
)

type transactionManager interface {
	Do(ctx context.Context, fn func(ctx context.Context) error) (err error)
}

type ordersRepository interface {
	GetByNumbers(ctx context.Context, orderNumbers []string) ([]model.Order, error)
	UpdateOrder(ctx context.Context, orderNumber string, setters ...orderrepo.Setter) error
}

type withdrawalLogsRepository interface {
	CreateWithdrawalLog(ctx context.Context, withdrawal model.WithdrawalLog, now time.Time) error
}

type Handler struct {
	transactionManager       transactionManager
	ordersRepository         ordersRepository
	withdrawalLogsRepository withdrawalLogsRepository
	now                      func() time.Time
}

func NewHandler(
	transactionManager transactionManager,
	ordersRepository ordersRepository,
	withdrawalLogsRepository withdrawalLogsRepository,
) *Handler {
	return &Handler{
		transactionManager:       transactionManager,
		ordersRepository:         ordersRepository,
		withdrawalLogsRepository: withdrawalLogsRepository,
		now:                      time.Now().UTC,
	}
}

type Request struct {
	OrderNumber string
	Sum         model.Money
}

func (h *Handler) Handle(ctx context.Context, req Request) error {
	user, authorized := authctx.GetUserFromContext(ctx)
	if !authorized {
		return domainerrors.UnauthorizedError()
	}

	now := h.now()

	err := h.transactionManager.Do(ctx, func(ctx context.Context) error {
		orders, txErr := h.ordersRepository.GetByNumbers(ctx, []string{req.OrderNumber})
		if txErr != nil {
			return fmt.Errorf("ordersRepository.GetByNumbers <number:%s>: %w", req.OrderNumber, txErr)
		}

		if len(orders) == 0 {
			return domainerrors.OrderNotFoundError(req.OrderNumber)
		}
		order := orders[0]

		if order.Owner.Login != user.Login {
			return domainerrors.NoAccessToOrderError(req.OrderNumber)
		}

		orderAccrual := pointer.Deref(order.Accrual, model.NewMoney(0))
		if orderAccrual.Less(req.Sum) {
			return domainerrors.InsufficientOrderAccrualError(req.OrderNumber, orderAccrual, req.Sum)
		}
		withdrawnAccrual := orderAccrual.Sub(req.Sum)

		txErr = h.ordersRepository.UpdateOrder(ctx, req.OrderNumber, orderrepo.SetOrderAccrual(withdrawnAccrual))
		if txErr != nil {
			return fmt.Errorf("ordersRepository.UpdateOrder: %w", txErr)
		}

		txErr = h.withdrawalLogsRepository.CreateWithdrawalLog(ctx, model.WithdrawalLog{
			Sum:         req.Sum,
			OrderNumber: req.OrderNumber,
			CreatedAt:   now,
			Actor:       user,
		}, now)
		if txErr != nil {
			return fmt.Errorf("withdrawalLogsRepository.CreateWithdrawalLog: %w", txErr)
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("transaction failed with err: %w", err)
	}

	return nil
}
