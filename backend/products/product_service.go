package products

import (
	"GoShop/backend/models"
	"GoShop/database"
	"database/sql"
	"fmt"
)

// GetProductsService retrieves all products from the database
func GetProductsService() ([]models.Product, error) {
	if database.DB == nil {
		return nil, fmt.Errorf("database is unavailable")
	}

	rows, err := database.DB.Query("SELECT id, name, quantity, price FROM products")
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve products")
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var p models.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Quantity, &p.Price); err != nil {
			return nil, fmt.Errorf("error scanning product data")
		}
		products = append(products, p)
	}

	return products, nil
}

// AddOrUpdateProductService adds or updates a product
func AddOrUpdateProductService(product models.Product) (models.Product, error) {
	var existingProduct models.Product
	err := database.DB.QueryRow("SELECT id, name, quantity, price FROM products WHERE id = ?", product.ID).Scan(
		&existingProduct.ID, &existingProduct.Name, &existingProduct.Quantity, &existingProduct.Price,
	)

	if err != nil && err != sql.ErrNoRows {
		return product, fmt.Errorf("database query failed")
	}

	if existingProduct.ID == 0 {
		// Insert new product
		stmt, err := database.DB.Prepare("INSERT INTO products (id, name, quantity, price) VALUES (?, ?, ?, ?)")
		if err != nil {
			return product, fmt.Errorf("failed to prepare insert statement")
		}
		defer stmt.Close()

		_, err = stmt.Exec(product.ID, product.Name, product.Quantity, product.Price)
		if err != nil {
			return product, fmt.Errorf("failed to insert product")
		}

		return product, nil
	}

	// Update quantity of existing product
	stmt, err := database.DB.Prepare("UPDATE products SET quantity = ? WHERE id = ?")
	if err != nil {
		return product, fmt.Errorf("failed to prepare update statement")
	}
	defer stmt.Close()

	newQuantity := existingProduct.Quantity + product.Quantity
	_, err = stmt.Exec(newQuantity, product.ID)
	if err != nil {
		return product, fmt.Errorf("failed to update product")
	}

	product.Quantity = newQuantity
	return product, nil
}

// DeleteProductService deletes a product by ID
func DeleteProductService(productID int) error {
	var quantity int
	err := database.DB.QueryRow("SELECT quantity FROM products WHERE id = ?", productID).Scan(&quantity)
	if err == sql.ErrNoRows {
		return fmt.Errorf("product not found")
	} else if err != nil {
		return fmt.Errorf("database query failed")
	}

	_, err = database.DB.Exec("DELETE FROM products WHERE id = ?", productID)
	if err != nil {
		return fmt.Errorf("failed to delete product")
	}

	return nil
}
