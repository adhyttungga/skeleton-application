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

// Update godoc
// @Summary Update user by ID
// @Description Update user by ID with the provided details
// @Tags User
// @Param id path string true "User ID"
// @Accept json
// @Produce json
// @Param user body models.User true "User details"
// @Success 200 {object} models.User
// @Failure 400 {object} gin.H "Bad Request", []utils.ErrorAPI
// @Failure 404 {object} gin.H "User not found", []utils.ErrorAPI
// @Failure 500 {object} gin.H "Internal Server Error", []utils.ErrorAPI
// @Router /api/v1/user/{id} [put]
func (hd *UserHandlerImpl) Update(c *gin.Context) {
	userID := c.Param("id")
	var updatedUser models.User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": []utils.ErrorAPI{{
			Field: "",
			Msg:   fmt.Sprintf("Invalid request body: %v", err),
		}}})
		return
	}

	if err := hd.UserUsecase.Update(c, userID, &updatedUser); err != nil {
		// Return bad request if error from validator or
		// duplicate key error.
		// Return error not found if no user with the id
		// Default is Internal server error
		var ve validator.ValidationErrors
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
					Msg:   fmt.Sprintf("User already exists with email: %s", updatedUser.Email),
				}},
			})
		} else if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": []utils.ErrorAPI{{
				Field: "",
				Msg:   fmt.Sprintf("No user found with ID: %s", userID)},
			}})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": []utils.ErrorAPI{{
				Field: "",
				Msg:   fmt.Sprintf("Failed to update user: %v", err),
			}}})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("User updated successfully with ID: %s", userID), "user": updatedUser})
}
