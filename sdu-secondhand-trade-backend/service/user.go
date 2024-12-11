package service

import (
	"github.com/gin-gonic/gin"
	"regexp"
	"sdu-secondhand-trade-backend/app"
	"sdu-secondhand-trade-backend/conf"
	"sdu-secondhand-trade-backend/model"
	"sdu-secondhand-trade-backend/util"
	"strconv"
)

type UserService struct {
}
type UserVO struct {
	ID        int    `json:"id"`
	RoleID    int    `json:"role_id"`
	StudentID string `json:"student_id"`
	model.UserInfo
	Token     *string         `json:"token,omitempty"`
	Addresses []model.Address `json:"addresses,omitempty"`
}

func (receiver UserService) TestGetJWT(c *gin.Context) {
	aw := app.NewWrapper(c)
	type getJWTReq struct {
		UserID    int    `form:"user_id" binding:"required"`
		JWTSecret string `form:"jwt_secret" binding:"required"`
	}
	var req getJWTReq
	if err := c.ShouldBind(&req); err != nil {
		aw.Error(err.Error())
		return
	}
	if req.JWTSecret != conf.Conf.JWTSecret {
		aw.Error("jwtSecret不正确")
		return
	}
	user, err := userModel.FindUserByID(req.UserID)
	if err != nil {
		aw.Error(err.Error())
		return
	}
	if user == nil {
		aw.Error("userID不存在")
		return
	}
	aw.Success(util.GenerateJWT(user.ID, user.RoleID))
}
func (receiver UserService) Login(c *gin.Context) {
	aw := app.NewWrapper(c)
	type loginReq struct {
		StudentID string `form:"student_id" json:"student_id" binding:"required"`
		Password  string `form:"password" json:"password" binding:"required"`
	}
	var req loginReq
	if err := c.ShouldBind(&req); err != nil {
		aw.Error(err.Error())
		return
	}
	user := userModel.FindUserByStudentID(req.StudentID)
	if user == nil {
		aw.Error("用户名不存在")
		return
	}
	if !util.CheckPassword(req.Password, user.Password) {
		aw.Error("用户名或密码错误")
		return
	}
	token := util.GenerateJWT(user.ID, user.RoleID)
	addresses := addressModel.FindAddressesByUserID(user.ID)

	userVO := UserVO{
		ID:        user.ID,
		RoleID:    user.RoleID,
		StudentID: user.StudentID,
		UserInfo:  user.UserInfo,
		Token:     &token,
		Addresses: addresses,
	}
	aw.Success(userVO)
}
func (receiver UserService) Register(c *gin.Context) {
	aw := app.NewWrapper(c)
	type registerReq struct {
		Password  string `form:"password" json:"password" binding:"required"`
		StudentID string `form:"student_id" json:"student_id" binding:"required"`
		Nickname  string `form:"nickname" json:"nickname" binding:"required"`
	}
	var req registerReq
	if err := c.ShouldBind(&req); err != nil {
		aw.Error(err.Error())
		return
	}
	pattern := `^202\d{9}$`
	regexp := regexp.MustCompile(pattern)
	if !regexp.MatchString(req.StudentID) {
		aw.Error("学号不合法")
		return
	}
	user := userModel.FindUserByStudentID(req.StudentID)
	if user != nil {
		aw.Error("账号已被注册")
		return
	}
	encryptPassword, err := util.EncryptPassword(req.Password)
	if err != nil {
		aw.Error(err.Error())
		return
	}
	user = &model.User{
		RoleID:    1,
		Password:  encryptPassword,
		StudentID: req.StudentID,
		UserInfo: model.UserInfo{
			Nickname: req.Nickname,
		},
	}
	userModel.CreateUser(user)
	aw.OK()
}
func (receiver UserService) Me(c *gin.Context) {
	aw := app.NewWrapper(c)
	userClaims := util.ExtractUserClaims(c)
	user, err := receiver.getUserInfo(userClaims.UserID)
	if err != nil {
		aw.Error(err.Error())
		return
	}
	aw.Success(user)
}
func (receiver UserService) GetUser(c *gin.Context) {
	aw := app.NewWrapper(c)
	userID, err := strconv.Atoi(c.Param("id"))
	user, err := receiver.getUserInfo(userID)
	if err != nil {
		aw.Error(err.Error())
		return
	}
	aw.Success(user)
}
func (receiver UserService) getUserInfo(id int) (UserVO, error) {
	user, err := userModel.FindUserByID(id)
	if err != nil {
		return UserVO{}, err
	}
	addresses := addressModel.FindAddressesByUserID(user.ID)
	return UserVO{
		ID:        user.ID,
		RoleID:    user.RoleID,
		StudentID: user.StudentID,
		UserInfo:  user.UserInfo,
		Addresses: addresses,
	}, nil
}

func (receiver UserService) GetAllUser(c *gin.Context) {
	aw := app.NewWrapper(c)
	users := userModel.FindAllUsers()

	// 将 users 转换为 gin.H 格式
	var userList []gin.H
	for _, user := range users {
		userData := gin.H{
			"id":         user.ID,
			"nickname":   user.UserInfo.Nickname,
			"student_id": user.StudentID,
			"role_id":    user.RoleID,
		}
		userList = append(userList, userData)
	}

	aw.Success(userList)
}

func (receiver UserService) UpdatePassword(c *gin.Context) {
	aw := app.NewWrapper(c)
	userClaim := util.ExtractUserClaims(c)
	type updateReq struct {
		oldPassword string `json:"old_password" binding:"required"`
		newPassword string `json:"new_password" binding:"required"`
	}
	var req updateReq
	if err := c.ShouldBind(&req); err != nil {
		aw.Error(err.Error())
	}
	user, err := userModel.FindUserByID(userClaim.UserID)
	if err != nil {
		aw.Error(err.Error())
		return
	}
	if user == nil {
		aw.Error("用户名不存在")
		return
	}
	if !util.CheckPassword(req.oldPassword, user.Password) {
		aw.Error("旧密码错误")
		return
	}
	encryptPassword, err := util.EncryptPassword(req.newPassword)
	if err != nil {
		aw.Error(err.Error())
		return
	}
	user = &model.User{
		RoleID:    user.RoleID,
		Password:  encryptPassword,
		StudentID: user.StudentID,
		UserInfo:  user.UserInfo,
	}
	userModel.UpdateUser(user)
	aw.OK()
}
func (receiver UserService) CreateAddress(c *gin.Context) {
	aw := app.NewWrapper(c)
	userClaim := util.ExtractUserClaims(c)
	type updateReq struct {
		address string `form:"address" json:"address" binding:"required"`
	}
	var req updateReq
	if err := c.ShouldBind(&req); err != nil {
		aw.Error(err.Error())
	}
	user, err := userModel.FindUserByID(userClaim.UserID)
	if err != nil {
		aw.Error(err.Error())
		return
	}
	if user == nil {
		aw.Error("用户名不存在")
		return
	}
	address := req.address
	newAddress := &model.Address{
		Address: address,
		UserID:  userClaim.UserID,
	}
	addressModel.CreateAddress(newAddress)
	aw.OK()
}

func (receiver UserService) UpdateAddress(c *gin.Context) {
	aw := app.NewWrapper(c)
	userClaim := util.ExtractUserClaims(c)
	type updateReq struct {
		address string `form:"address" json:"address" binding:"required"`
	}
	var req updateReq
	if err := c.ShouldBind(&req); err != nil {
		aw.Error(err.Error())
	}
	id, err := strconv.Atoi(c.Param("id"))
	address, err := addressModel.FindAddressesByID(id)
	if err != nil {
		aw.Error(err.Error())
		return
	}
	if address == nil {
		aw.Error("地址不存在")
		return
	}
	address = &model.Address{
		Address: req.address,
		UserID:  userClaim.UserID,
	}
	addressModel.UpdateAddress(address)
	aw.OK()
}

func (receiver UserService) DeleteAddress(c *gin.Context) {
	aw := app.NewWrapper(c)
	type updateReq struct {
		address string `form:"address" json:"address" binding:"required"`
	}
	var req updateReq
	if err := c.ShouldBind(&req); err != nil {
		aw.Error(err.Error())
	}
	id, err := strconv.Atoi(c.Param("id"))
	address, err := addressModel.FindAddressesByID(id)
	if err != nil {
		aw.Error(err.Error())
		return
	}
	if address == nil {
		aw.Error("地址不存在")
		return
	}
	addressModel.DeleteAddress(id)
	aw.OK()
}

func (receiver UserService) UpdateUser(c *gin.Context) {
	aw := app.NewWrapper(c)
	userClaim := util.ExtractUserClaims(c)
	type updateReq struct {
		model.UserInfo
	}
	var req updateReq
	if err := c.ShouldBind(&req); err != nil {
		aw.Error(err.Error())
	}
	user, err := userModel.FindUserByID(userClaim.UserID)
	if err != nil {
		aw.Error(err.Error())
		return
	}
	if user == nil {
		aw.Error("用户名不存在")
		return
	}
	user = &model.User{
		RoleID:    user.RoleID,
		Password:  user.Password,
		StudentID: user.StudentID,
		UserInfo:  req.UserInfo,
	}
	userModel.UpdateUser(user)
	aw.OK()
}
