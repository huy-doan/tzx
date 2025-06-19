package token

import (
	"context"

	"github.com/test-tzs/nomraeite/internal/domain/model/token"
)

type TokenRepository interface {
	Create(ctx context.Context, token *token.Token) error
	Update(ctx context.Context, token *token.Token) error
	FindByToken(ctx context.Context, token string) (*token.Token, error)
}
