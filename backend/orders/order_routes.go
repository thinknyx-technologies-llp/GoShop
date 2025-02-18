package orders

import (
	"net/http"
	"strconv"

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

func CheckItemExists(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID"})
		return
	}

	exists := ItemExists(id)
	c.JSON(http.StatusOK, gin.H{"exists": exists})
}
