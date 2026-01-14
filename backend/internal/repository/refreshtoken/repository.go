package refreshtoken

import (
	"context"
	"time"

	"github.com/Asuzaka/chat-app/backend/internal/domain/refreshtoken"
)

type RefreshTokenRepository interface {
	Save(
		ctx context.Context,
		userID string,
		token string,
		expiresAt time.Time,
	) error

	Find(
		ctx context.Context,
		token string,
	) (*refreshtoken.RefreshToken, error)

	Delete(
		ctx context.Context,
		token string,
	) error

	DeleteByUser(
		ctx context.Context,
		userID string,
	) error
}
