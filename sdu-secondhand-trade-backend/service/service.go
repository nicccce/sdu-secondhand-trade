package service

import "sdu-secondhand-trade-backend/model"

var userModel model.UserModel

func Setup() {
	userModel = model.UserModel{AbstractModel: model.AbstractModel{Tx: model.DB}}
}
