package service

import (
	"github.com/smartwalle/alipay/v3"
	"log"
	"sdu-secondhand-trade-backend/conf"
	"sdu-secondhand-trade-backend/model"
)

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
	var err error
	kAppId := conf.Conf.KAppId
	kPrivateKey := conf.Conf.KPrivateKey
	//支付宝提供了用于开发时测试的 sandbox 环境，对接的时候需要注意相关的 app id 和密钥是 sandbox 环境还是 production 环境的。如果是 sandbox 环境，本参数应该传 false，否则为 true。
	if client, err = alipay.New(kAppId, kPrivateKey, false); err != nil {
		log.Println("初始化支付宝失败", err)
		return
	}

	// 加载证书
	// 加载应用公钥证书
	if err = client.LoadAppCertPublicKeyFromFile("./cert/appPublicCert.crt"); err != nil {
		log.Println("加载证书发生错误", err)
		return
	}
	// 加载支付宝根证书
	if err = client.LoadAliPayRootCertFromFile("./cert/alipayRootCert.crt"); err != nil {
		log.Println("加载证书发生错误", err)
		return
	}
	// 加载支付宝公钥证书
	if err = client.LoadAlipayCertPublicKeyFromFile("./cert/alipayPublicCert.crt"); err != nil {
		log.Println("加载证书发生错误", err)
		return
	}
	//接口内容加密
	//	if err = client.SetEncryptKey("iotxR/d99T9Awom/UaSqiQ=="); err != nil {
	//log.Println("加载内容加密密钥发生错误", err)
	//return
}
