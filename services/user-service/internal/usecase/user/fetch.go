package user_usecase

import (
	"context"
	"errors"
	"log"

	"github.com/adhyttungga/skeleton-application/services/user-service/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (usecase *UserUsecaseImpl) Fetch(c context.Context, id string) (*models.User, error) {
	// Validate the ID format
	if id == "" {
		log.Println("User ID cannot be empty")
		return nil, errors.New("user ID cannot be empty")
	}

	// Convert the string ID to ObjectID
	userID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Invalid user ID format:", err)
		return nil, err
	}

	// Call the repository to fetch the user by ID
	user, err := usecase.Repository.Fetch(c, userID)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
