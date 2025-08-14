package user_handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/adhyttungga/skeleton-application/services/user-service/internal/models"
	"github.com/adhyttungga/skeleton-application/services/user-service/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

// Create godoc
// @Summary Create a new user
// @Description Create a new user with the provided details
// @Tags User
// @Accept json
// @Produce json
// @Param user body models.User true "User details"
// @Success 201 {object} gin.H "Created", models.User
// @Failure 400 {object} gin.H "Bad Request", []utils.ErrorAPI
// @Failure 500 {object} gin.H "Internal Server Error", []utils.ErrorAPI
// @Router /api/v1/user [post]
func (hd *UserHandlerImpl) Create(c *gin.Context) {
	var req models.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": []utils.ErrorAPI{{
			Field: "",
			Msg:   fmt.Sprintf("Invalid request body: %v", err),
		}}})
		return
	}

	if err := hd.UserUsecase.Create(c, &req); err != nil {
		var ve validator.ValidationErrors
		// Return bad request if error from validator or
		// duplicate key error. Default is Internal server error
		if errors.As(err, &ve) {
			out := make([]utils.ErrorAPI, len(ve))
			for i, fe := range ve {
				out[i] = utils.ErrorAPI{
					Field: fe.Field(),
					Msg:   fmt.Sprintf("%s %s", fe.Field(), utils.MsgFromTag(fe.Tag())),
				}
			}
			c.JSON(http.StatusBadRequest, gin.H{"error": out})
		} else if mongo.IsDuplicateKeyError(err) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": []utils.ErrorAPI{{
					Field: "Email",
					Msg:   fmt.Sprintf("User already exists with email: %s", req.Email),
				}},
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": []utils.ErrorAPI{{
				Field: "",
				Msg:   "Failed to create user",
			}}})
		}
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user": req})
}
