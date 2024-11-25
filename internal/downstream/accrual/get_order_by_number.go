package accrual

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) GetOrderByNumber(ctx context.Context, orderNumber string) (Order, error) {
	url := fmt.Sprintf("/api/orders/%s", orderNumber)

	respBytes, err := c.doRequest(ctx, http.MethodGet, url, nil, getOrderCallStatusCodeToError)
	if err != nil {
		return Order{}, fmt.Errorf("doRequest: %w", err)
	}

	dto := orderDTO{}
	err = json.Unmarshal(respBytes, &dto)
	if err != nil {
		return Order{}, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	order := parseOrderDTO(dto)

	return order, err
}

func getOrderCallStatusCodeToError(status string, code int) error {
	switch code {
	case http.StatusNoContent:
		return ErrOrderNotFound
	case http.StatusTooManyRequests:
		return ErrResourceExhausted
	case http.StatusOK:
		return nil
	}

	return fmt.Errorf("got status <code:%d>: %s", code, status)
}
