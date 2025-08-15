package auth_usecase

import (
	"context"

	"github.com/adhyttungga/skeleton-application/services/user-service/internal/models"
)

type AuthUsecase interface {
	SignIn(c context.Context, req models.User) (string, error)
}
