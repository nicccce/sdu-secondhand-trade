package model

type Address struct {
	ID      int    `json:"id"`
	UserID  int    `json:"user_id"`
	Address string `json:"address"`
}

type AddressModel struct {
	AbstractModel
}
