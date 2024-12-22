package service

import "sdu-secondhand-trade-backend/model"

var userModel model.UserModel
var addressModel model.AddressModel
var categoryModel model.CategoryModel
var bannerModel model.BannerModel
var goodModel model.GoodModel
var picturesModel model.PicturesModel
var campusModel model.CampusModel
var orderModel model.OrderModel
var problemModel model.ProblemModel

func Setup() {
	userModel = model.UserModel{AbstractModel: model.AbstractModel{Tx: model.DB}}
	addressModel = model.AddressModel{AbstractModel: model.AbstractModel{Tx: model.DB}}
	categoryModel = model.CategoryModel{AbstractModel: model.AbstractModel{Tx: model.DB}}
	bannerModel = model.BannerModel{AbstractModel: model.AbstractModel{Tx: model.DB}}
	goodModel = model.GoodModel{AbstractModel: model.AbstractModel{Tx: model.DB}}
	picturesModel = model.PicturesModel{
		AbstractModel: model.AbstractModel{Tx: model.DB},
	}
	campusModel = model.CampusModel{AbstractModel: model.AbstractModel{Tx: model.DB}}
	orderModel = model.OrderModel{AbstractModel: model.AbstractModel{Tx: model.DB}}
	picturesModel = model.PicturesModel{AbstractModel: model.AbstractModel{Tx: model.DB}}
	problemModel = model.ProblemModel{AbstractModel: model.AbstractModel{Tx: model.DB}}
}
