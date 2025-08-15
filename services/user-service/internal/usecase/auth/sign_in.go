package auth_usecase

import (
	"context"
	"log"

	"github.com/adhyttungga/skeleton-application/services/user-service/internal/models"
	"github.com/adhyttungga/skeleton-application/services/user-service/utils"
	"golang.org/x/crypto/bcrypt"
)

func (uc *AuthUsecaseImpl) SignIn(c context.Context, req models.User) (string, error) {
	// Ignore name required tag
	req.Name = "lorem ipsum"

	// Validate user data
	if err := uc.Validate.Struct(req); err != nil {
		log.Println("Validation error: ", err)
		return "", err
	}

	user := models.User{
		Email: req.Email,
	}

	// Retrieve user by email
	if err := uc.Repository.FetchEmail(c, &user); err != nil {
		return "", err
	}

	// Validate password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		log.Println("Validation error: ", err)
		return "", err
	}

	// Generate JWT token
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	// Set token to cookies
	return token, nil
}
