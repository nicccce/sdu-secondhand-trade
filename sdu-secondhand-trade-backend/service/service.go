package service

import "sdu-secondhand-trade-backend/model"

var userModel model.UserModel
var addressModel model.AddressModel

func Setup() {
	userModel = model.UserModel{AbstractModel: model.AbstractModel{Tx: model.DB}}
	addressModel = model.AddressModel{AbstractModel: model.AbstractModel{Tx: model.DB}}
}
