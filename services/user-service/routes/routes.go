package routes

import (
	"strings"
	"time"

	auth_handler "github.com/adhyttungga/skeleton-application/services/user-service/api/v1/auth"
	user_handler "github.com/adhyttungga/skeleton-application/services/user-service/api/v1/user"
	"github.com/adhyttungga/skeleton-application/services/user-service/config"
	repository "github.com/adhyttungga/skeleton-application/services/user-service/internal/repository/mongo"
	auth_usecase "github.com/adhyttungga/skeleton-application/services/user-service/internal/usecase/auth"
	user_usecase "github.com/adhyttungga/skeleton-application/services/user-service/internal/usecase/user"
	"github.com/adhyttungga/skeleton-application/services/user-service/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewRouter(DB *mongo.Database) *gin.Engine {
	router := gin.Default()
	allowOrigins := strings.Split(config.Config.Origin.AllowedOrigins, ",")

	router.Use(cors.New(cors.Config{
		AllowOrigins:     allowOrigins,
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-length"},
		MaxAge:           12 * time.Hour,
		AllowCredentials: true,
	}))

	// Initialize validator
	validate := validator.New()

	// Initialize repository
	repository := repository.NewRepository(DB)

	// Initialize user usecase
	userUsecase := user_usecase.NewUserUsecase(repository, validate)
	authUsecase := auth_usecase.NewAuthUsecase(repository, validate)

	// Initialize user handler
	userHandler := user_handler.NewUserHandler(userUsecase)
	authHandler := auth_handler.NewAuthHandler(authUsecase)

	userRoutes := router.Group("/api/v1/user")
	{
		userRoutes.POST("/", userHandler.Create)
		userRoutes.Use(middleware.ProtectRoute()).GET("/", userHandler.ListAll)
		userRoutes.Use(middleware.ProtectRoute()).GET("/:id", userHandler.Fetch)
		userRoutes.Use(middleware.ProtectRoute()).PUT("/:id", userHandler.Update)
		userRoutes.Use(middleware.ProtectRoute()).DELETE("/:id", userHandler.Delete)
	}
	authRoutes := router.Group("/api/v1/auth")
	{
		authRoutes.POST("/signin", authHandler.SignIn)
		authRoutes.POST("/signout", authHandler.SignOut)
	}

	return router
}
