package mongo

import "go.mongodb.org/mongo-driver/mongo"

type RepositoryImpl struct {
	DB *mongo.Database
}

func NewRepository(db *mongo.Database) Repository {
	return &RepositoryImpl{DB: db}
}
