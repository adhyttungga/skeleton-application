package mongo

import (
	"context"
	"log"

	"github.com/adhyttungga/skeleton-application/services/user-service/internal/models"
	"go.mongodb.org/mongo-driver/bson"
)

func (repo *RepositoryImpl) FetchEmail(c context.Context, user *models.User) error {
	if c.Err() == context.DeadlineExceeded {
		log.Println("Context deadline exceeded while creating user")
		return c.Err()
	}

	if err := repo.DB.Collection("users").FindOne(c, bson.M{"email": user.Email}).Decode(&user); err != nil {
		log.Println("Error fetching user from database:", err)
		return err
	}

	return nil
}
