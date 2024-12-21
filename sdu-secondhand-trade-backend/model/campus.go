package model

import "sdu-secondhand-trade-backend/util"

type Campus struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type CampusModel struct {
	AbstractModel
}

func (receiver CampusModel) FindCampusByID(ID int) (*Campus, error) {
	var campus Campus
	err := receiver.Tx.Where("id = ?", ID).Find(&campus).Error
	util.ForwardOrPanic(err)
	return &campus, err
}

func (receiver CampusModel) CreateCampus(campus *Campus) {
	err := receiver.Tx.Create(campus).Error
	util.ForwardOrPanic(err)
}

func (receiver CampusModel) DeleteCampus(campusID int) {
	err := receiver.Tx.Where("id = ?", campusID).Delete(&Campus{}).Error
	util.ForwardOrPanic(err)
}

func (receiver CampusModel) UpdateCampus(campus *Campus) {
	err := receiver.Tx.Save(campus).Error
	util.ForwardOrPanic(err)
}
