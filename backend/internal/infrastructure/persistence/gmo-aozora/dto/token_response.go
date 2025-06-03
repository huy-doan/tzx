package dto

import (
	model "github.com/test-tzs/nomraeite/internal/domain/model/api/gmo-aozora"
)

type GmoAozoraTokenClaimResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	Scope        string `json:"scope"`
}

func (dto *GmoAozoraTokenClaimResponse) ToModel() *model.GmoAozoraTokenClaimResult {
	return &model.GmoAozoraTokenClaimResult{
		AccessToken:  dto.AccessToken,
		RefreshToken: dto.RefreshToken,
		TokenType:    dto.TokenType,
		ExpiresIn:    dto.ExpiresIn,
		Scope:        dto.Scope,
	}
}
