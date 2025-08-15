package auth_handler

import (
	"errors"
	"net/http"

	"github.com/adhyttungga/skeleton-application/services/user-service/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

// CreateTags godoc
// @Summary Sign in endpoint
// @Tags auth
// @Accept json
// @Produce json
// @Param user body models.User true "User details"
// @Success 202 {object} gin.H "Accept"
// @Success 400 {object} gin.H "Bad Request"
// @Success 500 {object} gin.H "Internal Server Error"
// @Success 202 {object} gin.H "Accept"
// @Router /api/v1/auth/signin [post]
func (hd *AuthHandlerImpl) SignIn(c *gin.Context) {
	var req models.User

	// Parse json request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid credential"})
		return
	}

	token, err := hd.AuthUsecase.SignIn(c, req)
	if err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) || errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) || errors.Is(err, mongo.ErrNoDocuments) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid credential"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to signing in"})
		}
		return
	}

	c.SetCookie("jwt", token, (15 * 24 * 60 * 60), "/", "", false, true)
	c.JSON(http.StatusAccepted, gin.H{"message": "Signing in successfully"})
}
