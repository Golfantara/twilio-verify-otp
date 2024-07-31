package main

import (
	"github.com/Golfantara/twilio-verify-otp/api"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	app := api.Config{Router: router}

	app.Routes()

	router.Run(":8000")
}