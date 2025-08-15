package auth_handler

import auth_usecase "github.com/adhyttungga/skeleton-application/services/user-service/internal/usecase/auth"

type AuthHandlerImpl struct {
	AuthUsecase auth_usecase.AuthUsecase
}

func NewAuthHandler(authUsecase auth_usecase.AuthUsecase) AuthHandler {
	return &AuthHandlerImpl{
		AuthUsecase: authUsecase,
	}
}
