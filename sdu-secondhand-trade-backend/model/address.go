package model

import (
	"sdu-secondhand-trade-backend/util"
)

type Address struct {
	ID      int    `json:"id"`
	UserID  int    `json:"user_id"`
	Address string `json:"address"`
}

type AddressModel struct {
	AbstractModel
}

func (receiver AddressModel) FindAddressesByUserID(userID int) []Address {
	var addresses []Address
	err := receiver.Tx.Where("user_id = ?", userID).Find(&addresses).Error
	util.ForwardOrPanic(err)
	return addresses
}
func (receiver AddressModel) FindAddressesByID(ID int) (*Address, error) {
	var address Address
	err := receiver.Tx.Where("id = ?", ID).Find(&address).Error
	util.ForwardOrPanic(err)
	return &address, err
}

func (receiver AddressModel) CreateAddress(address *Address) {
	err := receiver.Tx.Create(address).Error
	util.ForwardOrPanic(err)
}

func (receiver AddressModel) DeleteAddress(addressID int) {
	err := receiver.Tx.Where("id = ?", addressID).Delete(&Address{}).Error
	util.ForwardOrPanic(err)
}

func (receiver AddressModel) UpdateAddress(address *Address) {
	err := receiver.Tx.Save(address).Error
	util.ForwardOrPanic(err)
}
