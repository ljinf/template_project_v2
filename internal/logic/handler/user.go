package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ljinf/template_project_v2/internal/logic/appservice"
	"github.com/ljinf/template_project_v2/pkg/app"
	"github.com/ljinf/template_project_v2/pkg/errcode"
)

type UserHandler interface {
	GetUserProfile(ctx *gin.Context)
}

type userHandler struct {
	userAppSrv appservice.UserAppService
}

func (h *userHandler) GetUserProfile(ctx *gin.Context) {
	uid := GetUserIdFromCtx(ctx)

	userProfile := h.userAppSrv.GetUserProfile(ctx, uid)

	if userProfile == nil {
		app.HandleError(ctx, errcode.ErrParams)
		return
	}
	app.HandleSuccess(ctx, userProfile)
}

func NewUserHandler(u appservice.UserAppService) UserHandler {
	return &userHandler{
		userAppSrv: u,
	}
}
