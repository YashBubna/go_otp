package main

import (
	"fmt"

	"github.com/RaghavTheGreat1/go_otp/constants"
	"github.com/RaghavTheGreat1/go_otp/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	port := ":8080"
	godotenv.Load(".env")
	constants.InitializeTwilioClient()

	router := gin.New()
	router.Use(gin.Logger())

	routes.SetUpSmsOtpRoute(router)

	/// Starts server at specified port
	fmt.Println("Server running at port", port, "...")
	fmt.Println("http://localhost", port)
	router.Run(port)
}
