package router

import (
	"github.com/gin-gonic/gin"
	"sdu-secondhand-trade-backend/app"
	"sdu-secondhand-trade-backend/middleware"
	"sdu-secondhand-trade-backend/model"
	"sdu-secondhand-trade-backend/service"
)

func Setup(engine *gin.Engine) {
	// 测试 上线后注释掉
	test := engine.Group("/test")
	{
		//测试panic
		test.GET("/panic", func(c *gin.Context) {
			panic("test panic")
		})
		//初始化数据库
		test.GET("/database_initialization", func(c *gin.Context) {
			aw := app.NewWrapper(c)
			err := model.Database_initialization()
			if err != nil {
				aw.Error(err.Error())
			}
			aw.Success("success!")
		})
	}

	// 用户
	user := engine.Group("/users")
	{
		hub := service.UserService{}
		user.GET("/test_get_jwt", hub.TestGetJWT)
		user.POST("/login", hub.Login)
		user.POST("/register", hub.Register)
	}
	user.Use(middleware.JWT(1))
	{
	}
}
