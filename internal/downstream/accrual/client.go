package accrual

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/hashicorp/go-retryablehttp"
)

type Client struct {
	client  *http.Client
	address string
}

func NewClient(address string) *Client {
	rt := http.DefaultTransport

	retryableHTTPClient := retryablehttp.NewClient()
	retryableHTTPClient.HTTPClient.Transport = rt

	return &Client{
		client:  retryableHTTPClient.StandardClient(),
		address: address,
	}
}

func (c *Client) doRequest(
	ctx context.Context,
	method string,
	url string,
	body io.Reader,
	presentError func(status string, code int) error,
) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, method, c.address+url, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create http request: %w", err)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to do http request: %w", err)
	}
	defer resp.Body.Close()

	err = presentError(resp.Status, resp.StatusCode)
	if err != nil {
		return nil, fmt.Errorf("http call responded with err: %w", err)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return bytes, nil
}
