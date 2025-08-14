package user_handler

import "github.com/gin-gonic/gin"

type UserHandler interface {
	Create(c *gin.Context)
	ListAll(c *gin.Context)
	Fetch(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}
