package user_handler

import user_usecase "github.com/adhyttungga/skeleton-application/services/user-service/internal/usecase/user"

type UserHandlerImpl struct {
	UserUsecase user_usecase.UserUsecase
}

func NewUserHandler(userUsecase user_usecase.UserUsecase) UserHandler {
	return &UserHandlerImpl{
		UserUsecase: userUsecase,
	}
}
