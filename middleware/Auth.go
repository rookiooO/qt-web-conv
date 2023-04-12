package middleware

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/rookiooO/qt-web-conv/constant"
	"github.com/rookiooO/qt-web-conv/server"
	"net/http"
)

type User struct {
	UserId        int64    `json:"userId"`
	Name          string   `json:"name"`
	DeptId        int64    `json:"deptId"`
	LoginTime     int64    `json:"loginTime"`
	ExpireTime    int64    `json:"expireTime"`
	LoginIP       string   `json:"loginIP"`
	LoginLocation string   `json:"loginLocation"`
	Browser       string   `json:"browser"`
	OS            string   `json:"os"`
	Permissions   []string `json:"permissions"`
	IsAdmin       bool     `json:"isAdmin"`
}

func AuthHandler(redis *redis.Client) gin.HandlerFunc {
	return func(context *gin.Context) {
		var (
			err   error
			key   string
			value string
			bytes []byte
			user  User
		)

		key = context.Request.Header.Get(constant.TokenKey)
		if key == "" {
			server.Result(http.StatusUnauthorized,
				http.StatusUnauthorized,
				nil,
				"请重新登录",
				context)
			context.Abort()
			return
		}

		if value, err = redis.Get(key).Result(); err != nil {
			server.Result(http.StatusUnauthorized,
				http.StatusUnauthorized,
				nil,
				"请重新登录",
				context)
			context.Abort()
			return
		}

		bytes = []byte(value)
		if err = json.Unmarshal(bytes, &user); err != nil {
			server.Result(http.StatusInternalServerError,
				http.StatusInternalServerError,
				nil,
				"系统错误，请联系管理员",
				context)
			// TODO: LOG
			context.Abort()
			return
		}

		redis.Expire(key, constant.TokenExpireTime)

		context.Set(constant.UserKey, &user)
		context.Next()
	}
}
