package user_usecase

import (
	"context"
	"errors"
	"log"

	"github.com/adhyttungga/skeleton-application/services/user-service/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (usecase *UserUsecaseImpl) Update(c context.Context, id string, user *models.User) error {
	// Validate the user ID
	if id == "" {
		log.Println("User ID cannot be empty")
		return errors.New("user ID cannot be empty")
	}

	// Convert the string ID to ObjectID
	userID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Invalid user ID format:", err)
		return err
	}

	// Validate the user data
	if err := usecase.Validate.Struct(user); err != nil {
		log.Println("Validation error:", err)
		return err
	}

	// Call the repository to update the user
	if err := usecase.Repository.Update(c, userID, user); err != nil {
		return err
	}

	return nil
}
