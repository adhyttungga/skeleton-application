package auth_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateTags godoc
// @Summary Sign out endpoint
// @Tags auth
// @Success 200 {object} gin.H "Ok"
// @Router /api/v1/auth/signout [post]
func (hd *AuthHandlerImpl) SignOut(c *gin.Context) {
	c.SetCookie("jwt", "", 0, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Singing out successfully"})
}
