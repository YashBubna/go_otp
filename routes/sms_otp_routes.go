package routes

import (
	"github.com/RaghavTheGreat1/go_otp/controllers"
	"github.com/gin-gonic/gin"
)

func SetUpSmsOtpRoute(router *gin.Engine) {

	router.POST("/phone/send-otp", controllers.SendSmsOtpController)
	router.POST("/phone/verify", controllers.SmsOtpVerifyController)
}
