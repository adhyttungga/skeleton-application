package user_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

// ListAll godoc
// @Summary List all users
// @Description Get a list of all users
// @Tags User
// @Produce json
// @Success 200 {array} models.User
// @Failure 404 {object} gin.H "Users not found"
// @Failure 500 {object} gin.H "Internal Server Error"
// @Router /api/v1/user [get]
func (hd *UserHandlerImpl) ListAll(c *gin.Context) {
	users, err := hd.UserUsecase.ListAll(c)
	if err != nil {
		if err == mongo.ErrNilDocument {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list users"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully listed all users", "user": users})
}
