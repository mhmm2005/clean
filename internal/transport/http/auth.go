package http

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

// validateToken - validates an incoming jwt token
func validateToken(accessToken string) bool {
	var mySigningKey = []byte("s3kR3tK3y")
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("could not validate auth token")
		}
		return mySigningKey, nil
	})

	if err != nil {
		return false
	}

	return token.Valid
}

// JWTAuth - a handy middleware function that will provide basic auth around specific endpoints
func JWTAuth(original func(c *gin.Context)) func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{
				"message": "an unauthorized request has been made",
			})
			return
		}

		authHeaderParts := strings.Split(authHeader, " ")
		if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
			c.JSON(401, gin.H{
				"message": "authorization header could not be parsed",
			})
			return
		}

		if validateToken(authHeaderParts[1]) {
			c.Header("Content-Type", "application/json")
			original(c)
		} else {
			c.JSON(401, gin.H{
				"message": "could not validate incoming token",
			})
			return
		}
	}
}
