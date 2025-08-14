package mongo

import (
	"context"
	"log"

	"github.com/adhyttungga/skeleton-application/services/user-service/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (repo *RepositoryImpl) Update(c context.Context, id primitive.ObjectID, user *models.User) error {
	if c.Err() == context.DeadlineExceeded {
		log.Println("Context deadline exceeded while updating user")
		return c.Err()
	}

	// Update the user with the given ID in the database
	filter := bson.M{"_id": id}
	update := bson.M{"$set": user}
	result, err := repo.DB.Collection("users").UpdateOne(c, filter, update)
	if err != nil {
		log.Println("Error updating user in database:", err)
		return err
	}

	if result.MatchedCount == 0 {
		log.Println("No user found with ID:", id.String())
		return mongo.ErrNoDocuments // No user found to update
	}

	return nil
}
