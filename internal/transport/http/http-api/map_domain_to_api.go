package httpapi

import (
	"fmt"

	"github.com/aridae/gophermart-diploma/internal/model"
	oapispec "github.com/aridae/gophermart-diploma/internal/transport/http/http-api/oapi-spec"
	"github.com/aridae/gophermart-diploma/pkg/pointer"
)

func mapDomainToAPIWithdrawal(withdrawal model.WithdrawalLog) oapispec.Withdrawal {
	return oapispec.Withdrawal{
		Order:       withdrawal.OrderNumber,
		ProcessedAt: withdrawal.CreatedAt,
		Sum:         withdrawal.Sum.Float32(),
	}
}

var dom2apiOrderStatus = map[model.OrderStatus]oapispec.OrderStatus{
	model.OrderStatusInvalid:    oapispec.INVALID,
	model.OrderStatusNew:        oapispec.NEW,
	model.OrderStatusProcessed:  oapispec.PROCESSED,
	model.OrderStatusProcessing: oapispec.PROCESSING,
}

func mapDomainToAPIOrderStatus(status model.OrderStatus) (oapispec.OrderStatus, error) {
	apiStatus, ok := dom2apiOrderStatus[status]
	if ok {
		return apiStatus, nil
	}

	return "", fmt.Errorf("unknown order status '%s'", status)
}

func mapDomainToAPIOrder(order model.Order) (oapispec.Order, error) {
	apiStatus, err := mapDomainToAPIOrderStatus(order.Status)
	if err != nil {
		return oapispec.Order{}, fmt.Errorf("error converting order status '%s' to API model: %w", order.Status, err)
	}

	var accrual *float32
	if order.Accrual != nil {
		accrual = pointer.To(order.Accrual.Float32())
	}

	return oapispec.Order{
		Accrual:    accrual,
		Number:     order.Number,
		Status:     apiStatus,
		UploadedAt: order.UploadedAt,
	}, nil
}

func mapDomainToAPIBalance(balance model.Balance) oapispec.Balance {
	return oapispec.Balance{
		Current:   balance.Current.Float32(),
		Withdrawn: balance.Withdrawn.Float32(),
	}
}
