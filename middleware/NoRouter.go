package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/rookiooO/qt-web-conv/server"
	"net/http"
)

func NoRouteHandler(context *gin.Context) {
	server.Result(http.StatusNotFound, http.StatusNotFound, nil, "", context)
}
