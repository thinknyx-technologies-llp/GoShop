package orders

import (
	"GoShop/backend/models"
	"GoShop/database"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetOrdersService(c *gin.Context) {
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

// Define bill details struct
type Bill struct {
	Details     []models.Order `json:"details"`
	TotalAmount int            `json:"total_amount"`
}

func PlaceOrdersService(c *gin.Context) {
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

	// Declare billDetails and totalAmount here
	var billDetails []models.Order
	var totalAmount int

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

		totalPrice := price * order.Quantity
		totalAmount += totalPrice

		_, err = tx.Exec("INSERT INTO orders (product_id, quantity, price) VALUES (?, ?, ?)", order.ProductID, order.Quantity, totalPrice)
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		res, err := tx.Exec("UPDATE products SET quantity = quantity - ? WHERE id = ? AND quantity >= ?", order.Quantity, order.ProductID, order.Quantity)
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		rowsAffected, err := res.RowsAffected()
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if rowsAffected == 0 {
			tx.Rollback()
			c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient stock. Order could not be placed."})
			return
		}

		// Append order details to billDetails
		billDetails = append(billDetails, models.Order{
			ProductID: order.ProductID,
			Quantity:  order.Quantity,
			Price:     totalPrice,
		})
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Order placed successfully",
		"bill": Bill{
			Details:     billDetails,
			TotalAmount: totalAmount,
		},
		"orders": billDetails,
	})
}

func CheckItemExistsService(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID"})
		return
	}

	var exists bool
	err = database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM products WHERE id=? AND quantity >= 1)", id).Scan(&exists)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"exists": exists})
}
