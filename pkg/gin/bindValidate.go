package taskifyGin

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func BindAndValidate(gCtx *gin.Context, request interface{}) error {
	var validate = validator.New()

	if err := gCtx.ShouldBindJSON(request); err != nil {
		return err
	}

	if err := validate.Struct(request); err != nil {
		return err
	}

	return nil
}
