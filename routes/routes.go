package routes

import (
	"github.com/ayushwar/ecommerce/controllers"
	"github.com/ayushwar/ecommerce/middlewares"
	"github.com/gin-gonic/gin"
)
func Router(r *gin.Engine)  {
	r.POST("/register",controllers.Register)
	r.POST("/login",controllers.Login)

	api:=r.Group("/api",middlewares.AuthMiddleware())
	api.POST("/products", controllers.CreateProduct)   // Admin should use this
	// api.GET("/products", controllers.GetProducts)
	// api.PUT("/products/:id", controllers.UpdateProduct)
	api.DELETE("/products/:id", controllers.DeleteProduct)

	// 🛍️ Cart routes
	api.POST("/cart", controllers.AddToCart)
	api.GET("/cart/:user_id", controllers.GetCart)
	api.DELETE("/cart/:id", controllers.RemoveFromCart)

	// 📦 Order routes
	api.POST("/orders", controllers.PlaceOrder)
	api.GET("/orders/:user_id", controllers.GetOrders)
	
}