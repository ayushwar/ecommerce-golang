package main

import (
	"github.com/ayushwar/ecommerce/database"
	"github.com/ayushwar/ecommerce/routes"
	"github.com/gin-gonic/gin"
	"log"
)

func main()  {
	database.ConnectDB()
	r:=gin.Default()
	routes.Router(r)
	log.Println("🚀 Server running at http://localhost:8080")
	r.Run(":8000")
}