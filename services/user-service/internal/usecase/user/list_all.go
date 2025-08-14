package user_usecase

import (
	"context"

	"github.com/adhyttungga/skeleton-application/services/user-service/internal/models"
)

func (usecase *UserUsecaseImpl) ListAll(c context.Context) (*[]models.User, error) {
	// Call the repository to list all users
	users, err := usecase.Repository.ListAll(c)
	if err != nil {
		return nil, err
	}

	return &users, nil
}
