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
