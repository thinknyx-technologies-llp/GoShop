package main

import (
	"fmt"
	"GoShop/backend/orders"
	"GoShop/backend/products"
	"GoShop/database"

	"net/http"

	"sync"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Admin Service (Port 8080)
func startAdminService(wg *sync.WaitGroup) {
	defer wg.Done()
	router := gin.Default()

	// Custom CORS config (allows everything)
	config := cors.Config{
		AllowOrigins:     []string{"*"},                                       // Allow all origins
		AllowMethods:     []string{"GET", "POST", "DELETE", "PUT"},            // Allow specific methods
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // Allow headers
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}

	router.Use(cors.New(config))

	// Serve static files (CSS, JS)
	router.Static("/static", "./frontend/static")

	// Serve templates
	router.LoadHTMLGlob("frontend/templates/*")

	// Admin Page
	router.GET("/admin", func(c *gin.Context) {
		c.HTML(200, "admin.html", nil)
	})

	// Handle form submission
	router.POST("/product/add", products.AddOrUpdateProduct)
	router.GET("/products", products.GetProducts)
	router.DELETE("/product/delete/:id", products.DeleteProduct)

	// Set mode to debug
	gin.SetMode(gin.DebugMode)

	fmt.Println("Product Service running on port 8080")
	router.Run(":8080")
}

// User Service (Port 8081)
func startUserService(wg *sync.WaitGroup) {
	defer wg.Done()
	router := gin.Default()

	// Custom CORS config (allows everything)
	config := cors.Config{
		AllowOrigins:     []string{"*"},                                       // Allow all origins
		AllowMethods:     []string{"GET", "POST", "DELETE", "PUT"},            // Allow specific methods
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // Allow headers
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}

	router.Use(cors.New(config))

	// Serve static files (CSS, JS)
	router.Static("/static", "./frontend/static")

	// Serve templates
	router.LoadHTMLGlob("frontend/templates/*")

	// User home route
	router.GET("/user", func(c *gin.Context) {
		c.HTML(200, "user.html", nil)
	})

	router.POST("/product/add", products.AddOrUpdateProduct)
	router.GET("/products", products.GetProducts)
	router.DELETE("/product/delete/:id", products.DeleteProduct)

	// Order routes
	router.GET("/orders", orders.GetOrders)
	router.POST("/orders", orders.PlaceOrders)
	router.GET("/orders/exists/:id", orders.CheckItemExists)

	fmt.Println("👤 User Service running on port 8081")
	router.Run(":8081")
}

// Main Service (Port 8082)
func startMainService(wg *sync.WaitGroup) {
	defer wg.Done()
	router := gin.Default()

	// Serve templates
	router.LoadHTMLGlob("frontend/templates/*")

	// Main Page
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	// Admin login page
	router.GET("/admin/login", func(c *gin.Context) {
		c.HTML(200, "admin_login.html", nil)
	})

	// Admin login route
	router.POST("/admin/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")

		// Simple authentication check
		if username == "admin" && password == "password" {
			c.Redirect(http.StatusFound, "http://localhost:8080/admin")
		} else {
			c.Redirect(http.StatusFound, "/?unauthorized=true")
		}
	})

	// User redirect route
	router.GET("/user", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "http://localhost:8081/user")
	})

	fmt.Println("Main Service running on port 8082")
	router.Run(":8082")
}

func main() {

	database.Connect()

	var wg sync.WaitGroup
	wg.Add(3)

	// Run all services concurrently
	go startUserService(&wg)
	go startAdminService(&wg)
	go startMainService(&wg)

	// Wait for all services to finish
	wg.Wait()
}
