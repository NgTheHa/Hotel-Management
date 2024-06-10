package main

import (
	"hotel-management/utils"
	"log"
)

func main() {
	utils.ConnectDatabase()
	// Your code to start the Gin server, define routes, etc.
	log.Println("Server is running...")

}
