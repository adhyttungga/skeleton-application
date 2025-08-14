package user_usecase

import (
	"github.com/adhyttungga/skeleton-application/services/user-service/internal/repository/mongo"
	"github.com/go-playground/validator/v10"
)

type UserUsecaseImpl struct {
	Repository mongo.Repository
	Validate   *validator.Validate
}

func NewUserUsecase(repository mongo.Repository, validate *validator.Validate) UserUsecase {
	return &UserUsecaseImpl{
		Repository: repository,
		Validate:   validate,
	}
}
