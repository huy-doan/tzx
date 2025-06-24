package http

import (
	"context"
	"fmt"
	"io"
	"maps"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/test-tzs/nomraeite/internal/pkg/config"
	"github.com/test-tzs/nomraeite/internal/pkg/logger"
)

const (
	defaultTimeoutSeconds = 30
)

type Option struct {
	Header         http.Header
	TimeoutSeconds int
	Host           string
	BasicAuth      *BasicAuth
}

type BasicAuth struct {
	Username string
	Password string
}

type Client interface {
	Get(ctx context.Context, path string, params url.Values, option *Option) (*http.Response, error)
	Post(ctx context.Context, path string, reqBody io.Reader, option *Option) (*http.Response, error)
	PostForm(ctx context.Context, path string, formData url.Values, option *Option) (*http.Response, error)
	BuildURL(baseURL, path string) string
}

type ClientImpl struct {
	config *config.Config
	logger logger.Logger
}

func NewClient(logger logger.Logger) Client {
	cfg := config.GetConfig()
	return &ClientImpl{
		config: cfg,
		logger: logger,
	}
}

func (c *ClientImpl) Get(ctx context.Context, url string, params url.Values, option *Option) (*http.Response, error) {
	header := http.Header{}
	if option != nil {
		maps.Copy(header, option.Header)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	if params != nil {
		req.URL.RawQuery = params.Encode()
	}

	return c.doRequest(ctx, req, header, option)
}

func (c *ClientImpl) Post(ctx context.Context, url string, reqBody io.Reader, option *Option) (*http.Response, error) {
	header := http.Header{
		"Content-Type": {"application/json"},
	}
	if option != nil {
		maps.Copy(header, option.Header)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, reqBody)
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, req, header, option)
}

func (c *ClientImpl) PostForm(ctx context.Context, url string, formData url.Values, option *Option) (*http.Response, error) {
	header := http.Header{
		"Content-Type": {"application/x-www-form-urlencoded"},
		"Accept":       {"application/json;charset=UTF-8"},
	}
	if option != nil {
		maps.Copy(header, option.Header)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, strings.NewReader(formData.Encode()))
	if err != nil {
		return nil, err
	}

	return c.doRequest(ctx, req, header, option)
}

func (c *ClientImpl) BuildURL(baseURL, path string) string {
	if path == "" {
		return baseURL
	}

	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	baseURL = strings.TrimSuffix(baseURL, "/")

	return baseURL + path
}

func (c *ClientImpl) doRequest(ctx context.Context, req *http.Request, header http.Header, option *Option) (response *http.Response, err error) {
	if len(header) > 0 {
		req.Header = header
	}

	op := &Option{}
	if option != nil {
		op = option
	}
	if op.BasicAuth != nil {
		req.SetBasicAuth(op.BasicAuth.Username, op.BasicAuth.Password)
	}
	if op.Host != "" {
		req.Host = op.Host
	}

	timeout := defaultTimeoutSeconds
	if op != nil && op.TimeoutSeconds > 0 {
		timeout = op.TimeoutSeconds
	}

	client := &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}

	response, err = client.Do(req)
	if err != nil {
		c.logger.Error("HTTP request failed", map[string]any{
			"error":  err.Error(),
			"method": req.Method,
			"url":    req.URL.String(),
		})
		return nil, fmt.Errorf("HTTP request failed: %w", err)
	}

	return
}
