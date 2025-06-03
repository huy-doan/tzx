package adapter

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	adapter "github.com/test-tzs/nomraeite/internal/domain/adapter/gmo-aozora"
	model "github.com/test-tzs/nomraeite/internal/domain/model/api/gmo-aozora"
	dto "github.com/test-tzs/nomraeite/internal/infrastructure/persistence/gmo-aozora/dto"
	config "github.com/test-tzs/nomraeite/internal/pkg/config"
	httpApp "github.com/test-tzs/nomraeite/internal/pkg/http"
	"github.com/test-tzs/nomraeite/internal/pkg/logger"
)

type ApiClientAdapter struct {
	client          httpApp.Client
	logger          logger.Logger
	config          *config.Config
	apiUrl          string
	responseHandler *TransferResponseHandler
}

func NewApiClient(logger logger.Logger) adapter.ApiClient {
	appConfig := config.GetConfig()
	if appConfig.GmoAozoraNetBankAPIEndPoint == "" {
		logger.Error("GmoAozoraNetBankAPIEndPoint is not set in the configuration", nil)
		return nil
	}
	return &ApiClientAdapter{
		client:          httpApp.NewClient(logger),
		logger:          logger,
		config:          appConfig,
		apiUrl:          appConfig.GmoAozoraNetBankAPIEndPoint,
		responseHandler: NewTransferResponseHandler(logger),
	}
}

func (c *ApiClientAdapter) RefreshToken(ctx context.Context, refreshToken string) (*model.GmoAozoraTokenClaimResult, error) {
	formData := url.Values{}
	formData.Set("grant_type", "refresh_token")
	formData.Set("refresh_token", refreshToken)
	response, err := c.requestToken(ctx, formData)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *ApiClientAdapter) RequestTransfer(ctx context.Context, header model.TransferHeaderRequest, request model.TransferParamsRequest) *model.TransferResponse {
	dtoHeader := dto.ToDTOTransferHeaderRequest(header)
	option := &httpApp.Option{
		Header: http.Header{
			"x-access-token":  {dtoHeader.AccessToken},
			"Idempotency-Key": {dtoHeader.IdempotencyKey},
		},
	}

	dtoRequest := dto.ToDTOTransferParamsRequest(request)
	requestJSON, err := json.Marshal(dtoRequest)
	if err != nil {
		c.logger.Error("Failed to marshal transfer request", map[string]any{
			"error": err.Error(),
		})
		return model.NewSystemErrorTransferResponse("Failed to marshal transfer request")
	}

	return c.requestTransfer(ctx, bytes.NewBuffer(requestJSON), option)
}

func (c *ApiClientAdapter) requestTransfer(ctx context.Context, body io.Reader, option *httpApp.Option) *model.TransferResponse {
	url := c.client.BuildURL(c.apiUrl, "/corporation/v1/transfer/request")
	response, err := c.client.Post(ctx, url, body, option)
	if err != nil {
		c.logger.Error("Failed to make transfer request", map[string]any{
			"error": err.Error(),
		})
		return model.NewSystemErrorTransferResponse("Failed to make transfer request")
	}

	defer func() {
		err = response.Body.Close()
		if err != nil {
			c.logger.Error("Failed to close response body", map[string]any{
				"error": err.Error(),
			})
		}
	}()

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		c.logger.Error("Error reading response body", map[string]any{
			"error": err.Error(),
		})
		return model.NewSystemErrorTransferResponse("Error reading response body")
	}

	return c.responseHandler.HandleResponse(response.StatusCode, bodyBytes)
}

func (c *ApiClientAdapter) requestToken(ctx context.Context, formData url.Values) (*model.GmoAozoraTokenClaimResult, error) {
	option := &httpApp.Option{
		BasicAuth: &httpApp.BasicAuth{
			Username: c.config.GmoAozoraNetBankClientID,
			Password: c.config.GmoAozoraNetBankClientSecret,
		},
	}

	url := c.client.BuildURL(c.apiUrl, "/auth/v1/token")

	resp, err := c.client.PostForm(ctx, url, formData, option)
	if err != nil {
		c.logger.Error("Failed to get token", map[string]any{
			"error": err.Error(),
		})
		return nil, fmt.Errorf("failed to get token: %w", err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			c.logger.Error("Failed to close response body", map[string]any{
				"error": err.Error(),
			})
		}
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.logger.Error("Failed to read response body", map[string]any{
			"error": err.Error(),
		})
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return nil, c.handleErrorResponse(resp.StatusCode, body)
	}

	var response dto.GmoAozoraTokenClaimResponse
	if err := json.Unmarshal(body, &response); err != nil {
		c.logger.Error("Failed to parse token response", map[string]any{
			"error": err.Error(),
			"body":  string(body),
		})
		return nil, fmt.Errorf("failed to parse token response: %w", err)
	}
	return response.ToModel(), nil
}

func (c *ApiClientAdapter) handleErrorResponse(statusCode int, body []byte) error {
	c.logger.Error("API returned error", map[string]any{
		"status_code": statusCode,
		"body":        string(body),
	})

	var errResp dto.CommonErrorResponse
	err := json.Unmarshal(body, &errResp)
	if err == nil && (errResp.ErrorCode != "" || errResp.ErrorMessage != "") {
		return fmt.Errorf("API error [%s]: %s", errResp.ErrorCode, errResp.ErrorMessage)
	}

	return fmt.Errorf("API returned error with status code: %d", statusCode)
}
