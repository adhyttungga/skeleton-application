package user_handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

// Fetch godoc
// @Summary Fetch user by ID
// @Description Fetch user by ID
// @Tags User
// @Param id path string true "User ID"
// @Produce json
// @Success 200 {object} models.User
// @Failure 404 {object} gin.H "User not found"
// @Failure 500 {object} gin.H "Internal Server Error"
// @Router /api/v1/user/{id} [get]
func (hd *UserHandlerImpl) Fetch(c *gin.Context) {
	userID := c.Param("id")
	user, err := hd.UserUsecase.Fetch(c, userID)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintln("User fetched successfully with ID: ", userID), "user": user})
}
