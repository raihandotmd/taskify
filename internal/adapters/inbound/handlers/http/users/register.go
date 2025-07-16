package users

import (
	"net/http"

	"github.com/gin-gonic/gin"

	ibModel "github.com/raihandotmd/taskify/internal/adapters/inbound/handlers/http/model/user"
	taskifyGin "github.com/raihandotmd/taskify/pkg/gin"
)

func Register(gCtx *gin.Context) {
	var request ibModel.RegisterRequest

	if err := taskifyGin.BindAndValidate(gCtx, &request); err != nil {
		taskifyGin.NewJSONResponse(gCtx, http.StatusBadRequest, nil, err)
		return
	}

	response, err := usecaseUser.Register(gCtx.Request.Context(), request.ToUcModel())
	if err != nil {
		taskifyGin.NewJSONResponse(gCtx, http.StatusInternalServerError, nil, err)
		return
	}

	taskifyGin.NewJSONResponse(gCtx, http.StatusCreated, response, nil)
}
