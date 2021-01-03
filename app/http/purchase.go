package http

import (
	"github.com/gin-gonic/gin"
	"github.com/gopher-lego/ginger/app/repository"
	"github.com/gopher-lego/ginger/app/request"
	"github.com/gopher-lego/response"
)

// AliPay params means, https://docs.open.alipay.com/api_1/alipay.trade.page.pay
func PurchaseAliPC(c *gin.Context) {
	var params request.PurchaseAliParams
	if validateErr := c.ShouldBind(&params); validateErr != nil {
		response.Failure(c, validateErr.Error())
		return
	}

	// 1.Persist our order into database
	var result = repository.PurchaseAliPCOurSystemCreate(params)
	if result.Error != nil {
		response.Failure(c, "Create order failed")
		return
	}

	// 2.Create ali pay order to get payUrl
	var payUrl, tradeErr = repository.PurchaseAliPCTheirSystemCreate(params)
	if tradeErr != nil {
		response.Failure(c, tradeErr.Error())
		return
	}

	response.Success(c, payUrl)
}
