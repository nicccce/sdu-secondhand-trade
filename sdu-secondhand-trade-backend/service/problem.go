package service

import (
	"github.com/gin-gonic/gin"
	"sdu-secondhand-trade-backend/app"
	"sdu-secondhand-trade-backend/model"
	"sdu-secondhand-trade-backend/util"
)

type ProblemService struct{}

func (receiver ProblemService) CreateProblem(c *gin.Context) {
	aw := app.NewWrapper(c)
	userClaim := util.ExtractUserClaims(c)
	type updateReq struct {
		OrderId int    `json:"order_id"`
		Problem string `json:"problem"`
	}
	var req updateReq
	if err := c.ShouldBind(&req); err != nil {
		aw.Error(err.Error())
		return
	}
	newProblem := &model.Problem{
		OrderId: req.OrderId,
		Problem: req.Problem,
		Status:  0,
		UserId:  userClaim.UserID,
	}
	problemModel.CreateProblem(newProblem)
	aw.OK()
}

type ProblemVO struct {
	Total int             `json:"total"`
	List  []model.Problem `json:"problems"`
}

func (receiver ProblemService) GetAllProblem(c *gin.Context) {
	aw := app.NewWrapper(c)

	type ProblemReq struct {
		Page     int `json:"page"`
		PageSize int `json:"page_size"`
		Status   int `json:"status"`
	}

	var req ProblemReq
	if err := c.ShouldBind(&req); err != nil {
		aw.Error(err.Error())
		return
	}

	problem, tt, err := problemModel.GetAllProblems(req.Status, req.Page, req.PageSize)
	if err != nil {
		aw.Error(err.Error())
		return
	}
	problemVO := &ProblemVO{
		Total: tt,
		List:  problem,
	}

	aw.Success(problemVO)
}

func (receiver ProblemService) GetMyProblem(c *gin.Context) {
	aw := app.NewWrapper(c)
	userClaim := util.ExtractUserClaims(c)

	type ProblemReq struct {
		Page     int `json:"page"`
		PageSize int `json:"page_size"`
		Status   int `json:"status"`
	}

	var req ProblemReq
	if err := c.ShouldBind(&req); err != nil {
		aw.Error(err.Error())
		return
	}

	problem, tt, err := problemModel.GetMyProblem(req.Status, req.Page, req.PageSize, userClaim.UserID)
	if err != nil {
		aw.Error(err.Error())
		return
	}
	problemVO := &ProblemVO{
		Total: tt,
		List:  problem,
	}

	aw.Success(problemVO)
}

func (receiver ProblemService) UpdateProblem(c *gin.Context) {
	aw := app.NewWrapper(c)
	type updateReq struct {
		Id     int    `json:"id"`
		Status int    `json:"status"`
		Reply  string `json:"reply"`
	}
	var req updateReq
	if err := c.ShouldBind(&req); err != nil {
		aw.Error(err.Error())
		return
	}
	problem, err := problemModel.FindProblemByID(req.Id)
	if err != nil {
		aw.Error(err.Error())
		return
	}

	updated := &model.Problem{
		ID:      req.Id,
		Status:  req.Status,
		Reply:   req.Reply,
		OrderId: problem.OrderId,
		UserId:  problem.UserId,
		Problem: problem.Problem,
	}
	problemModel.UpdateProblem(updated)
	aw.OK()
}
