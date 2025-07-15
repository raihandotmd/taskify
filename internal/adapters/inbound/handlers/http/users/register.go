package users

import (
	"github.com/gin-gonic/gin"

	ibModel "github.com/raihandotmd/taskify/internal/adapters/inbound/handlers/http/model/user"
	ucUser "github.com/raihandotmd/taskify/internal/usecase/users"
	taskifyGin "github.com/raihandotmd/taskify/pkg/gin"
)

var usecaseUser = ucUser.New()

func Register(gCtx *gin.Context) error {
	var request ibModel.RegisterRequest

	if err := taskifyGin.BindAndValidate(gCtx, &request); err != nil {
		return err
	}

	response, err := usecaseUser.Register(gCtx.Request.Context(), request.ToUcModel())
	if err != nil {
		return err
	}

	gCtx.JSON(201, response)
	return nil
}
