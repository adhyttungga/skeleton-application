package mongo

import (
	"context"
	"errors"
	"log"

	"github.com/adhyttungga/skeleton-application/services/user-service/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (repo *RepositoryImpl) Create(c context.Context, user *models.User) error {
	if c.Err() == context.DeadlineExceeded {
		log.Println("Context deadline exceeded while creating user")
		return c.Err()
	}

	res, err := repo.DB.Collection("users").InsertOne(c, user)
	if err != nil {
		log.Println("Error inserting user into database: ", err)
		return err
	}

	userID, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		err := errors.New("Failed to convert inserted ID")
		log.Printf("Error converting inserted id: %v", err)
		return err
	}

	user.ID = userID.Hex()
	return nil
}
