package model

import "log"

type Banner struct {
	ID      int    `json:"id"`
	ImgUrl  string `json:"img_url"`
	HrefUrl string `json:"href_url"`
	Type    string `json:"type"`
}

type BannerModel struct {
	AbstractModel
}

// getBanners 方法：从数据库中获取所有横幅数据
func (receiver BannerModel) GetBanners() ([]Banner, error) {
	var banners []Banner

	// 查询所有横幅数据
	if err := receiver.Tx.Find(&banners).Error; err != nil {
		// 记录错误日志并返回错误
		log.Println("Error fetching banners:", err)
		return nil, err
	}

	return banners, nil
}
