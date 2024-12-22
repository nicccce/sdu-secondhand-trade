package service

import (
	"github.com/gin-gonic/gin"
	"sdu-secondhand-trade-backend/app"
	"sdu-secondhand-trade-backend/model"
	"sdu-secondhand-trade-backend/util"
	"strconv"
	"sync"
	"time"
)

type OrderService struct{}

var orderLocks sync.Map // 用于保存每个订单的锁，防止多个操作修改同一订单

// 生成订单并启动协程管理订单状态
func (receiver *OrderService) CreateOrder(c *gin.Context) {
	aw := app.NewWrapper(c)
	userClaim := util.ExtractUserClaims(c)
	user, err := userModel.FindUserByID(userClaim.UserID)
	if err != nil {
		aw.Error(err.Error())
		return
	}

	var req struct {
		GoodID int `json:"good_id" binding:"required"`
	}

	// 绑定请求参数
	if err := c.ShouldBind(&req); err != nil {
		aw.Error("请求参数错误: " + err.Error())
		return
	}

	// 查找商品
	good, err := goodModel.GetGoodByID(req.GoodID)
	if err != nil {
		aw.Error("获取商品失败: " + err.Error())
		return
	}

	// 如果商品已经不可用，则返回错误
	if !good.IsEffective {
		aw.Error("商品已下架或不可用")
		return
	}

	// 创建订单
	order := &model.Order{
		BuyerId:   userClaim.UserID,
		BuyerName: user.Nickname,
		GoodId:    req.GoodID,
		Status:    new(int), // 状态初始化为0（待提交）
		CreatedAt: time.Now(),
	}

	// 将订单状态设为待提交
	*order.Status = 0

	// 创建订单
	if err := orderModel.CreateOrder(order); err != nil {
		aw.Error("创建订单失败: " + err.Error())
		return
	}

	// 在创建订单时，设置商品为不可用
	good.IsEffective = false
	if err := goodModel.UpdateGood(&good); err != nil {
		aw.Error("更新商品状态失败: " + err.Error())
		return
	}

	// 启动协程管理订单的状态
	go receiver.manageOrderStatus(order.ID, req.GoodID)

	// 返回订单ID
	aw.Success(order.ID)
}

// 管理订单状态
func (receiver *OrderService) manageOrderStatus(orderID int, goodID int) {
	// 为了防止并发访问导致的竞态条件，我们使用一个锁来确保每个订单的状态管理是独立的
	lock, _ := orderLocks.LoadOrStore(orderID, &sync.Mutex{})
	orderLock := lock.(*sync.Mutex)

	// 锁住该订单，确保状态变更时不会被其他操作干扰
	orderLock.Lock()
	defer orderLock.Unlock()

	// 设置一个 10 分钟的超时，超时后自动取消订单
	select {
	case <-time.After(10 * time.Minute):
		// 超过 10 分钟后自动取消订单
		order, err := orderModel.GetOrderByID(orderID)
		if err != nil {
			return
		}
		// 如果订单的状态还没被改变，设置为已取消（状态为4）
		if *order.Status == 0 { // 只有在订单状态仍为0（待提交）时才会取消
			order.Status = new(int)
			*order.Status = 4 // 已取消
			if err := orderModel.UpdateOrder(order); err != nil {
				// 错误处理：更新订单失败
				return
			}
		}

		// 超时后，恢复商品为有效
		good, err := goodModel.GetGoodByID(goodID)
		if err != nil {
			return
		}
		good.IsEffective = true
		if err := goodModel.UpdateGood(&good); err != nil {
			// 错误处理：恢复商品状态失败
			return
		}
	}
}

// 提交订单，状态改为待支付
func (receiver *OrderService) SubmitOrder(c *gin.Context) {
	aw := app.NewWrapper(c)
	var req struct {
		model.Order
	}

	// 绑定请求参数
	if err := c.ShouldBind(&req); err != nil {
		aw.Error("请求参数错误: " + err.Error())
		return
	}

	// 查找订单
	order, err := orderModel.GetOrderByID(req.ID)
	if err != nil {
		aw.Error("获取订单失败: " + err.Error())
		return
	}

	// 如果订单的状态已经不是待提交（0），则无法提交
	if *order.Status != 0 {
		aw.Error("订单状态不正确，无法提交")
		return
	}

	// 提交订单，更新订单状态为待支付（1）
	order.Status = new(int)
	*order.Status = 1
	if order.Payment == 2 || order.Payment == 3 {
		*order.Status = 2
	}
	order.Payment = req.Payment
	order.BuyerId = req.BuyerId
	order.BuyerName = req.BuyerName
	order.BuyerAddress = req.BuyerAddress
	order.BuyerContact = req.BuyerContact

	// 更新订单信息
	if err := orderModel.UpdateOrder(order); err != nil {
		aw.Error("提交订单失败: " + err.Error())
		return
	}

	// 提交时，恢复商品为有效
	good, err := goodModel.GetGoodByID(order.GoodId)
	if err != nil {
		aw.Error("获取商品失败: " + err.Error())
		return
	}
	good.IsEffective = true
	if err := goodModel.UpdateGood(&good); err != nil {
		aw.Error("更新商品状态失败: " + err.Error())
		return
	}

	aw.Success("订单提交成功，待支付")
}

// 取消订单
func (receiver *OrderService) CancelOrder(c *gin.Context) {
	aw := app.NewWrapper(c)
	var orderId int
	var err error
	if orderId, err = strconv.Atoi(c.Param("id")); err != nil {
		aw.Error(err.Error())
		return
	}

	// 查找订单
	order, err := orderModel.GetOrderByID(orderId)
	if err != nil {
		aw.Error("获取订单失败: " + err.Error())
		return
	}

	// 检查订单状态是否为待提交（0），只有待提交的订单才能取消
	if *order.Status != 0 {
		aw.Error("订单状态不正确，无法取消")
		return
	}

	// 获取当前时间，判断订单是否超过 10 分钟
	if time.Since(order.CreatedAt) > 10*time.Minute {
		aw.Error("订单超时，无法取消")
		return
	}

	// 获取当前用户的身份（买家或卖家）
	userClaim := util.ExtractUserClaims(c)
	if userClaim == nil {
		aw.Error("用户未登录")
		return
	}

	GG, _ := goodModel.GetGoodByID(order.GoodId)
	// 判断是否为卖家或买家
	if order.BuyerId != userClaim.UserID && GG.Seller != userClaim.UserID {
		aw.Error("只有卖家或买家才能取消订单")
		return
	}

	order.Status = new(int)
	*order.Status = 3

	// 更新订单
	if err := orderModel.UpdateOrder(order); err != nil {
		aw.Error("更新订单状态失败: " + err.Error())
		return
	}

	// 恢复商品为有效
	good, err := goodModel.GetGoodByID(order.GoodId)
	if err != nil {
		aw.Error("获取商品失败: " + err.Error())
		return
	}

	good.IsEffective = true
	if err := goodModel.UpdateGood(&good); err != nil {
		aw.Error("恢复商品状态失败: " + err.Error())
		return
	}

	aw.Success("订单已取消")
}

// 获取订单信息
func (receiver *OrderService) GetOrder(c *gin.Context) {
	aw := app.NewWrapper(c)

	// 获取订单 ID
	orderID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		aw.Error("非法订单")
		return
	}

	// 查询订单
	order, err := orderModel.GetOrderByID(orderID)
	if err != nil {
		aw.Error("获取订单失败: " + err.Error())
		return
	}

	// 获取当前用户的身份（买家或卖家）
	userClaim := util.ExtractUserClaims(c)
	if userClaim == nil {
		aw.Error("用户未登录")
		return
	}

	good, err := goodModel.GetGoodByID(order.GoodId)

	// 判断当前用户是否为订单的买家或卖家
	if order.BuyerId != userClaim.UserID && good.Seller != userClaim.UserID {
		aw.Error("您没有权限查看此订单")
		return
	}

	// 返回订单信息
	aw.Success(order)
}

type OrderVo struct {
	Total int           `json:"total"`
	LL    []model.Order `json:"order"`
}

func (receiver *OrderService) GetBuyerOrder(c *gin.Context) {
	aw := app.NewWrapper(c)
	userClaim := util.ExtractUserClaims(c)
	type OrderReq struct {
		Status   *int `json:"status" binding:"required"`
		Page     int  `json:"page" binding:"required,min=1"`
		PageSize int  `json:"page_size" binding:"required,min=1"`
	}
	var req OrderReq
	if err := c.ShouldBind(&req); err != nil {
		aw.Error(err.Error())
		return
	}

	if *req.Status == -1 {
		order, tt, err := orderModel.GetOrdersByBuyerID(userClaim.UserID, req.Page, req.PageSize)
		if err != nil {
			aw.Error(err.Error())
			return
		}
		orderVo := &OrderVo{
			Total: tt,
			LL:    order,
		}
		aw.Success(orderVo)
		return
	}

	order, tt, err := orderModel.GetOrdersByBuyerIDAndStatus(userClaim.UserID, *req.Status, req.Page, req.PageSize)
	if err != nil {
		aw.Error(err.Error())
		return
	}
	orderVo := &OrderVo{
		Total: tt,
		LL:    order,
	}
	// 返回订单信息
	aw.Success(orderVo)
}

func (receiver *OrderService) GetSellerOrder(c *gin.Context) {
	aw := app.NewWrapper(c)
	userClaim := util.ExtractUserClaims(c)
	type OrderReq struct {
		Status   *int `json:"status" binding:"required"`
		Page     int  `json:"page" binding:"required,min=1"`
		PageSize int  `json:"page_size" binding:"required,min=1"`
	}
	var req OrderReq
	if err := c.ShouldBind(&req); err != nil {
		aw.Error(err.Error())
		return
	}
	goods, err := goodModel.GetGoodBySellerID(userClaim.UserID)
	if err != nil {
		aw.Error(err.Error())
		return
	}
	var order []model.Order
	var orders []model.Order
	var tt int
	var sum int
	if *req.Status == -1 {
		for _, good := range goods {
			order, tt, err = orderModel.GetOrdersByGoodID(good.ID, req.Page, req.PageSize)
			if err != nil {
				aw.Error(err.Error())
				return
			}
			sum += tt
			orders = append(orders, order...)
		}
	} else {
		for _, good := range goods {
			order, tt, err = orderModel.GetOrdersBySellerIDAndStatus(good.ID, *req.Status, req.Page, req.PageSize)
			if err != nil {
				aw.Error(err.Error())
				return
			}
			sum += tt
			orders = append(orders, order...)
		}
	}

	orderVo := &OrderVo{
		Total: sum,
		LL:    orders,
	}
	// 返回订单信息
	aw.Success(orderVo)
}
