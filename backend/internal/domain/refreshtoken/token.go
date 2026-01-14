package refreshtoken

import "time"

type RefreshToken struct {
	Token     string
	UserID    string
	ExpiresAt time.Time
	CreatedAt time.Time
}
