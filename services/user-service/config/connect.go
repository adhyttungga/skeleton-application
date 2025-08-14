package config

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(ctx context.Context) *mongo.Client {
	log.Println("Connecting to the database...")
	// Create options with credential
	// credential := options.Credential{
	// 	Username: Config.Database.Username,
	// 	Password: Config.Database.Password,
	// }
	// clientOptions := options.Client().ApplyURI(Config.Database.URI).SetAuth(credential)
	clientOptions := options.Client().ApplyURI(Config.Database.URI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}

	return client
}
