package model

type Category struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Picture      string `json:"picture"`
	Introduction string `json:"introduction"`
	FatherID     int    `json:"-"`
}

type CategoryModel struct {
	AbstractModel
}

type FirstCategory struct {
	ID           int        `json:"id"`
	Name         string     `json:"name"`
	Picture      string     `json:"picture"`
	Introduction string     `json:"introduction"`
	Children     []Category `json:"children"`
}

// GetSubCategories 获取指定父分类的所有子分类
func (receiver *CategoryModel) GetSubCategories(fatherID int) ([]Category, error) {
	var subCategories []Category
	if err := receiver.Tx.Where("father_id = ?", fatherID).Find(&subCategories).Error; err != nil {
		return nil, err
	}
	return subCategories, nil
}

func (receiver *CategoryModel) GetCategoryById(categoryID int) (Category, error) {
	var subCategories Category
	if err := receiver.Tx.Where("id = ?", categoryID).Find(&subCategories).Error; err != nil {
		return Category{}, err
	}
	return subCategories, nil
}

// GetCategoriesWithSubCategories 获取所有一级分类及其对应的二级分类
func (receiver *CategoryModel) GetCategoriesWithSubCategories() ([]FirstCategory, error) {
	var firstCategories []Category
	var categories []FirstCategory
	// 获取所有一级分类（父ID为0）
	if err := receiver.Tx.Where("father_id = ?", 0).Find(&firstCategories).Error; err != nil {
		return nil, err
	}

	// 遍历一级分类并获取其对应的二级分类
	for i := range firstCategories {
		// 获取当前一级分类的所有子分类（二级分类）
		subCategories, err := receiver.GetSubCategories(firstCategories[i].ID)
		if err != nil {
			return nil, err
		}
		// 将二级分类添加到一级分类的 Children 字段
		categories = append(categories, FirstCategory{
			ID:           firstCategories[i].ID,
			Name:         firstCategories[i].Name,
			Picture:      firstCategories[i].Picture,
			Introduction: firstCategories[i].Introduction,
			Children:     subCategories,
		})
	}

	return categories, nil
}
