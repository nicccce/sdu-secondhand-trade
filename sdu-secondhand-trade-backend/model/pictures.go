package model

import (
	"sdu-secondhand-trade-backend/util"
)

type Pictures struct {
	ID     int    `json:"id"`
	URL    string `json:"url"`
	GoodId int    `json:"good_id"`
}

type PicturesModel struct {
	AbstractModel
}

// FindPicturesByGoodID 根据商品 ID 查找所有图片
func (receiver PicturesModel) FindPicturesByGoodID(goodID int) []Pictures {
	var pictures []Pictures
	err := receiver.Tx.Where("good_id = ?", goodID).Find(&pictures).Error
	util.ForwardOrPanic(err)
	return pictures
}

// FindPictureByID 根据图片 ID 查找单张图片
func (receiver PicturesModel) FindPictureByID(ID int) (*Pictures, error) {
	var picture Pictures
	err := receiver.Tx.Where("id = ?", ID).Find(&picture).Error
	util.ForwardOrPanic(err)
	return &picture, err
}

// CreatePicture 创建新图片
func (receiver PicturesModel) CreatePicture(picture *Pictures) {
	err := receiver.Tx.Create(picture).Error
	util.ForwardOrPanic(err)
}

// DeletePicture 删除图片
func (receiver PicturesModel) DeletePicture(pictureID int) {
	err := receiver.Tx.Where("id = ?", pictureID).Delete(&Pictures{}).Error
	util.ForwardOrPanic(err)
}

// UpdatePicture 更新图片
func (receiver PicturesModel) UpdatePicture(picture *Pictures) {
	err := receiver.Tx.Save(picture).Error
	util.ForwardOrPanic(err)
}

// DeletePictureByGoodId 删除图片
func (receiver PicturesModel) DeletePictureByGoodId(goodID int) {
	err := receiver.Tx.Where("good_id = ?", goodID).Delete(&Pictures{}).Error
	util.ForwardOrPanic(err)
}
