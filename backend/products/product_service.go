package products

import (
	"database/sql"
	"encoding/json"
	"go-shop/backend/models"
	"go-shop/database"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetProducts retrieves all products using raw SQL query
func GetProducts(c *gin.Context) {
	// Execute the raw SQL query using database.DB (which is *sql.DB)
	rows, err := database.DB.Query("SELECT id, name, quantity, price FROM products")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	// Slice to hold all the products
	var products []models.Product

	// Loop through the rows
	for rows.Next() {
		var p models.Product
		// Scan each row's data into a Product model
		if err := rows.Scan(&p.ID, &p.Name, &p.Quantity, &p.Price); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		// Append the product to the slice
		products = append(products, p)
	}

	// Return the products in JSON format
	c.JSON(http.StatusOK, products)
}

// AddOrUpdateProduct inserts a new product into the database or updates the quantity if the product already exists
func AddOrUpdateProduct(c *gin.Context) {
	var product models.Product
	// Decode the request body into the Product model
	if err := json.NewDecoder(c.Request.Body).Decode(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("Request Body: %+v", product)
	log.Println("Received product:", product)

	// Check if the product already exists
	var existingProduct models.Product
	err := database.DB.QueryRow("SELECT id, name, quantity, price FROM products WHERE id = ?", product.ID).Scan(&existingProduct.ID, &existingProduct.Name, &existingProduct.Quantity, &existingProduct.Price)
	if err != nil && err != sql.ErrNoRows {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if existingProduct.ID == 0 {
		// Product does not exist, insert a new product
		stmt, err := database.DB.Prepare("INSERT INTO products (id, name, quantity, price) VALUES (?, ?, ?, ?)")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer stmt.Close()

		_, err = stmt.Exec(product.ID, product.Name, product.Quantity, product.Price)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, product)
	} else {
		// Product exists, update the quantity by adding the new quantity to the existing quantity
		stmt, err := database.DB.Prepare("UPDATE products SET quantity = ? WHERE id = ?")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer stmt.Close()

		newQuantity := existingProduct.Quantity + product.Quantity
		_, err = stmt.Exec(newQuantity, product.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Update the product's quantity to reflect the new total quantity
		product.Quantity = newQuantity

		c.JSON(http.StatusOK, product)
	}
}

// DeleteProduct deletes a product from the database using raw SQL
func DeleteProduct(c *gin.Context) {
	id := c.Param("id") // Get the product ID from the URL parameter
	productID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	// Check if the product exists
	var quantity int
	err = database.DB.QueryRow("SELECT quantity FROM products WHERE id = ?", productID).Scan(&quantity)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Prepare the SQL statement for deleting the product by its ID
	// Execute the SQL statement directly
	_, err = database.DB.Exec("DELETE FROM products WHERE id = ?", productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with a status of "No Content" to indicate successful deletion
	c.Status(http.StatusNoContent)
}
