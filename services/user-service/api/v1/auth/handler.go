package auth_handler

import "github.com/gin-gonic/gin"

type AuthHandler interface {
	SignIn(c *gin.Context)
	SignOut(c *gin.Context)
}
