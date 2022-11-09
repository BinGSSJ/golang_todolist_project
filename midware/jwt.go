package midware

import (
	"github.com/BINGSSJ/golang_todolist_project/pkg/utils"
	"github.com/gin-gonic/gin"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := 200
		token := c.GetHeader("Authorization")
		if token == "" {
			code = 404 // 无token
		} else {
			claims, err := utils.ParseToken(token)
			if err != nil { // 解析错误
				code = 403
			} else if time.Now().Unix() > claims.ExpiresAt { // token过期
				code = 401
			}
		}
		if code != 200 {
			c.JSON(code, gin.H{
				"status": code,
				"msg":    "token parse ERROR",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
