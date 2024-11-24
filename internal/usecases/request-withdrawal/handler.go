package requestwithdrawal

import (
	"context"
	"fmt"
	"time"

	"github.com/aridae/gophermart-diploma/internal/auth/authctx"
	"github.com/aridae/gophermart-diploma/internal/model"
	domainerrors "github.com/aridae/gophermart-diploma/internal/model/domain-errors"
)

type transactionManager interface {
	Do(ctx context.Context, fn func(ctx context.Context) error) (err error)
}

type userBalanceRepository interface {
	GetUserBalance(ctx context.Context, user model.User) (model.Balance, error)
}

type withdrawalLogsRepository interface {
	CreateWithdrawalLog(ctx context.Context, withdrawal model.WithdrawalLog, now time.Time) error
}

type Handler struct {
	transactionManager       transactionManager
	withdrawalLogsRepository withdrawalLogsRepository
	userBalanceRepository    userBalanceRepository
	now                      func() time.Time
}

func NewHandler(
	transactionManager transactionManager,
	withdrawalLogsRepository withdrawalLogsRepository,
	userBalanceRepository userBalanceRepository,
) *Handler {
	return &Handler{
		transactionManager:       transactionManager,
		withdrawalLogsRepository: withdrawalLogsRepository,
		userBalanceRepository:    userBalanceRepository,
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
		return domainerrors.ErrUnauthorized()
	}

	now := h.now()

	err := h.transactionManager.Do(ctx, func(ctx context.Context) error {
		userBalance, txErr := h.userBalanceRepository.GetUserBalance(ctx, user)
		if txErr != nil {
			return fmt.Errorf("userBalanceRepository.GetUserBalance: %w", txErr)
		}

		if userBalance.Current.Less(req.Sum) {
			return domainerrors.ErrInsufficientOrderAccrual(req.OrderNumber, userBalance.Current, req.Sum)
		}

		txErr = h.withdrawalLogsRepository.CreateWithdrawalLog(ctx, model.WithdrawalLog{
			Sum:         req.Sum,
			OrderNumber: req.OrderNumber,
			Actor:       user,
			CreatedAt:   now,
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
