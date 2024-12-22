package service

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"os"
	"path/filepath"
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

func (receiver GoodService) UpdateGoodCover(c *gin.Context) {
	aw := app.NewWrapper(c)
	var goodId int
	var err error

	// 获取商品ID
	if goodId, err = strconv.Atoi(c.Param("good_id")); err != nil || goodId < 1 {
		aw.Error("无效的商品ID")
		return
	}

	// 获取商品对象
	good, err := goodModel.GetGoodByID(goodId)
	if err != nil {
		aw.Error(err.Error())
		return
	}

	// 获取前端上传的文件，字段名为"cover"
	file, err := c.FormFile("cover")
	if err != nil {
		aw.Error("文件上传失败: " + err.Error())
		return
	}

	// 检查上传的文件类型，确保是图片格式
	// 你可以根据需要进一步完善这个检查，比如限制文件的扩展名或 MIME 类型
	if !util.IsImage(file) {
		aw.Error("请上传有效的图片文件")
		return
	}

	// 获取文件扩展名
	ext := filepath.Ext(file.Filename)
	if ext == "" {
		aw.Error("文件扩展名无效")
		return
	}

	// 生成新的文件名，避免文件名重复
	timestamp := time.Now().Unix()
	newFileName := strconv.FormatInt(timestamp, 10) + ext

	// 文件存储路径（项目根目录下的 "cover" 文件夹）
	savePath := "../cover/" + newFileName

	// 确保文件夹存在
	if err := os.MkdirAll("../cover", os.ModePerm); err != nil {
		aw.Error("无法创建文件夹: " + err.Error())
		return
	}

	// 保存文件到服务器
	if err := c.SaveUploadedFile(file, savePath); err != nil {
		aw.Error("文件保存失败: " + err.Error())
		return
	}

	// 将文件路径赋值给good对象的cover字段
	good.Cover = savePath

	//判断
	pictures := picturesModel.FindPicturesByGoodID(goodId)
	if len(pictures) != 0 {
		good.IsEffective = true
	}

	// 更新数据库中的good记录
	if err := goodModel.UpdateGood(&good); err != nil {
		aw.Error("更新商品封面失败: " + err.Error())
		return
	}

	// 返回成功
	aw.OK()
}

func (receiver GoodService) UpdateGoodPictures(c *gin.Context) {
	aw := app.NewWrapper(c)
	var goodId int
	var err error

	// 获取商品ID
	if goodId, err = strconv.Atoi(c.Param("good_id")); err != nil || goodId < 1 {
		aw.Error("无效的商品ID")
		return
	}

	// 获取商品对象
	good, err := goodModel.GetGoodByID(goodId)
	if err != nil {
		aw.Error(err.Error())
		return
	}
	// 获取前端上传的文件，字段名为"picture"
	file, err := c.FormFile("picture")
	if err != nil {
		aw.Error("文件上传失败: " + err.Error())
		return
	}

	// 检查上传的文件类型，确保是图片格式
	// 你可以根据需要进一步完善这个检查，比如限制文件的扩展名或 MIME 类型
	if !util.IsImage(file) {
		aw.Error("请上传有效的图片文件")
		return
	}

	// 获取文件扩展名
	ext := filepath.Ext(file.Filename)
	if ext == "" {
		aw.Error("文件扩展名无效")
		return
	}

	// 生成新的文件名，避免文件名重复
	timestamp := time.Now().Unix()
	newFileName := strconv.FormatInt(timestamp, 10) + ext

	// 文件存储路径（项目根目录下的 "cover" 文件夹）
	savePath := "../pictures/" + newFileName

	// 确保文件夹存在
	if err := os.MkdirAll("../pictures", os.ModePerm); err != nil {
		aw.Error("无法创建文件夹: " + err.Error())
		return
	}

	// 保存文件到服务器
	if err := c.SaveUploadedFile(file, savePath); err != nil {
		aw.Error("文件保存失败: " + err.Error())
		return
	}

	picture := &model.Pictures{
		URL:    savePath,
		GoodId: goodId,
	}
	picturesModel.CreatePicture(picture)

	goodL, err := goodModel.GetGoodByID(goodId)
	if err != nil {
		aw.Error(err.Error())
		return
	}
	if goodL.Cover != "" {
		good.IsEffective = true
	}
	// 更新数据库中的good记录
	if err := goodModel.UpdateGood(&good); err != nil {
		aw.Error(err.Error())
		return
	}

	// 返回成功
	aw.OK()
}

func (receiver GoodService) DeleteGood(c *gin.Context) {
	aw := app.NewWrapper(c)
	userClaim := util.ExtractUserClaims(c)
	goodId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		aw.Error(err.Error())
		return
	}
	if userClaim.RoleID == 1 {
		good, err := goodModel.GetGoodByID(goodId)
		if err != nil {
			aw.Error(err.Error())
			return
		}
		if good.Seller != userClaim.UserID {
			aw.Error("无权限删除")
			return
		}
	}

	picturesModel.DeletePictureByGoodId(goodId)
	goodModel.DeleteGood(goodId)
	aw.OK()

}

type GoodDTO struct {
	Total  int          `json:"total"`
	GoodsV []model.Good `json:"goods"`
}

func (receiver GoodService) GetAllGoods(c *gin.Context) {
	aw := app.NewWrapper(c)
	type GoodReq struct {
		IsEffective int    `json:"is_effective"`
		Page        int    `json:"page"`
		PageSize    int    `json:"page_size"`
		Search      string `json:"search"`
	}
	var req GoodReq
	if err := c.ShouldBind(&req); err != nil {
		aw.Error(err.Error())
		return
	}
	goods, tt, err := goodModel.GetAllGoods(req.IsEffective, req.Page, req.PageSize, req.Search)
	if err != nil {
		aw.Error(err.Error())
		return
	}
	goodLL := &GoodDTO{
		Total:  tt,
		GoodsV: goods,
	}
	aw.Success(goodLL)
}

func (receiver GoodService) GetMyGood(c *gin.Context) {
	aw := app.NewWrapper(c)
	userClaim := util.ExtractUserClaims(c)
	type GoodReq struct {
		IsEffective int    `json:"is_effective"`
		Page        int    `json:"page"`
		PageSize    int    `json:"page_size"`
		Search      string `json:"search"`
	}
	var req GoodReq
	if err := c.ShouldBind(&req); err != nil {
		aw.Error(err.Error())
		return
	}
	goods, tt, err := goodModel.GetMyGoods(req.IsEffective, req.Page, req.PageSize, userClaim.UserID, req.Search)
	if err != nil {
		aw.Error(err.Error())
		return
	}
	goodLL := &GoodDTO{
		Total:  tt,
		GoodsV: goods,
	}
	aw.Success(goodLL)

}
