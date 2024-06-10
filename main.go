package main

import (
	"github.com/gin-gonic/gin"
	"hotel-management/routes"
	"hotel-management/utils"
	"log"
)

func main() {

	// connect DB
	utils.ConnectDatabase()

	router := gin.Default()

	routes.SetupRoutes(router)

	log.Println("Server is running...")

	//server port
	router.Run(":8080")
}
