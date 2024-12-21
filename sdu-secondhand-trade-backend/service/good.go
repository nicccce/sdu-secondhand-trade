package service

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"sdu-secondhand-trade-backend/app"
	"sdu-secondhand-trade-backend/model"
	"sdu-secondhand-trade-backend/util"
	"strconv"
	"time"
)

type GoodService struct{}

type GoodVo struct {
	model.Good
	Pictures []model.Pictures `json:"pictures,omitempty"`
}

func (receiver GoodService) GetAllCategory(c *gin.Context) {
	aw := app.NewWrapper(c)
	var Category []model.FirstCategory
	var err error
	if Category, err = categoryModel.GetCategoriesWithSubCategories(); err != nil {
		aw.Error(err.Error())
		return
	}
	aw.Success(Category)
}

func (receiver GoodService) GetBanner(c *gin.Context) {
	aw := app.NewWrapper(c)
	var Banners []model.Banner
	var err error
	if Banners, err = bannerModel.GetBanners(); err != nil {
		aw.Error(err.Error())
		return
	}
	aw.Success(Banners)
}

func (receiver GoodService) GetRecommendedGoods(c *gin.Context) {
	aw := app.NewWrapper(c)
	var err error
	var categoryId int

	// 获取 categoryId
	if categoryId, err = strconv.Atoi(c.Param("category_id")); err == nil && categoryId < 0 {
		aw.Error("非法分类")
		return
	}

	var Children []model.Category
	var categoryMap = make(map[string]bool)
	var Goods []model.GoodBrief
	var recommendedGoods []model.GoodBrief

	// 获取所有未完成的商品
	if Goods, err = goodModel.GetAllUnfinishedGoods(); err != nil {
		aw.Error(err.Error())
		return
	}

	categoryLL, _ := categoryModel.GetCategoryById(categoryId)

	if categoryLL.FatherID == 0 { //父分类
		// 获取子分类列表
		if Children, err = categoryModel.GetSubCategories(categoryId); err != nil {
			aw.Error(err.Error())
			return
		}
		// 使用 map 存储子分类名称，提高查找效率
		for _, category := range Children {
			categoryMap[category.Name] = true
		}

		// 筛选商品
		for _, good := range Goods {
			categoryL, err := categoryModel.GetCategoryById(good.Category)
			if err != nil {
				aw.Error(err.Error())
				return
			}
			if categoryMap[categoryL.Name] {
				recommendedGoods = append(recommendedGoods, good)
			}
		}

	} else { //子分类

		categoryMap[categoryLL.Name] = true
		// 筛选商品
		for _, good := range Goods {
			categoryL, err := categoryModel.GetCategoryById(good.Category)
			if err != nil {
				aw.Error(err.Error())
				return
			}
			if categoryMap[categoryL.Name] {
				recommendedGoods = append(recommendedGoods, good)
			}
		}

	}

	recommendedGoods = GetRecommended(recommendedGoods)
	aw.Success(recommendedGoods)

}

// GetRecommended 根据用户喜好推荐商品
func GetRecommended(recommendedGoods []model.GoodBrief) []model.GoodBrief {
	if len(recommendedGoods) > 16 {
		rand.Seed(time.Now().UnixNano())
		var selectedGoods []model.GoodBrief
		for len(selectedGoods) < 16 {
			randomIndex := rand.Intn(len(recommendedGoods))
			selectedGoods = append(selectedGoods, recommendedGoods[randomIndex])
			recommendedGoods = append(recommendedGoods[:randomIndex], recommendedGoods[randomIndex+1:]...)
		}
		return selectedGoods
	} else {
		return recommendedGoods
	}

}

func (receiver GoodService) GetLatestGoods(c *gin.Context) {
	aw := app.NewWrapper(c)
	var New []model.GoodBrief
	var err error
	if New, err = goodModel.GetLatestGoods(); err != nil {
		aw.Error(err.Error())
		return
	}
	aw.Success(New)
}

func (receiver GoodService) GetGoodsByCampus(c *gin.Context) {
	aw := app.NewWrapper(c)
	var err error
	var campusId int

	if campusId, err = strconv.Atoi(c.Param("campus_id")); err == nil && campusId < 0 {
		aw.Error("非法校区")
		return
	}
	var CampusGoods []model.GoodBrief
	if CampusGoods, err = goodModel.GetGoodsByCampus(campusId); err != nil {
		aw.Error(err.Error())
		return
	}
	aw.Success(CampusGoods)
}
func (receiver GoodService) GetGoodsDetailed(c *gin.Context) {
	aw := app.NewWrapper(c)
	var err error
	var goodId int

	if goodId, err = strconv.Atoi(c.Param("id")); err == nil && goodId < 0 {
		aw.Error("非法商品")
		return
	}
	var GoodsDetailed GoodVo
	var CGoods model.Good
	var Picture []model.Pictures
	if CGoods, err = goodModel.GetGoodByID(goodId); err != nil {
		aw.Error(err.Error())
		return
	}
	Picture = picturesModel.FindPicturesByGoodID(goodId)

	GoodsDetailed = GoodVo{
		Good:     CGoods,
		Pictures: Picture,
	}
	aw.Success(GoodsDetailed)
}

func (receiver GoodService) GetUnfinishedGoodsByTime(c *gin.Context) {
	aw := app.NewWrapper(c)
	type GetUnfinishedGoodsReq struct {
		CategoryID int `json:"category_id" binding:"required"`
		Page       int `json:"page" binding:"required,min=1"`
		PageSize   int `json:"page_size" binding:"required,min=1"`
	}

	var req GetUnfinishedGoodsReq
	if err := c.ShouldBind(&req); err != nil {
		aw.Error("请求参数错误: " + err.Error())
		return
	}
	goods, err := goodModel.GetUnfinishedGoodsPaginated(req.CategoryID, req.Page, req.PageSize, "created_at", 0)
	if err != nil {
		aw.Error(err.Error())
		return
	}
	aw.Success(goods)
}

func (receiver GoodService) GetUnfinishedGoodsByCampus(c *gin.Context) {
	aw := app.NewWrapper(c)
	userClaim := util.ExtractUserClaims(c)

	type GetUnfinishedGoodsReq struct {
		CategoryID int `json:"category_id" binding:"required"`
		Page       int `json:"page" binding:"required,min=1"`
		PageSize   int `json:"page-size" binding:"required,min=1"`
	}

	var req GetUnfinishedGoodsReq
	if err := c.ShouldBind(&req); err != nil {
		aw.Error("请求参数错误: " + err.Error())
		return
	}

	user, err := userModel.FindUserByID(userClaim.UserID)
	if err != nil {
		aw.Error(err.Error())
		return
	}
	if user == nil {
		aw.Error("用户名不存在")
		return
	}

	var campusId int
	campusId = user.Campus

	goods, err := goodModel.GetUnfinishedGoodsPaginated(req.CategoryID, req.Page, req.PageSize, "campus", campusId)
	if err != nil {
		aw.Error(err.Error())
		return
	}

	aw.Success(goods)
}

func (receiver GoodService) CreateGood(c *gin.Context) {
	aw := app.NewWrapper(c)
	userClaim := util.ExtractUserClaims(c)
	type updateReq struct {
		model.Good
	}
	var req updateReq
	if err := c.ShouldBind(&req); err != nil {
		aw.Error(err.Error())
		return
	}
	req.Seller = userClaim.UserID
	goodModel.CreateGood(&req.Good)
	aw.OK()
}
