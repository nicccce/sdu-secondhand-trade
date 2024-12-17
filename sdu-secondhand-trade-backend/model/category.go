package model

type Category struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Picture      string `json:"picture"`
	Introduction string `json:"introduction"`
}

type CategoryModel struct {
	AbstractModel
}
