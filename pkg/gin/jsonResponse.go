package taskifyGin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	Response struct {
		Success bool   `json:"success"`
		Data    any    `json:"data"`
		Error   *Error `json:"error"`
	}

	Error struct {
		Message string `json:"message"`
		Title   string `json:"title"`
	}

	ApiError struct {
		ErrCode string `json:"error_code"`
		Message string `json:"message"`
	}
)

var errorMap = map[int]ApiError{
	400: {
		ErrCode: "BAD_REQUEST",
		Message: "The request could not be understood by the server due to malformed syntax.",
	},
	401: {
		ErrCode: "UNAUTHORIZED",
		Message: "The request requires user authentication.",
	},
	404: {
		ErrCode: "NOT_FOUND",
		Message: "The requested resource could not be found on the server.",
	},
	500: {
		ErrCode: "INTERNAL_SERVER_ERROR",
		Message: "The server encountered an unexpected condition which prevented it from fulfilling the request.",
	},
}

func NewJSONResponse(c *gin.Context, statusCode int, data any, err any) error {
	if statusCode >= http.StatusBadRequest {
		c.JSON(statusCode, errorResponse(statusCode, err))
	} else {
		c.JSON(statusCode, successResponse(data))
	}

	return nil
}

func AbortJSONResponse(c *gin.Context, statusCode int, data any, err any) {
	if statusCode >= http.StatusBadRequest {
		c.AbortWithStatusJSON(statusCode, errorResponse(statusCode, err))
	} else {
		c.AbortWithStatusJSON(statusCode, successResponse(data))
	}
}

func successResponse(data any) Response {
	var response = Response{
		Success: true,
		Data:    data,
	}
	return response
}

func errorResponse(statusCode int, err any) Response {
	var response = Response{
		Success: false,
	}

	if err != nil {
		errResponse, ok := err.(error)
		if ok {
			apiError, exists := errorMap[statusCode]
			if !exists {
				apiError = ApiError{
					ErrCode: "UNKNOWN_ERROR",
					Message: "An unknown error occurred.",
				}
				response.Error = &Error{
					Title:   apiError.ErrCode,
					Message: apiError.Message,
				}
			}
			response.Error = &Error{
				Title:   apiError.ErrCode,
				Message: errResponse.Error(),
			}

		}
	}
	return response
}
