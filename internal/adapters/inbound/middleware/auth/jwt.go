package taskifyAuth

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	taskifyConfig "github.com/raihandotmd/taskify/internal/config"
	tokenModel "github.com/raihandotmd/taskify/internal/usecase/models/token"
	taskifyGin "github.com/raihandotmd/taskify/pkg/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenConfig, err := taskifyConfig.NewTokenConfig()
		if err != nil {
			taskifyGin.AbortJSONResponse(c, http.StatusInternalServerError, nil, errors.New("Failed to load token configuration"))
			return
		}

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			taskifyGin.AbortJSONResponse(c, http.StatusUnauthorized, nil, errors.New("Authorization header is required"))
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			taskifyGin.AbortJSONResponse(c, http.StatusUnauthorized, nil, errors.New("Invalid authorization header format"))
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
			taskifyGin.AbortJSONResponse(c, http.StatusUnauthorized, nil, errors.New("Invalid token"))
			return
		}

		if claims, ok := token.Claims.(*tokenModel.AccessToken); ok && token.Valid {
			// Set the user ID in the context for further use
			c.Set("user_id", claims.Data.UserId)
			c.Set("claims", claims)
			c.Next()
		} else {
			taskifyGin.AbortJSONResponse(c, http.StatusUnauthorized, nil, errors.New("Invalid token claims"))
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

func GetExpiresAt(c *gin.Context) (int64, error) {
	claims, exists := c.Get("claims")
	if !exists {
		return 0, errors.New("claims not found in context")
	}
	accessToken, ok := claims.(*tokenModel.AccessToken)
	if !ok {
		return 0, errors.New("claims is not of type AccessToken")
	}
	return accessToken.ExpiresAt.Unix(), nil
}
