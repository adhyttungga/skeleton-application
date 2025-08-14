package mongo

import (
	"context"
	"log"

	"github.com/adhyttungga/skeleton-application/services/user-service/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (repo *RepositoryImpl) ListAll(c context.Context) ([]models.User, error) {
	if c.Err() == context.DeadlineExceeded {
		log.Println("Context deadline exceeded while listing all users")
		return nil, c.Err()
	}

	collection := repo.DB.Collection("users")
	cursor, err := collection.Find(c, bson.M{})
	if err != nil {
		log.Println("Error fetching users from database:", err)
		return nil, err
	}

	// Ensure the cursor is closed after use
	defer cursor.Close(c)

	var users []models.User
	for cursor.Next(c) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			log.Println("Error decoding user from cursor:", err)
			return nil, err
		}
		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		log.Println("Error iterating over cursor:", err)
		return nil, err
	}

	if len(users) <= 0 {
		log.Println("Error no user found")
		return nil, mongo.ErrNilDocument
	}

	return users, nil
}
