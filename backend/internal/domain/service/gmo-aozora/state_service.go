package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/test-tzs/nomraeite/internal/pkg/config"
)

type gmoAozoraStateClaims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

type GmoAozoraStateService interface {
	GenerateStateToken(userID int) (string, error)
	TakeUserID(token string) (userID int, err error)
}

type gmoAozoraStateService struct {
	secretKey string
}

func NewGmoAozoraStateService() (GmoAozoraStateService, error) {
	appConfig := config.GetConfig()
	secret := appConfig.JWTSecret
	if secret == "" {
		return nil, errors.New("JWTSecret is not set in the configuration")
	}

	return &gmoAozoraStateService{
		secretKey: secret,
	}, nil
}

func (s *gmoAozoraStateService) GenerateStateToken(userID int) (string, error) {
	claims := gmoAozoraStateClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   fmt.Sprintf("aozora_oauth_state_%d", userID),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.secretKey))
}

func (s *gmoAozoraStateService) TakeUserID(tokenString string) (userID int, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &gmoAozoraStateClaims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.secretKey), nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(*gmoAozoraStateClaims); ok && token.Valid {
		return claims.UserID, nil
	}

	return 0, errors.New("invalid state token")
}
