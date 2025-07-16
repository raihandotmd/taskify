package taskifyAuth

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	taskifyConfig "github.com/raihandotmd/taskify/internal/config"
	tokenModel "github.com/raihandotmd/taskify/internal/usecase/models/token"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenConfig, err := taskifyConfig.NewTokenConfig()
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": "Failed to load token configuration"})
			return
		}

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Authorization header is required"})
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid authorization header format"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		// Parse and validate the token
		token, err := jwt.ParseWithClaims(tokenString, &tokenModel.AccessToken{}, func(token *jwt.Token) (interface{}, error) {
			// Validate the signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(tokenConfig.TokenSecret), nil
		})

		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid token"})
			return
		}

		if claims, ok := token.Claims.(*tokenModel.AccessToken); ok && token.Valid {
			// Set the user ID in the context for further use
			c.Set("user_id", claims.Data.UserId)
			c.Set("claims", claims)
			c.Next()
		} else {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid token claims"})
			return
		}

	}
}

func GetUserId(c *gin.Context) (string, error) {
	userID, exists := c.Get("user_id")
	if !exists {
		return "", errors.New("user_id not found in context")
	}
	id, ok := userID.(string)
	if !ok {
		return "", errors.New("user_id is not a string")
	}
	return id, nil
}
