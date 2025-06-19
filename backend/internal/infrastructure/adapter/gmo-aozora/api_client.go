package adapter

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	adapter "github.com/test-tzs/nomraeite/internal/domain/adapter/gmo-aozora"
	model "github.com/test-tzs/nomraeite/internal/domain/model/api/gmo-aozora"
	transferStatusModel "github.com/test-tzs/nomraeite/internal/domain/model/api/gmo-aozora/transfer_status"
	dto "github.com/test-tzs/nomraeite/internal/infrastructure/persistence/gmo-aozora/dto"
	transferStatusDto "github.com/test-tzs/nomraeite/internal/infrastructure/persistence/gmo-aozora/transfer_status/dto"
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

func (c *ApiClientAdapter) GetAuthURL(state string) (string, error) {
	authUrl := c.client.BuildURL(c.apiUrl, "/auth/v1/authorization")
	u, err := url.Parse(authUrl)

	if err != nil {
		return "", err
	}

	q := u.Query()
	q.Add("response_type", "code")
	q.Add("client_id", c.config.GmoAozoraNetBankClientID)
	q.Add("scope", strings.ReplaceAll(c.config.GmoAozoraNetBankAuthScope, ",", " "))
	q.Add("state", state)
	q.Add("redirect_uri", c.config.GmoAozoraNetBankAuthCallbackURL)
	u.RawQuery = q.Encode()

	return u.String(), nil
}

func (c *ApiClientAdapter) Connect(ctx context.Context, code string) (*model.GmoAozoraTokenClaimResult, error) {
	formData := url.Values{}
	formData.Set("grant_type", "authorization_code")
	formData.Set("code", code)
	formData.Set("redirect_uri", c.config.GmoAozoraNetBankAuthCallbackURL)

	response, err := c.requestToken(ctx, formData)

	if err != nil {
		return nil, err
	}

	c.logger.Info("Successfully fetched new access token using authorization code", map[string]any{
		"token_type": response.TokenType,
		"expires_in": response.ExpiresIn,
	})

	return response, nil
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
	fmt.Printf("Transfer Request DTO: %+v\n", dtoRequest)
	requestJSON, err := json.Marshal(dtoRequest)
	if err != nil {
		c.logger.Error("Failed to marshal transfer request", map[string]any{
			"error": err.Error(),
		})
		return model.NewSystemErrorTransferResponse(err.Error())
	}

	return c.requestTransfer(ctx, bytes.NewBuffer(requestJSON), option)
}

func (c *ApiClientAdapter) RequestBulkTransfer(ctx context.Context, header model.BulkTransferHeaderRequest, request model.BulkTransferParamsRequest) *model.TransferResponse {
	dtoHeader := dto.ToDTOBulkTransferHeaderRequest(header)
	option := &httpApp.Option{
		Header: http.Header{
			"x-access-token":  {dtoHeader.AccessToken},
			"Idempotency-Key": {dtoHeader.IdempotencyKey},
		},
	}

	dtoRequest := dto.ToDTOBulkTransferParamsRequest(request)
	requestJSON, err := json.Marshal(dtoRequest)
	if err != nil {
		c.logger.Error("Failed to marshal bulk transfer request", map[string]any{
			"error": err.Error(),
		})
		return model.NewSystemErrorTransferResponse(err.Error())
	}

	return c.requestBulkTransfer(ctx, bytes.NewBuffer(requestJSON), option)
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

func (c *ApiClientAdapter) requestBulkTransfer(ctx context.Context, body io.Reader, option *httpApp.Option) *model.TransferResponse {
	url := c.client.BuildURL(c.apiUrl, "/corporation/v1/bulktransfer/request")
	response, err := c.client.Post(ctx, url, body, option)
	if err != nil {
		c.logger.Error("Failed to make bulk transfer request", map[string]any{
			"error": err.Error(),
		})
		return model.NewSystemErrorTransferResponse(err.Error())
	}

	defer func() {
		err = response.Body.Close()
		if err != nil {
			c.logger.Error("Failed to close bulk transfer response body", map[string]any{
				"error": err.Error(),
			})
		}
	}()

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		c.logger.Error("Error reading bulk transfer response body", map[string]any{
			"error": err.Error(),
		})
		return model.NewSystemErrorTransferResponse(err.Error())
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

func (c *ApiClientAdapter) GetTransferStatus(ctx context.Context, header model.AuthHeader, request transferStatusModel.TransferStatusRequest) (*transferStatusModel.TransferStatusResponse, error) {
	requestDTO := transferStatusDto.ToTransferStatusRequestDTO(request)
	option := &httpApp.Option{
		Header: http.Header{
			"x-access-token": {header.AccessToken},
		},
	}

	queryParams := requestDTO.ToQueryParams()
	fmt.Printf("Query Params: %v\n", queryParams)
	return c.requestTransferStatus(ctx, queryParams, option)
}

func (c *ApiClientAdapter) GetBulkTransferStatus(ctx context.Context, header model.AuthHeader, request transferStatusModel.BulkTransferStatusRequest) (*transferStatusModel.BulkTransferStatusResponse, error) {
	requestDTO := transferStatusDto.ToBulkTransferStatusRequestDTO(request)
	option := &httpApp.Option{
		Header: http.Header{
			"x-access-token": {header.AccessToken},
		},
	}

	queryParams := requestDTO.ToQueryParams()
	return c.requestBulkTransferStatus(ctx, queryParams, option)
}

func (c *ApiClientAdapter) requestTransferStatus(ctx context.Context, queryParams url.Values, option *httpApp.Option) (*transferStatusModel.TransferStatusResponse, error) {
	url := c.client.BuildURL(c.apiUrl, "/corporation/v1/transfer/status")
	response, err := c.client.Get(ctx, url, queryParams, option)
	if err != nil {
		c.logger.Error("Failed to get transfer status", map[string]any{
			"error": err.Error(),
		})
		return nil, fmt.Errorf("failed to get transfer status: %w", err)
	}
	defer c.closeResponseBody(response.Body)

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		c.logger.Error("Error reading response body", map[string]any{
			"error": err.Error(),
		})
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	if response.StatusCode != http.StatusOK {
		return nil, c.handleErrorResponse(response.StatusCode, bodyBytes)
	}

	var responseDTO transferStatusDto.TransferStatusResponseDTO
	err = c.unmarshalResponse(bodyBytes, &responseDTO)
	if err != nil {
		return nil, fmt.Errorf("failed to parse transfer status response: %w", err)
	}

	return responseDTO.ToModel(), nil
}

func (c *ApiClientAdapter) requestBulkTransferStatus(ctx context.Context, queryParams url.Values, option *httpApp.Option) (*transferStatusModel.BulkTransferStatusResponse, error) {
	url := c.client.BuildURL(c.apiUrl, "/corporation/v1/bulktransfer/status")
	response, err := c.client.Get(ctx, url, queryParams, option)
	if err != nil {
		c.logger.Error("Failed to get bulk transfer status", map[string]any{
			"error": err.Error(),
		})
		return nil, fmt.Errorf("failed to get bulk transfer status: %w", err)
	}
	defer c.closeResponseBody(response.Body)

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		c.logger.Error("Error reading response body", map[string]any{
			"error": err.Error(),
		})
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	if response.StatusCode != http.StatusOK {
		return nil, c.handleErrorResponse(response.StatusCode, bodyBytes)
	}

	var responseDTO transferStatusDto.BulkTransferStatusResponseDTO
	err = c.unmarshalResponse(bodyBytes, &responseDTO)
	if err != nil {
		return nil, fmt.Errorf("failed to parse bulk transfer status response: %w", err)
	}

	return responseDTO.ToModel(), nil
}

func (c *ApiClientAdapter) unmarshalResponse(body []byte, target any) error {
	if err := json.Unmarshal(body, target); err != nil {
		c.logger.Error("Failed to unmarshal response", map[string]any{
			"error": err.Error(),
			"body":  string(body),
		})
		return err
	}
	return nil
}

func (c *ApiClientAdapter) closeResponseBody(body io.ReadCloser) {
	if err := body.Close(); err != nil {
		c.logger.Error("Failed to close response body", map[string]any{
			"error": err.Error(),
		})
	}
}
