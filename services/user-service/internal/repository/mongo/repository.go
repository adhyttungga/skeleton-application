package mongo

import (
	"context"

	"github.com/adhyttungga/skeleton-application/services/user-service/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Repository interface {
	Create(c context.Context, user *models.User) error
	Fetch(c context.Context, id primitive.ObjectID) (models.User, error)
	ListAll(c context.Context) ([]models.User, error)
	Update(c context.Context, id primitive.ObjectID, user *models.User) error
	Delete(c context.Context, id primitive.ObjectID) error
}
