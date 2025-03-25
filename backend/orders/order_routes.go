package orders

import (
	"GoShop/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	orderGroup := router.Group("/orders")
	{
		orderGroup.GET("/", GetOrders)
		orderGroup.POST("/", PlaceOrders)
		orderGroup.GET("/exists/:id", CheckItemExists)
	}
}

func GetOrders(c *gin.Context) {
	if !database.DBConnected {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Database not available"})
		return
	}
	// Call the service function
	GetOrdersService(c)
}

func PlaceOrders(c *gin.Context) {
	if !database.DBConnected {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Database not available"})
		return
	}
	// Call the service function
	PlaceOrdersService(c)
}

func CheckItemExists(c *gin.Context) {
	if !database.DBConnected {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Database not available"})
		return
	}
	// Call the service function
	CheckItemExistsService(c)
}
