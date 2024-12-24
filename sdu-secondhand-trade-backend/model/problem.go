package model

import "sdu-secondhand-trade-backend/util"

type Problem struct {
	ID      int    `json:"id"`
	OrderId int    `json:"order_id"`
	UserId  int    `json:"user_id"`
	Problem string `json:"problem"`
	Status  int    `json:"status"`
	Reply   string `json:"reply"`
}

type ProblemModel struct {
	AbstractModel
}

// FindProblemByOrderID 根据Order ID 查找所有problem
func (receiver ProblemModel) FindProblemByOrderID(orderID int) []Problem {
	var problem []Problem
	err := receiver.Tx.Where("order_id = ?", orderID).Find(&problem).Error
	util.ForwardOrPanic(err)
	return problem
}

// FindProblemByID 根据 ID 查找problem
func (receiver ProblemModel) FindProblemByID(ID int) (*Problem, error) {
	var problem Problem
	err := receiver.Tx.Where("id = ?", ID).Find(&problem).Error
	util.ForwardOrPanic(err)
	return &problem, err
}

// CreateProblem 创建新problem
func (receiver ProblemModel) CreateProblem(problem *Problem) {
	err := receiver.Tx.Create(problem).Error
	util.ForwardOrPanic(err)
}

// DeleteProblem 删除Problem
func (receiver ProblemModel) DeleteProblem(problemID int) {
	err := receiver.Tx.Where("id = ?", problemID).Delete(&Problem{}).Error
	util.ForwardOrPanic(err)
}

// UpdateProblem 更新Problem
func (receiver ProblemModel) UpdateProblem(problem *Problem) {
	err := receiver.Tx.Save(problem).Error
	util.ForwardOrPanic(err)
}

func (receiver ProblemModel) GetAllProblems(status int, page int, pageSize int) ([]Problem, int, error) {
	var problems []Problem
	var total int64

	// 构建查询条件
	query := receiver.Tx

	// 如果 status 不是 -1，则添加 status 条件
	if status != -1 {
		query = query.Where("status = ?", status)
	}

	// 查询符合条件的总记录数
	if err := query.Model(&Problem{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	// 分页查询
	if err := query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&problems).Error; err != nil {
		return nil, 0, err
	}

	return problems, int(total), nil
}

func (receiver ProblemModel) GetMyProblem(status int, page int, pageSize int, userId int) ([]Problem, int, error) {
	var problems []Problem
	var total int64

	// 构建查询条件
	query := receiver.Tx.Where("user_id = ?", userId)

	// 如果 status 不是 -1，则添加 status 条件
	if status != -1 {
		query = query.Where("status = ?", status)
	}

	// 查询符合条件的总记录数
	if err := query.Model(&Problem{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	// 分页查询
	if err := query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&problems).Error; err != nil {
		return nil, 0, err
	}

	return problems, int(total), nil
}
