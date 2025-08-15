package utils

import (
	"log"
	"time"

	"github.com/adhyttungga/skeleton-application/services/user-service/config"
	jwt "github.com/dgrijalva/jwt-go"
)

func GenerateToken(userID string) (string, error) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(config.Config.PrivateKey))
	if err != nil {
		log.Println("Error parse the private key: ", err.Error())
		return "", err
	}

	now := time.Now().UTC()
	claims := make(jwt.MapClaims)
	claims["dat"] = userID                             // Stored data.
	claims["exp"] = now.Add(5 * 24 * time.Hour).Unix() // The expiration time after which the token must be disregard.
	claims["iat"] = now.Unix()                         // The time at which the token was issued.
	claims["nbf"] = now.Unix()                         // The time before which the token must be disregarded.

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)
	if err != nil {
		log.Println("Error generating token: ", err.Error())
		return "", err
	}

	return token, nil
}
