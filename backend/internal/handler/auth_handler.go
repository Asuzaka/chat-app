package handler

import (
	"github.com/Asuzaka/chat-app/backend/internal/auth"
	"github.com/Asuzaka/chat-app/backend/internal/repository/refreshtoken"
	"github.com/Asuzaka/chat-app/backend/internal/repository/user"
)

type AuthService struct {
	users      user.UserRepository
	tokens     refreshtoken.RefreshTokenRepository
	jwtManager *auth.JWTManager
}
