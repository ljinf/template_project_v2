package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func demo(g *gin.RouterGroup) {

	g.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			":)": "Thank you for using delicate moods!",
		})
	})

}
