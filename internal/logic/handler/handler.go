package handler

import (
	"github.com/gin-gonic/gin"
)

func GetUserIdFromCtx(ctx *gin.Context) string {
	//v, exists := ctx.Get("claims")
	//if !exists {
	//	return ""
	//}
	//return v.(*jwt.MyCustomClaims).UserId

	return ""
}
