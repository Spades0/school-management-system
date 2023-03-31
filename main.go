package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ucheonukaji/sms/routes"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.AllRoutes(router)

	router.Run(":" + port)

}
