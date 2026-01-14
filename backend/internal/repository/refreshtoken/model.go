package refreshtoken

import "time"

type RefreshToken struct {
	ID        string
	UserID    string
	TokenHash string

	Device    string
	IP        string
	UserAgent string

	ExpiresAt time.Time
	CreatedAt time.Time
	RevokedAt *time.Time
}
