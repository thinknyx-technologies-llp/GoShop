package products

import (
	"GoShop/backend/models"
	"GoShop/database"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes sets up product-related routes
func RegisterRoutes(router *gin.Engine) {
	productGroup := router.Group("/products")
	{
		productGroup.GET("/", GetProducts)
		productGroup.POST("/", AddOrUpdateProduct)
		productGroup.DELETE("/:id", DeleteProduct)
	}
}

// GetProducts fetches all products
func GetProducts(c *gin.Context) {
	if !database.DBConnected {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Database not available"})
		return
	}

	products, err := GetProductsService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}

// AddOrUpdateProduct adds a new product or updates an existing one
func AddOrUpdateProduct(c *gin.Context) {
	if !database.DBConnected {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Database not available"})
		return
	}

	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	updatedProduct, err := AddOrUpdateProductService(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product processed successfully", "product": updatedProduct})
}

// DeleteProduct removes a product by ID
func DeleteProduct(c *gin.Context) {
	if !database.DBConnected {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Database not available"})
		return
	}

	id := c.Param("id")
	productID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	err = DeleteProductService(productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
