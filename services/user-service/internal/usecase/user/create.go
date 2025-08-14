package user_usecase

import (
	"context"
	"log"

	"github.com/adhyttungga/skeleton-application/services/user-service/internal/models"
)

func (usecase *UserUsecaseImpl) Create(c context.Context, user *models.User) error {
	// Validate the user data
	if err := usecase.Validate.Struct(user); err != nil {
		log.Println("Validation error:", err)
		return err
	}

	// Call the repository to create the user
	if err := usecase.Repository.Create(c, user); err != nil {
		return err
	}

	return nil
}
