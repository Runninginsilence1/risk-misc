package riskmisc

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// simulate echo gin web server for testing

func GenGinEngine(relativePath string, data any) (r *gin.Engine, err error) {
	r = gin.New()
	r.GET(relativePath, func(ctx *gin.Context) {
		ctx.Json(http.StatusOK, gin.H{
			"code": 200,
			"data": data,
		})
	})
	return
}
