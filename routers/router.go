package routers

import (
	"github.com/gin-gonic/gin"
	"go-chats/app/http/controller"
	"go-chats/app/http/middleware"
)

func InitRouter(r *gin.Engine) {

	r.Any("/test", (&controller.PublicController{}).Test)                    // 测试
	r.Any("/login", (&controller.PublicController{}).Login)                  // 登录
	r.Any("/register", (&controller.PublicController{}).Register)            // 注册
	r.Any("/reset-password", (&controller.PublicController{}).ResetPassword) // 找回密码

	authorized := r.Group("/").Use(middleware.SessionValidate())
	{
		authorized.GET("logout", (&controller.PublicController{}).Logout) // 退出登录
		authorized.GET("index", (&controller.IndexController{}).Index)    // 主页
	}
}
