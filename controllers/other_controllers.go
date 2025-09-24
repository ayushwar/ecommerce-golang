package controllers

import (
	"github.com/ayushwar/ecommerce/database"
	"github.com/ayushwar/ecommerce/models"
	"github.com/gin-gonic/gin"
	"time"
)

/////////////////////////////////////////////////////
// PRODUCT CONTROLLERS(Admin only)
/////////////////////////////////////////////////////

func CreateProduct(ctx *gin.Context)  {
	var product models.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(400, gin.H{"error": "invalid request"})
		return
	}

	if err := database.DB.Create(&product).Error; err != nil {
		ctx.JSON(400, gin.H{"error": "could not create product"})
		return
	}

	ctx.JSON(200,product)
	

}
func Getallproduc(ctx *gin.Context)  {
	var product []models.Product
	database.DB.Find(&product)
	ctx.JSON(200,product)

	
}
func Updateproduct(ctx *gin.Context)  {
	id:=ctx.Param("id")
	var product models.Product
	if err := database.DB.First(&product, id).Error; err != nil {
		ctx.JSON(400, gin.H{"error": "product not found"})
		return
	}

	var input models.Product
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(400, gin.H{"error": "invalid request"})
		return
	}

	product.Description = input.Description
	product.Price = input.Price
	product.Stock = input.Stock
	product.UpdatedAt = time.Now()

	database.DB.Save(&product)
	ctx.JSON(200, product)
	
}
func DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := database.DB.Delete(&models.Product{}, id).Error; err != nil {
		ctx.JSON(500, gin.H{"error": "could not delete product"})
		return
	}
	ctx.JSON(200, gin.H{"message": "product deleted"})
}

/////////////////////////////////////////////////////
// CART CONTROLLERS
/////////////////////////////////////////////////////
func AddToCart(ctx *gin.Context) {
	var cartItem models.CartItem
	if err := ctx.ShouldBindJSON(&cartItem); err != nil {
		ctx.JSON(400, gin.H{"error": "invalid request"})
		return
	}
	if err := database.DB.Create(&cartItem).Error; err != nil {
		ctx.JSON(500, gin.H{"error": "could not add to cart"})
		return
	}
	ctx.JSON(200, cartItem)
}
func GetCart(ctx *gin.Context) {
	userID := ctx.Param("user_id")
	var cart []models.CartItem
	database.DB.Preload("Product").Where("user_id = ?", userID).Find(&cart)
	ctx.JSON(200, cart)
}
func RemoveFromCart(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := database.DB.Delete(&models.CartItem{}, id).Error; err != nil {
		ctx.JSON(500, gin.H{"error": "could not remove cart item"})
		return
	}
	ctx.JSON(200, gin.H{"message": "item removed"})
}
/////////////////////////////////////////////////////
// ORDER CONTROLLERS
/////////////////////////////////////////////////////

// Place order from cart
func PlaceOrder(ctx *gin.Context) {
	var input struct {
		UserID uint `json:"user_id"`
	}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(400, gin.H{"error": "invalid request"})
		return
	}
		// Get cart items
		var cartItems []models.CartItem
	database.DB.Preload("Product").Where("user_id = ?", input.UserID).Find(&cartItems)
	if len(cartItems) == 0 {
		ctx.JSON(400, gin.H{"error": "cart is empty"})
		return
	}

	// Create order
	order := models.Order{
		UserID: input.UserID,
		Status: "Pending",
	}
	database.DB.Create(&order)
	
	var total int64
	var orderItems []models.OrderItem
	for _, item := range cartItems {
		price := item.Product.Price * int64(item.Quantity)
		orderItems = append(orderItems, models.OrderItem{
			OrderID:   order.ID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     item.Product.Price,
		})
		total += price
	}


	// Save order items
	for _, oi := range orderItems {
		database.DB.Create(&oi)
	}

	// Update order total
	order.Total = total
	database.DB.Save(&order)

	// Clear cart
	database.DB.Where("user_id = ?", input.UserID).Delete(&models.CartItem{})

	ctx.JSON(200, gin.H{"message": "order placed", "order_id": order.ID})

}

// Get user orders
func GetOrders(ctx *gin.Context) {
	userID := ctx.Param("user_id")
	var orders []models.Order
	database.DB.Preload("Items.Product").Where("user_id = ?", userID).Find(&orders)
	ctx.JSON(200, orders)
}