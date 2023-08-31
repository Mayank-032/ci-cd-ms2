package router

import (
	"microservice2/handler/payment"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	paymentGroup := r.Group("payment")

	paymentGroup.POST("/verify", payment.VerifyPayment)
}
