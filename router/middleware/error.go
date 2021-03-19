package middleware

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				context.Status(http.StatusInternalServerError)

				h := gin.H{"code": "server err", "msg": err.(error).Error()}
				val, _ := json.Marshal(h)
				context.Writer.Write(val)
				context.Writer.Flush()
				context.Abort()
			}
		}()

		context.Next()
	}
}
