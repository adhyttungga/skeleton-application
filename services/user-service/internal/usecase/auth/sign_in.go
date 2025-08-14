package auth_usecase

import (
	"context"
	"log"

	"github.com/adhyttungga/skeleton-application/services/user-service/internal/models"
)

func (uc *AuthUsecaseImpl) SignIn(c context.Context, user *models.User) (string, error) {
	// Validate user data
	if err := uc.Validate.Struct(user); err != nil {
		log.Println("Validation error: ", err)
		return "", err
	}

	// Retrieve user data by email
	// Validate password
	// JWT token
	// Set token to cookies
	return "", nil
}
