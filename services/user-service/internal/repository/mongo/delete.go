package mongo

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (repo *RepositoryImpl) Delete(c context.Context, id primitive.ObjectID) error {
	if c.Err() == context.DeadlineExceeded {
		log.Println("Context deadline exceeded while deleting user")
		return c.Err()
	}

	// Delete the user with the given ID from the database
	collection := repo.DB.Collection("users")
	filter := bson.M{"_id": id}
	_, err := collection.DeleteOne(c, filter)
	if err != nil {
		log.Println("Error deleting user from database:", err)
		return err
	}

	return nil
}
