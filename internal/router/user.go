package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ljinf/template_project_v2/internal/logic/handler"
)

func userRouter(
	root *gin.RouterGroup,
	userHandler handler.UserHandler,
) {

	userGroup := root.Group("/user")
	{
		userGroup.GET("/profile", userHandler.GetUserProfile)
	}
}
