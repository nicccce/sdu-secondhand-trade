package model

import "time"

type Order struct {
	ID           int       `json:"id"`
	Status       *int      `json:"status"`
	GoodId       int       `json:"good_id"`
	BuyerId      int       `json:"buyer_id"`
	BuyerAddress string    `json:"buyer_address"`
	BuyerName    string    `json:"buyer_name"`
	BuyerContact string    `json:"buyer_contact"`
	Payment      int       `json:"payment"`
	CreatedAt    time.Time `json:"created_at" binging:"-"`
}

type OrderModel struct {
	AbstractModel
}

// CreateOrder 创建一个新订单
func (receiver *OrderModel) CreateOrder(order *Order) error {
	if err := receiver.Tx.Create(order).Error; err != nil {
		return err
	}
	return nil
}

// GetOrderByID 根据 ID 获取订单
func (receiver *OrderModel) GetOrderByID(id int) (*Order, error) {
	var order Order
	if err := receiver.Tx.Where("id = ?", id).First(&order).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

// UpdateOrder 更新订单信息
func (receiver *OrderModel) UpdateOrder(order *Order) error {
	if err := receiver.Tx.Save(order).Error; err != nil {
		return err
	}
	return nil
}

// DeleteOrder 根据订单 ID 删除订单
func (receiver *OrderModel) DeleteOrder(id int) error {
	if err := receiver.Tx.Where("id = ?", id).Delete(&Order{}).Error; err != nil {
		return err
	}
	return nil
}

// GetOrdersByBuyerIDAndStatus 分页查询，根据 buyer_id 和 status 获取订单列表
func (receiver *OrderModel) GetOrdersByBuyerIDAndStatus(buyerID, status, page, pageSize int) ([]Order, int, error) {
	var orders []Order
	var total int64
	offset := (page - 1) * pageSize // 计算偏移量

	// 构建查询条件
	query := receiver.Tx.Where("buyer_id = ? AND status = ?", buyerID, status).
		Limit(pageSize).Offset(offset) // 设置分页条件

	// 查询符合条件的总订单数
	if err := receiver.Tx.Model(&Order{}).Where("buyer_id = ? AND status = ?", buyerID, status).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	// 执行查询
	if err := query.Find(&orders).Error; err != nil {
		return nil, 0, err
	}

	return orders, int(total), nil
}

// GetOrdersBySellerIDAndStatus 分页查询，根据 seller_id 和 status 获取订单列表
func (receiver *OrderModel) GetOrdersBySellerIDAndStatus(goodID, status, page, pageSize int) ([]Order, int, error) {
	var orders []Order
	var total int64
	offset := (page - 1) * pageSize // 计算偏移量

	// 构建查询条件
	query := receiver.Tx.Where("good_id = ? AND status = ?", goodID, status).
		Limit(pageSize).Offset(offset) // 设置分页条件

	// 查询符合条件的总订单数
	if err := receiver.Tx.Model(&Order{}).Where("good_id = ? AND status = ?", goodID, status).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	// 执行查询
	if err := query.Find(&orders).Error; err != nil {
		return nil, 0, err
	}

	return orders, int(total), nil
}

// GetOrderByGoodID 根据 ID 获取订单
func (receiver *OrderModel) GetOrderByGoodID(id int) (*Order, error) {
	var order Order
	if err := receiver.Tx.Where("good_id = ?", id).First(&order).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

// GetOrdersByBuyerID 根据买家 ID 获取订单列表
func (receiver *OrderModel) GetOrdersByBuyerID(buyerID, page, pageSize int) ([]Order, int, error) {
	var orders []Order
	var total int64
	offset := (page - 1) * pageSize // 计算偏移量

	// 构建查询条件
	query := receiver.Tx.Where("buyer_id = ? ", buyerID).
		Limit(pageSize).Offset(offset) // 设置分页条件

	// 查询符合条件的总订单数
	if err := receiver.Tx.Model(&Order{}).Where("buyer_id = ?", buyerID).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	// 执行查询，获取订单列表
	if err := query.Find(&orders).Error; err != nil {
		return nil, 0, err
	}

	return orders, int(total), nil
}

// GetOrdersByGoodID 根据商品 ID 获取订单列表
func (receiver *OrderModel) GetOrdersByGoodID(goodID, page, pageSize int) ([]Order, int, error) {
	var orders []Order
	var total int64
	offset := (page - 1) * pageSize // 计算偏移量

	// 构建查询条件
	query := receiver.Tx.Where("good_id = ? ", goodID).
		Limit(pageSize).Offset(offset) // 设置分页条件

	// 查询符合条件的总订单数
	if err := receiver.Tx.Model(&Order{}).Where("good_id = ? ", goodID).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	// 执行查询，获取订单列表
	if err := query.Find(&orders).Error; err != nil {
		return nil, 0, err
	}

	return orders, int(total), nil
}
