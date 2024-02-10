package controllers

import (
	"fmt"
	"net/http"

	"github.com/RaghavTheGreat1/go_otp/models"
	"github.com/gin-gonic/gin"
)

func SendSmsOtpController(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	defer c.Request.Body.Close()

	var phoneNumber models.PhoneNumber

	if err := c.BindJSON(&phoneNumber); err != nil {
		panic(err)
	}

	_, err := phoneNumber.SendOtp()

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": fmt.Sprintf("OTP sent to %s %d", phoneNumber.CountryCode, phoneNumber.PhoneNumber),
	})
}

func SmsOtpVerifyController(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	defer c.Request.Body.Close()

	var phoneNumber models.PhoneNumber

	if err := c.BindJSON(&phoneNumber); err != nil {
		panic(err)
	}

	status, err := phoneNumber.CheckOtp()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	if !status {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": fmt.Sprintf("OTP verification failed for %s %d", phoneNumber.CountryCode, phoneNumber.PhoneNumber),
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": fmt.Sprintf("OTP sent to %s %d has been successfully verified", phoneNumber.CountryCode, phoneNumber.PhoneNumber),
	})

}
