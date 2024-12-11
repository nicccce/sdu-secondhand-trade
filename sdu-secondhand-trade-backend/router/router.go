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
	user := engine.Group("/user")
	address := user.Group("/address")
	{
		service := service.UserService{}
		user.GET("/test_get_jwt", service.TestGetJWT)
		user.POST("/login", service.Login)
		user.POST("/register", service.Register)
	}
	user.Use(middleware.JWT(1))
	{
		service := service.UserService{}
		user.GET("/me", service.Me)
		user.POST("/password", service.UpdatePassword)
		user.POST("/me", service.UpdateUser)
		address.POST("/", service.CreateAddress)
		address.POST("/:id", service.UpdateAddress)
		address.DELETE("/:id", service.DeleteAddress)
	}
	user.Use(middleware.JWT(2))
	{
		service := service.UserService{}
		user.GET("/:id", service.GetUser)
		user.GET("/", service.GetAllUser)
		user.POST("/password", service.UpdatePassword)
		user.POST("/me", service.UpdateUser)
		address.POST("/", service.CreateAddress)
		address.POST("/:id", service.UpdateAddress)
		address.DELETE("/:id", service.DeleteAddress)
	}
}
