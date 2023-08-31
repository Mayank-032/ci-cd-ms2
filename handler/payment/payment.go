package payment

import (
	"log"
	"microservice2/handler/payment/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

func VerifyPayment(c *gin.Context) {
	request := VerifyPaymentReq{}
	response := gin.H{"Status": "false"}

	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println("Error: " + err.Error())
		response["message"] = "invalid request"
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err := request.Validate()
	if err != nil {
		log.Println("Error: " + err.Error())
		response["message"] = "invalid request"
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err = usecase.ValidatePayment(c, request.Email)
	if err != nil {
		log.Println("Error: " + err.Error())
		if err.Error() == "no_payment_initiated" {
			response["message"] = "invalid payment"
			response["is_payment_verified"] = false
			c.JSON(http.StatusOK, response)
			return
		}
		response["message"] = "unable to valiadate user due to internal server error"
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	log.Println("user validated successfully")
	response["status"] = true
	response["message"] = "successfully validated user"
	response["is_payment_verified"] = true
	c.JSON(http.StatusOK, response)
	return
}
