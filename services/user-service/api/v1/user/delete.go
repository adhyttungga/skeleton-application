package user_handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Delete godoc
// @Summary Delete a user
// @Description Delete a user by ID
// @Tags User
// @Param id path string true "User ID"
// @Produce json
// @Success 200 {object} gin.H "User deleted successfully"
// @Failure 500 {object} gin.H "Internal Server Error"
// @Router /api/v1/user/{id} [delete]
func (hd *UserHandlerImpl) Delete(c *gin.Context) {
	userID := c.Param("id")
	err := hd.UserUsecase.Delete(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("User deleted successfully with ID: %s", userID)})
}
