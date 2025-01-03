package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/smartwalle/alipay/v3"
	"github.com/smartwalle/xid"
	"log"
	"net/http"
	"sdu-secondhand-trade-backend/app"
	"sdu-secondhand-trade-backend/conf"
	"sdu-secondhand-trade-backend/model"
	"strconv"
)

type AlipayService struct{}

var client *alipay.Client

func (receiver AlipayService) Pay(c *gin.Context) {
	aw := app.NewWrapper(c)
	var orderId int
	var err error
	if orderId, err = strconv.Atoi(c.Query("order_id")); err != nil {
		aw.Error(err.Error())
		return
	}
	redirect := c.Query("redirect")
	var order *model.Order
	var good model.Good
	if order, err = orderModel.GetOrderByID(orderId); err != nil {
		aw.Error(err.Error())
		return
	}
	var tradeNo = fmt.Sprintf("%d", xid.Next())
	good, err = goodModel.GetGoodByID(order.GoodId)
	if err != nil {
		aw.Error(err.Error())
		return
	}
	kServerDomain := conf.Conf.KServerDomain
	var p = alipay.TradePagePay{}
	p.NotifyURL = kServerDomain + "/pay/alipay/notify/" + strconv.Itoa(orderId)
	p.ReturnURL = redirect
	p.Subject = "支付测试:" + tradeNo
	p.OutTradeNo = tradeNo
	p.TotalAmount = strconv.Itoa(good.Price)
	p.ProductCode = "FAST_INSTANT_TRADE_PAY"

	url, _ := client.TradePagePay(p)
	c.Redirect(http.StatusTemporaryRedirect, url.String())
	//http.Redirect(writer, request, url.String(), http.StatusTemporaryRedirect)
}

func (receiver AlipayService) Callback(c *gin.Context) {
	// 解析请求参数
	if err := c.Request.ParseForm(); err != nil {
		log.Println("解析请求参数发生错误", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "解析请求参数发生错误"})
		return
	}

	// 获取通知参数
	outTradeNo := c.Request.Form.Get("out_trade_no")

	// 验证签名
	if err := client.VerifySign(c.Request.Form); err != nil {
		log.Println("回调验证签名发生错误", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "回调验证签名发生错误"})
		return
	}

	log.Println("回调验证签名通过")

	// 查询订单状态
	var p = alipay.TradeQuery{}
	p.OutTradeNo = outTradeNo

	rsp, err := client.TradeQuery(c, p)
	if err != nil {
		log.Printf("验证订单 %s 信息发生错误: %s", outTradeNo, err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("验证订单 %s 信息发生错误: %s", outTradeNo, err.Error())})
		return
	}

	if rsp.IsFailure() {
		log.Printf("验证订单 %s 信息发生错误: %s-%s", outTradeNo, rsp.Msg, rsp.SubMsg)
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("验证订单 %s 信息发生错误: %s-%s", outTradeNo, rsp.Msg, rsp.SubMsg)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("订单 %s 支付成功", outTradeNo)})

}
func (receiver AlipayService) Notify(c *gin.Context) {
	aw := app.NewWrapper(c)
	// 解析请求参数
	if err := c.Request.ParseForm(); err != nil {
		log.Println("解析请求参数发生错误", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "解析请求参数发生错误"})
		return
	}
	orderId, err := strconv.Atoi(c.Param("order_id"))
	if err != nil {
		aw.Error(err.Error())
		return
	}

	// 解析异步通知
	notification, err := client.DecodeNotification(c.Request.Form)
	if err != nil {
		log.Println("解析异步通知发生错误", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "解析异步通知发生错误"})
		return
	}

	log.Println("解析异步通知成功:", notification.NotifyId)

	// 查询订单状态
	var p = alipay.NewPayload("alipay.trade.query")
	p.AddBizField("out_trade_no", notification.OutTradeNo)

	var rsp *alipay.TradeQueryRsp
	if err := client.Request(c, p, &rsp); err != nil {
		log.Printf("异步通知验证订单 %s 信息发生错误: %s", notification.OutTradeNo, err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("异步通知验证订单 %s 信息发生错误: %s", notification.OutTradeNo, err.Error())})
		return
	}

	if rsp.IsFailure() {
		log.Printf("异步通知验证订单 %s 信息发生错误: %s-%s", notification.OutTradeNo, rsp.Msg, rsp.SubMsg)
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("异步通知验证订单 %s 信息发生错误: %s-%s", notification.OutTradeNo, rsp.Msg, rsp.SubMsg)})
		return
	}
	order, err := orderModel.GetOrderByID(orderId)
	if err != nil {
		aw.Error(err.Error())
		return
	}
	*order.Status = 2
	orderModel.UpdateOrder(order)

	good, err := goodModel.GetGoodByID(order.GoodId)
	if err != nil {
		aw.Error(err.Error())
		return
	}
	good.IsEffective = false
	goodModel.UpdateGood(&good)
	log.Printf("订单 %s 支付成功", notification.OutTradeNo)

	client.ACKNotification(c.Writer)

}
