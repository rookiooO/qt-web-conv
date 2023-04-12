package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CorsHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		context.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		context.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		//放行所有OPTIONS方法
		method := context.Request.Method
		if method == "OPTIONS" {
			context.JSON(http.StatusOK, "Options Request!")
			return
		}
		context.Next()
	}
}
