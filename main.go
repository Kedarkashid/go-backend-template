package main

import (
	"fmt"

	"github.com/Kedarkashid/go-backend-template/database"
	"github.com/Kedarkashid/go-backend-template/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	database.Connect()
	routes.SetupRoutes(r)

	fmt.Println("Backend is running...")
	r.Run(":8080")

}
