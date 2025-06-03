package model

import (
	"time"

	model "github.com/test-tzs/nomraeite/internal/domain/model/api/gmo-aozora"
	util "github.com/test-tzs/nomraeite/internal/domain/object/basedatetime"
	object "github.com/test-tzs/nomraeite/internal/domain/object/connected_service_token"
)

// ConnectedServiceToken represents an external service connection
type ConnectedServiceToken struct {
	ID int
	util.BaseColumnTimestamp

	ServiceName          object.ConnectedServiceName
	UserID               int
	RefreshToken         string
	AccessToken          string
	AccessTokenExpiredAt *time.Time
}

// IsAccessTokenExpired checks if the access token is expired
func (s *ConnectedServiceToken) IsAccessTokenExpired() bool {
	if s.AccessTokenExpiredAt == nil {
		return true
	}
	return s.AccessTokenExpiredAt.Before(time.Now())
}

// IsAccessTokenValid checks if the access token is valid
func (s *ConnectedServiceToken) IsAccessTokenValid() bool {
	return s.AccessToken != "" && !s.IsAccessTokenExpired()
}

func (s *ConnectedServiceToken) SetGmoAozoraToken(response *model.GmoAozoraTokenClaimResult) {
	s.RefreshToken = response.RefreshToken
	s.AccessToken = response.AccessToken
	accessTokenExpiredAt := time.Now().Add(time.Second * time.Duration(response.ExpiresIn))
	s.AccessTokenExpiredAt = &accessTokenExpiredAt
}

// ClearAccessToken clears the access token and its expiry time
func (s *ConnectedServiceToken) ClearAccessToken() {
	s.AccessToken = ""
	s.AccessTokenExpiredAt = nil
}
