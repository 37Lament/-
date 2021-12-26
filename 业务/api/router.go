package api

import "github.com/gin-gonic/gin"

func InitEngine() {
	engine := gin.Default()

	engine.POST("/register", register) //注册
	engine.POST("/login", login)       //登陆
	engine.POST("ping",Ping)//检查网络连接
	userGroup := engine.Group("/user")
	{
		userGroup.Use(auth)
		userGroup.POST("/password", changePassword) //修改密码
	}
	viewGroup := engine.Group("/view")//游客模式可使用的

	adminGroup:=engine.Group("/admin")
	{
		adminGroup.GET("/",Admin)
	}
	rechargeGroup:=engine.Group("/recharge")
	{
		recharge.GET("/",recharege)
	}
	adGroup:=engine.Group("ad")//广告，还没做
	engine.Run()
}
