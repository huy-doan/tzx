package model

import (
	"time"

	"github.com/test-tzs/nomraeite/internal/domain/model/user"
	util "github.com/test-tzs/nomraeite/internal/domain/object/basedatetime"
)

type TwoFactorToken struct {
	ID        int
	UserID    int
	Token     string
	MFAType   int
	User      *user.User
	IsUsed    bool
	ExpiredAt time.Time

	util.BaseColumnTimestamp
}

func (t *TwoFactorToken) MarkAsUsed() {
	t.IsUsed = true
}
