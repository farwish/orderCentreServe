package route

import (
	"github.com/gin-gonic/gin"

	"github.com/gopher-lego/ginger/app/http"
	"github.com/gopher-lego/ginger/app/middleware"
)

var corsMiddleware = middleware.CorsMiddleware()
var limiterMiddleware = middleware.LimiterMiddleware()

func Set(e *gin.Engine) {
	e.Use(corsMiddleware)
	e.Use(limiterMiddleware)

	// Top routes prefix
	top := e.Group("/api")

	heartbeat := top.Group("/ping")
	{
		heartbeat.GET("/", http.Pong)
	}

	// Create an order for different place
	purchase := top.Group("/purchase")
	{
		purchase.POST("/ali", http.PurchaseAli)
	}

	// Payment event from different place
	payment := top.Group("/payment")
	{
		payment.Any("/ali-return-url", http.PaymentAliReturnUrl)
		payment.Any("/ali-notify-url", http.PaymentAliNotifyUrl)

		payment.Any("/wx-native-scanned-notify", http.PaymentWxNativeScannedNotify)
		payment.Any("/wx-native-paid-notify", http.PaymentWxNativePaidNotify)
	}
}
