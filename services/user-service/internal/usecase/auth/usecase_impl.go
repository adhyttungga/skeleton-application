package auth_usecase

import (
	"github.com/adhyttungga/skeleton-application/services/user-service/internal/repository/mongo"
	"github.com/go-playground/validator/v10"
)

type AuthUsecaseImpl struct {
	Repository mongo.Repository
	Validate   *validator.Validate
}

func NewAuthUsecase(
	repository mongo.Repository,
	validate *validator.Validate,
) AuthUsecase {
	return &AuthUsecaseImpl{
		Repository: repository,
		Validate:   validate,
	}
}
