package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/rookiooO/qt-web-conv/server"
	"github.com/sirupsen/logrus"
	"net/http"
)

func Recovery(logger *logrus.Entry) gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				switch err.(type) {
				case server.JsonResult:
					result := err.(server.JsonResult)
					server.Failed(result.Code, result.Msg, context)
					logger.Errorln(err)
					context.Abort()
					return
				default:
					logger.Errorln(err)
					server.Failed(http.StatusInternalServerError,
						server.StatusMsg[http.StatusInternalServerError],
						context)
					context.Abort()
					return
				}

			}
		}()
		context.Next()
	}
}
