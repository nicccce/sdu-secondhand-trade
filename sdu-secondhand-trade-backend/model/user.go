package model

import (
	"errors"
	"gorm.io/gorm"
	"sdu-secondhand-trade-backend/util"
	"time"
)

type User struct {
	ID        int    `json:"id"`
	RoleID    int    `json:"role_id"`
	Password  string `json:"password"`
	StudentID string `json:"student_id"`
	UserInfo
}
type UserInfo struct {
	Nickname     string    `json:"nickname"`
	IconUrl      string    `json:"icon_url"`
	Phone        string    `json:"phone"`
	Alipay       string    `json:"alipay"`
	Introduction string    `json:"introduction"`
	Campus       int       `json:"campus"`
	CreatedAt    time.Time `json:"created_at" binging:"-"`
	Gender       string    `json:"gender"`
}

type UserModel struct {
	AbstractModel
}

func (receiver UserModel) FindUserByID(id int) (*User, error) {
	var user User
	err := receiver.Tx.Take(&user, id).Error
	if errors.Is(gorm.ErrRecordNotFound, err) {
		return nil, nil
	}
	return &user, err
}
func (receiver UserModel) UpdateUser(user *User) {
	err := receiver.Tx.Save(user).Error
	util.ForwardOrPanic(err)
}
func (receiver UserModel) CreateUser(user *User) {
	err := receiver.Tx.Create(user).Error
	util.ForwardOrPanic(err)
}
func (receiver UserModel) FindAllUsers() []User {
	var users []User
	err := receiver.Tx.Order("id desc").Find(&users).Error
	util.ForwardOrPanic(err)
	return users
}
func (receiver UserModel) DeleteUser(user *User) {
	err := receiver.Tx.Unscoped().Delete(user).Error
	util.ForwardOrPanic(err)
}
func (receiver UserModel) FindUserByStudentID(studentID string) *User {
	var user User
	err := receiver.Tx.Where("student_id=?", studentID).Take(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil
	}
	util.ForwardOrPanic(err)
	return &user
}

func (receiver UserModel) GetAllUserGenders() []map[string]interface{} {
	var results []map[string]interface{}
	err := receiver.Tx.Model(&User{}).Select("gender, id").Order("id desc").Find(&results).Error
	util.ForwardOrPanic(err)
	return results
}
func (receiver UserModel) GetAllUserCampus() []map[string]interface{} {
	var results []map[string]interface{}
	err := receiver.Tx.Model(&Campus{}).Select("id, name").Order("id asc").Find(&results).Error
	util.ForwardOrPanic(err)
	return results
}
