package repository

import (
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"github.com/gopher-lego/ginger/app/model"
	"github.com/gopher-lego/ginger/app/request"
	"github.com/gopher-lego/ginger/config"
	"github.com/smartwalle/alipay/v3"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"time"
)

// Our system order
func PurchaseAliPCOurSystemCreate(params request.PurchaseAliParams) *gorm.DB {
	order := model.Order{
		OrderId:       uuid.New().String(),

		AccountId:     params.AccountId,
		ProductId:     params.ProductId,
		TradeNo:       params.TradeNo,
		TotalFee:      params.TotalFee,

		Status:        0,
		PaymentType:   0,
		PaymentStatus: 0,
		PaymentAt:     sql.NullTime{},

		Source:        params.Source,
		ProductName:   params.ProductName,

		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     gorm.DeletedAt{},
	}

	db := config.MySqlInit()

	return db.Create(order)
}

// Ali system order
// @return payUrl, error
func PurchaseAliPCTheirSystemCreate(params request.PurchaseAliParams) (string, error) {
	// Instance
	var client, clientErr = alipay.New(viper.GetString("payment.aliPay.appId"), viper.GetString("payment.aliPay.privateKey"), viper.GetBool("payment.aliPay.isProduction"))
	if clientErr != nil {
		return "", errors.New("ali pay sdk new instance error")
	}
	if client.LoadAliPayPublicKey(viper.GetString("payment.aliPay.aliPublicKey")) != nil {
		return "", errors.New("ali pay sdk load public key error")
	}

	// Params
	var tradeParams alipay.TradePagePay
	tradeParams.ReturnURL = viper.GetString("payment.aliPay.returnUrl")
	tradeParams.NotifyURL = viper.GetString("payment.aliPay.notifyUrl")
	// biz content, belows four is required
	tradeParams.OutTradeNo = params.TradeNo
	tradeParams.TotalAmount = cast.ToString(params.TotalFee / 100) // 支付宝:元
	tradeParams.Subject = params.ProductName
	tradeParams.ProductCode = "FAST_INSTANT_TRADE_PAY"

	// Create ali order
	var payUrl, tradeErr = client.TradePagePay(tradeParams)
	if tradeErr != nil {
		return "", errors.New("ali pay create order failure")
	}

	// URL for pay
	return payUrl.String(), nil
}
