package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Client struct {
	baseURL    string
	httpClient *http.Client
	retries    int
}

func NewClient(
	baseURL string,
	timeout time.Duration,
	retries int,
) *Client {
	return &Client{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: timeout,
		},
		retries: retries,
	}
}

func (c *Client) Post(
	ctx *gin.Context,
	path string,
	body any,
	respBody any,
) error {
	return c.SendRequest(
		ctx,
		http.MethodPost,
		path,
		body,
		respBody,
	)
}

func (c *Client) SendRequest(
	ctx *gin.Context,
	method string,
	path string,
	body any,
	respBody any,
) error {
	var bodyReader io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return err
		}
		bodyReader = bytes.NewBuffer(b)
	}

	url := c.baseURL + path

	var lastErr error
	for attempt := 0; attempt <= c.retries; attempt++ {
		req, err := http.NewRequestWithContext(
			context.Background(),
			method,
			url,
			bodyReader,
		)
		if err != nil {
			return err
		}

		req.Header.Set("Content-Type", "application/json")

		if ctx != nil {
			if requestID := ctx.GetHeader("X-Request-ID"); requestID != "" {
				req.Header.Set("X-Request-ID", requestID)
			}
			if auth := ctx.GetHeader("Authorization"); auth != "" {
				req.Header.Set("Authorization", auth)
			}
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			lastErr = err
			time.Sleep(time.Duration(attempt+1) * 100 * time.Millisecond)
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode >= 500 {
			lastErr = errors.New("model service error: " + resp.Status)
			time.Sleep(time.Duration(attempt+1) * 100 * time.Millisecond)
			continue
		}

		if respBody != nil {
			return json.NewDecoder(resp.Body).Decode(respBody)
		}
		return nil
	}

	return lastErr
}
