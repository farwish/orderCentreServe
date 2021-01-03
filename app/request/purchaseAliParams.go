package request

type PurchaseAliParams struct {
	AccountId	  uint	 `form:"account_id" json:"account_id" binding:"required"`
	ProductId	  uint	 `form:"product_id" json:"product_id" binding:"required"`
	TradeNo		  string `form:"trade_no" json:"trade_no" binding:"required"`
	TotalFee  	  uint	 `form:"total_fee" json:"total_fee" binding:"required"`

	Status		  uint	 `form:"status" json:"status" binding:"required"`
	PaymentType   uint	 `form:"payment_type" json:"payment_type" binding:"required"`
	PaymentStatus uint	 `form:"payment_status" json:"payment_status" binding:"required"`

	Source	  	  string `form:"source" json:"source" binding:"-"`
	ProductName	  string `form:"product_name" json:"product_name" binding:"-"`
}
