package user_usecase

import (
	"context"

	"github.com/adhyttungga/skeleton-application/services/user-service/internal/models"
)

type UserUsecase interface {
	Create(c context.Context, user *models.User) (string, error)
	Fetch(c context.Context, id string) (*models.User, error)
	ListAll(c context.Context) (*[]models.User, error)
	Update(c context.Context, id string, user *models.User) error
	Delete(c context.Context, id string) error
}
