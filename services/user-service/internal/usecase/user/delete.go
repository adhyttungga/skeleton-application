package user_usecase

import (
	"context"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (usecase *UserUsecaseImpl) Delete(c context.Context, id string) error {
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

	// Call the repository to delete the user
	if err := usecase.Repository.Delete(c, userID); err != nil {
		return err
	}

	return nil
}
