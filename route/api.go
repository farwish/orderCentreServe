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
		purchase.POST("/ali-pc", http.PurchaseAliPC)
	}

	// Payment event from different place
	payment := top.Group("/payment")
	{
		// To allow any request(GET/POST only), NOT Any to exclude others(HEAD/PUT..)
		payment.GET("/ali-return-url", http.PaymentAliReturnUrl)
		payment.POST("/ali-return-url", http.PaymentAliReturnUrl)

		payment.GET("/ali-notify-url", http.PaymentAliNotifyUrl)
		payment.POST("/ali-notify-url", http.PaymentAliNotifyUrl)


		payment.GET("/wx-native-scanned-notify", http.PaymentWxNativeScannedNotify)
		payment.POST("/wx-native-scanned-notify", http.PaymentWxNativeScannedNotify)

		payment.GET("/wx-native-paid-notify", http.PaymentWxNativePaidNotify)
		payment.POST("/wx-native-paid-notify", http.PaymentWxNativePaidNotify)
	}
}
