package mongo

import (
	"context"
	"log"

	"github.com/adhyttungga/skeleton-application/services/user-service/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (repo *RepositoryImpl) Fetch(c context.Context, id primitive.ObjectID) (models.User, error) {
	if c.Err() == context.DeadlineExceeded {
		log.Println("Context deadline exceeded while fetching user")
		return models.User{}, c.Err()
	}

	var user models.User
	if err := repo.DB.Collection("users").FindOne(c, bson.M{"_id": id}).Decode(&user); err != nil {
		log.Println("Error fetching user from database:", err)
		return models.User{}, err
	}

	return user, nil
}
