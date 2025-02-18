package orders

import (
	"encoding/json"
	"GoShop/backend/models"
	"GoShop/database"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetOrders(c *gin.Context) {
	rows, err := database.DB.Query("SELECT id, product_id, quantity, price FROM orders")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var orders []models.Order
	for rows.Next() {
		var o models.Order
		if err := rows.Scan(&o.ID, &o.ProductID, &o.Quantity, &o.Price); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		orders = append(orders, o)
	}

	c.JSON(http.StatusOK, orders)
}

func PlaceOrders(c *gin.Context) {
	var orders []models.Order
	if err := json.NewDecoder(c.Request.Body).Decode(&orders); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tx, err := database.DB.Begin()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var totalAmount int
	var billDetails []gin.H
	var returningOrder []gin.H

	for _, order := range orders {
		var price, availableQuantity int
		err := tx.QueryRow("SELECT price, quantity FROM products WHERE id=?", order.ProductID).Scan(&price, &availableQuantity)
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if availableQuantity < order.Quantity {
			tx.Rollback()
			c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient stock for product ID " + strconv.Itoa(order.ProductID)})
			return
		}

		price = price * order.Quantity
		totalAmount += order.Quantity * price

		result, err := tx.Exec("INSERT INTO orders (product_id, quantity, price) VALUES (?, ?, ?)", order.ProductID, order.Quantity, price)
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		orderID, err := result.LastInsertId()
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		_, err = tx.Exec("UPDATE products SET quantity = quantity - ? WHERE id = ?", order.Quantity, order.ProductID)
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		returningOrder = append(returningOrder, gin.H{
			"product_id":        order.ProductID,
			"quantity":          order.Quantity,
			"name":              order.Name,
			"id":                orderID,
			"price":             price,
			"availableQuantity": availableQuantity,
		})

		billDetails = append(billDetails, gin.H{
			"product_id":        order.ProductID,
			"quantity":          order.Quantity,
			"price":             price,
			"fetched_price":     price,
			"availableQuantity": availableQuantity,
		})
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"orders": returningOrder,
		"bill": gin.H{
			"total_amount": totalAmount,
			"details":      billDetails,
		},
	})
}

func ItemExists(id int) bool {
	var exists bool
	err := database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM products WHERE id=? AND quantity >= 1)", id).Scan(&exists)
	if err != nil {
		return false
	}
	return exists
}
