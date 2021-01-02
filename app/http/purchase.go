package http

import (
	"github.com/gin-gonic/gin"
	"github.com/gopher-lego/response"
)

type PurchaseAliParam struct {
	AccountId	  uint	 `form:"account_id" json:"account_id" binding:"required"`
	ProductId	  uint	 `form:"product_id" json:"product_id" binding:"required"`
	TradeNo		  string `form:"trade_no" json:"trade_no" binding:"required"`
	TotalFee  	  uint	 `form:"total_fee" json:"total_fee" binding:"required"`
	Status		  uint	 `form:"status" json:"status" binding:"required"`
	PaymentType   uint	 `form:"payment_type" json:"payment_type" binding:"required"`
	PaymentStatus uint	 `form:"payment_status" json:"payment_status" binding:"required"`

	ProductName	  string `form:"product_name" json:"product_name" binding:"-"`
}

// AliPay params means, https://docs.open.alipay.com/api_1/alipay.trade.page.pay
func PurchaseAli(c *gin.Context) {
	var param PurchaseAliParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failure(c, err.Error())
		return
	}

	// Persist into database

	// orderId := uuid.New().String()
}
