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
	{
		service := service.UserService{}
		user.GET("/test_get_jwt", service.TestGetJWT)
		user.POST("/login", service.Login)
		user.POST("/register", service.Register)
		user.POST("/forget", service.Forget)
	}
	user.Use(middleware.JWT(1))
	{
		service := service.UserService{}
		user.POST("/password", service.UpdatePassword)
		user.POST("/me", service.UpdateUser)
		user.GET("/me", service.Me)
	}
	user.Use(middleware.JWT(2))
	{
		service := service.UserService{}
		user.GET("/:id", service.GetUser)
		user.GET("/", service.GetAllUser)
	}
	address := engine.Group("/user/address")
	address.Use(middleware.JWT(1))
	{
		service := service.UserService{}
		address.POST("/", service.CreateAddress)
		address.POST("/:id", service.UpdateAddress)
		address.DELETE("/:id", service.DeleteAddress)
	}

	static := engine.Group("/static")
	{
		service := service.UserService{}
		static.GET("/gender", service.GetAllGender)
		static.GET("/campus", service.GetAllCampus)
	}

	good := engine.Group("/good")
	{
		service := service.GoodService{}
		good.GET("/banner", service.GetBanner)
		good.GET("/category", service.GetAllCategory)
		good.GET("/category/recommend/:category_id", service.GetRecommendedGoods)
		good.GET("/new", service.GetLatestGoods)
		good.GET("/campus/:campus_id", service.GetGoodsByCampus)
		good.POST("/temporary/time", service.GetUnfinishedGoodsByTime)
		good.POST("/all", service.GetAllGoods)
		good.POST("/cover/:good_id", service.UpdateGoodCover)
		good.POST("/picture/:good_id", service.UpdateGoodPictures)
		good.GET("/search", service.GetSearchGoods)
	}
	good.Use(middleware.JWT(1))
	{
		service := service.GoodService{}
		good.GET("/:id", service.GetGoodsDetailed)
		good.POST("/temporary/campus", service.GetUnfinishedGoodsByCampus)
		good.POST("/sell", service.CreateGood)
		good.DELETE("/:id", service.DeleteGood)
		good.POST("/my", service.GetMyGood)
	}

	order := engine.Group("/order")
	order.Use(middleware.JWT(1))
	{
		service := service.OrderService{}
		order.POST("/init", service.CreateOrder)
		order.DELETE("/:id", service.CancelOrder)
		order.POST("/submit", service.SubmitOrder)
		order.GET("/:id", service.GetOrder)
		order.POST("/buyer", service.GetBuyerOrder)
		order.POST("/seller", service.GetSellerOrder)
	}

	problem := engine.Group("/problem")
	problem.Use(middleware.JWT(1))
	{
		service := service.ProblemService{}
		problem.POST("/after_sale", service.CreateProblem)
		problem.POST("/get", service.GetMyProblem)
	}

	problem.Use(middleware.JWT(2))
	{
		service := service.ProblemService{}
		problem.POST("/all", service.GetAllProblem)
		problem.POST("/update", service.UpdateProblem)
	}

	r := engine.Group("/pay")
	{
		service := service.AlipayService{}
		r.GET("/alipay", service.Pay)
		r.GET("/alipay/callback", service.Callback)
		r.POST("/alipay/notify/:order_id", service.Notify)
	}

}
