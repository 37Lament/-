package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"yewu/tool"
)
//查看登陆状态
func auth(ctx *gin.Context) {
	username, err := ctx.Cookie("username")//检查cookie
	if err != nil {
		tool.RespErrorWithDate(ctx, "请登陆后进行操作")
		ctx.Abort()
	}

	ctx.Set("username", username)
	ctx.Next()
}
func Admin() gin.HandlerFunc {
	return func(c *gin.Context) {
		if cookie, err := c.Request.Cookie("username"); err == nil {
			value := cookie.Value
			fmt.Println(value)
			if value =="lxl" {
				c.Next()
				return
			}
		}
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "你不是管理员",
		})
		c.Abort()
		return
	}
}
func Ping(ctx *gin.Context)  {
	ctx.JSON(http.StatusOK,gin.H{
		"info": "ping",
	})
}