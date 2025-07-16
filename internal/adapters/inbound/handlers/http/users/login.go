package users

import (
	"github.com/gin-gonic/gin"
	ibModel "github.com/raihandotmd/taskify/internal/adapters/inbound/handlers/http/model/user"
	taskifyGin "github.com/raihandotmd/taskify/pkg/gin"
)

func Login(gCtx *gin.Context) error {
	var request ibModel.LoginRequest

	if err := taskifyGin.BindAndValidate(gCtx, &request); err != nil {
		return err
	}

	response, err := usecaseUser.Login(gCtx.Request.Context(), request.ToUcModel())
	if err != nil {
		gCtx.AbortWithStatusJSON(500, gin.H{"error": "Internal server error, please try again later"})
		return nil
	}

	gCtx.JSON(200, response)
	return nil
}
