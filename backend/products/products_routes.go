package products

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {
	productGroup := router.Group("/products")
	{
		productGroup.GET("/", GetProducts)
		productGroup.POST("/", AddOrUpdateProduct)
		productGroup.DELETE("/:id", DeleteProduct)
	}
}
