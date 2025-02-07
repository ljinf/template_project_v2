package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ljinf/template_project_v2/internal/logic/handler"
)

func RegistryRouter(
	engin *gin.RouterGroup,
	userHandler handler.UserHandler,
) {

	demo(engin)
	userRouter(engin, userHandler)
}
