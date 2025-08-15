package middleware

import (
	"fmt"
	"log"
	"net/http"

	"github.com/adhyttungga/skeleton-application/services/user-service/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func ProtectRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("jwt")
		if err != nil {
			log.Printf("Error retrieving token from cookie: %v", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token invalid"})
			return
		}

		key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(config.Config.PublicKey))
		if err != nil {
			log.Printf("Error parsing the public key: %v", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		parsedToken, err := jwt.Parse(token, func(jwtToken *jwt.Token) (any, error) {
			if _, ok := jwtToken.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("Unexpected method: %s", jwtToken.Header["alg"])
			}

			return key, nil
		})

		if err != nil {
			log.Printf("Error parsing jwt token: %v", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token invalid"})
			return
		}

		claims, ok := parsedToken.Claims.(jwt.MapClaims)
		if !ok || !parsedToken.Valid {
			log.Printf("Error validating token: %v", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token invalid"})
			return
		}

		c.Set("id", claims["dat"])
	}
}
