package model

import (
	"errors"
	"gorm.io/gorm"
	"sdu-secondhand-trade-backend/util"
	"time"
)

type Good struct {
	GoodBrief
	IsEffective       bool      `json:"is_effective" gorm:"default:false"`
	Brand             string    `json:"brand"`
	Status            string    `json:"status"`
	FullName          string    `json:"full_name"`
	Specification     string    `json:"specifications"`
	TransactionMethod int       `json:"transaction_method"`
	Detail            string    `json:"detail"`
	CreatedAt         time.Time `json:"created_at" binging:"-"`
	Seller            int       `json:"seller"`
	SellerAddress     string    `json:"seller_address"`
}
type GoodBrief struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	Campus      int    `json:"campus"` //campus_id
	Cover       string `json:"cover"`
	Category    int    `json:"category"`
}

type GoodModel struct {
	AbstractModel
}

// GetAllUnfinishedGoods 查询指定分类下所有未完成的商品
func (receiver GoodModel) GetAllUnfinishedGoods() ([]GoodBrief, error) {
	var goods []Good
	var goodsVO []GoodBrief
	if err := receiver.Tx.Where("is_effective = ?", true).Find(&goods).Error; err != nil {
		return nil, err
	}
	for _, good := range goods {
		goodsVO = append(goodsVO, good.GoodBrief)
	}
	return goodsVO, nil
}

// GetLatestGoods 获取最新的商品，按 CreatedAt 降序排序，返回前十个
func (receiver GoodModel) GetLatestGoods() ([]GoodBrief, error) {
	var goods []Good
	var goodsVO []GoodBrief
	if err := receiver.Tx.Where("is_effective = ?", true).Order("created_at desc").Limit(10).Find(&goods).Error; err != nil {
		return nil, err
	}
	for _, good := range goods {
		goodsVO = append(goodsVO, good.GoodBrief)
	}
	return goodsVO, nil
}

// GetGoodsByCampus 获取指定校区的商品，返回前十个
func (receiver GoodModel) GetGoodsByCampus(campusId int) ([]GoodBrief, error) {
	var goods []Good
	var goodsVO []GoodBrief
	if err := receiver.Tx.Where("is_effective = ?", true).Where("campus = ?", campusId).Limit(10).Find(&goods).Error; err != nil {
		return nil, err
	}
	for _, good := range goods {
		goodsVO = append(goodsVO, good.GoodBrief)
	}
	return goodsVO, nil
}

// GetGoodByID 根据商品 ID 获取商品详情
func (receiver GoodModel) GetGoodByID(id int) (Good, error) {
	var good Good

	if err := receiver.Tx.Where("id = ?", id).First(&good).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return Good{
				IsEffective: false,
			}, nil
		}
		return Good{}, err
	}

	return good, nil
}

// GetUnfinishedGoodsPaginated 根据分类ID、页码、页面大小以及排序字段查询商品，返回指定分页的商品列表
func (receiver GoodModel) GetUnfinishedGoodsPaginated(categoryId int, page int, pageSize int, sortField string, campusId int) ([]GoodBrief, error) {
	var goods []Good
	var goodsVO []GoodBrief
	query := receiver.Tx.Where("is_effective = ?", 1).Where("category = ?", categoryId)

	if sortField == "created_at" {
		query = query.Order("created_at desc")
	} else if sortField == "campus" {
		query = query.Order("campus asc").Where("campus = ?", campusId)
	}

	if err := query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&goods).Error; err != nil {
		return nil, err
	}

	for _, good := range goods {
		goodsVO = append(goodsVO, good.GoodBrief)
	}

	return goodsVO, nil
}

func (receiver GoodModel) UpdateGood(good *Good) error {
	// 检查商品是否存在
	var existingGood Good
	if err := receiver.Tx.Where("id = ?", good.ID).First(&existingGood).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("商品不存在")
		}
		return err
	}

	if err := receiver.Tx.Save(good).Error; err != nil {
		return err
	}

	return nil
}

// GetGoodBySellerID 根据卖家 ID 获取商品详情
func (receiver GoodModel) GetGoodBySellerID(id int) ([]Good, error) {
	var good []Good

	if err := receiver.Tx.Where("seller = ?", id).Find(&good).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return good, nil
}

func (receiver GoodModel) CreateGood(good *Good) {
	err := receiver.Tx.Create(good).Error
	util.ForwardOrPanic(err)
}

func (receiver GoodModel) DeleteGood(ID int) {
	err := receiver.Tx.Where("id = ?", ID).Delete(&Good{}).Error
	util.ForwardOrPanic(err)
}

/*// GetAllGoods 获取所有商品（包括有效的商品），支持根据search进行模糊查询
func (receiver GoodModel) GetAllGoods(isEffective int, page int, pageSize int, search string) ([]Good, int, error) {
	var goods []Good
	var total int64
	// 如果传入的search为空，直接进行分页查询
	if search == "" {
		if isEffective == -1 {
			if err := receiver.Tx.Offset((page - 1) * pageSize).Limit(pageSize).Find(&goods).Error; err != nil {
				return nil, 0, err
			}
			// 查询符合条件的总订单数
			if err := receiver.Tx.Model(&Order{}).Count(&total).Error; err != nil {
				return nil, 0, err
			}
			return goods, int(total), nil
		} else {
			query := receiver.Tx.Where("is_effective = ?", isEffective)
			if err := query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&goods).Error; err != nil {
				return nil, 0, err
			}
			// 查询符合条件的总订单数
			if err := query.Model(&Order{}).Count(&total).Error; err != nil {
				return nil, 0, err
			}
			return goods, int(total), nil
		}

	}

	// 否则进行模糊查询

	// 模糊查询所有字段
	query := receiver.Tx.Where(
		"full_name LIKE ? OR brand LIKE ? OR status LIKE ? OR specification LIKE ? OR detail LIKE ? OR name LIKE ? OR description LIKE ?",
		"%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%",
	)
	if isEffective == -1 {
		if err := query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&goods).Error; err != nil {
			return nil, 0, err
		}
		// 查询符合条件的总订单数
		if err := query.Model(&Order{}).Count(&total).Error; err != nil {
			return nil, 0, err
		}

		return goods, int(total), nil
	}

	query = query.Where("is_effective = ?", isEffective)

	if err := query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&goods).Error; err != nil {
		return nil, 0, err
	}
	// 查询符合条件的总订单数
	if err := query.Model(&Order{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return goods, int(total), nil
}
*/

// GetAllGoods 获取所有商品（包括有效的商品），支持根据search进行模糊查询
func (receiver GoodModel) GetAllGoods(isEffective int, page int, pageSize int, search string) ([]Good, int, error) {
	var goods []Good
	var total int64

	// 构建查询条件
	query := receiver.Tx
	if search != "" {
		// 模糊查询所有字段
		query = query.Where(
			"full_name LIKE ? OR brand LIKE ? OR status LIKE ? OR specification LIKE ? OR detail LIKE ? OR name LIKE ? OR description LIKE ?",
			"%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%",
		)
	}

	// 如果 isEffective 不是 -1，则添加 isEffective 条件
	if isEffective != -1 {
		query = query.Where("is_effective = ?", isEffective)
	}
	// 查询符合条件的总记录数
	if err := query.Model(&Good{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	if err := query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&goods).Error; err != nil {
		return nil, 0, err
	}

	return goods, int(total), nil
}

// GetMyGoods 根据分类ID、页码、页面大小以及排序字段查询商品，返回指定分页的商品列表
func (receiver GoodModel) GetMyGoods(isEffective int, page int, pageSize int, userID int, search string) ([]Good, int, error) {
	var goods []Good
	var total int64

	// 构建查询条件
	query := receiver.Tx.Where("seller = ?", userID)
	if search != "" {
		// 模糊查询所有字段
		query = query.Where(
			"full_name LIKE ? OR brand LIKE ? OR status LIKE ? OR specification LIKE ? OR detail LIKE ? OR name LIKE ? OR description LIKE ?",
			"%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%",
		)
	}

	// 如果 isEffective 不是 -1，则添加 isEffective 条件
	if isEffective != -1 {
		query = query.Where("is_effective = ?", isEffective)
	}
	// 查询符合条件的总记录数
	if err := query.Model(&Good{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	if err := query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&goods).Error; err != nil {
		return nil, 0, err
	}

	return goods, int(total), nil
}

func (receiver GoodModel) GetSearchGoods(search string) ([]Good, error) {
	var goods []Good

	// 构建查询条件
	query := receiver.Tx
	if search != "" {
		// 模糊查询所有字段
		query = query.Where(
			"full_name LIKE ? OR brand LIKE ? OR status LIKE ? OR specification LIKE ? OR detail LIKE ? OR name LIKE ? OR description LIKE ?",
			"%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%",
		)
	}

	query = query.Where("is_effective = ?", true)

	if err := query.Find(&goods).Error; err != nil {
		return nil, err
	}

	return goods, nil
}
