document.addEventListener("DOMContentLoaded", function () {
    console.log("scripts.js is running...");

    checkDatabaseStatus();
    loadProducts();

    // Handle product form submission with Bootstrap validation
    const addProductForm = document.getElementById("addProductForm");
    if (addProductForm) {
        addProductForm.addEventListener("submit", async function (e) {
            if (!addProductForm.checkValidity()) {
                e.preventDefault();
                e.stopPropagation();
                addProductForm.classList.add("was-validated");
                return;
            }
            e.preventDefault();

            const productID = document.getElementById("productID").value;
            const name = document.getElementById("name").value;
            const quantity = document.getElementById("quantity").value;
            const price = document.getElementById("price").value;

            try {
                const response = await fetch("http://localhost:8080/product/add", {
                    method: "POST",
                    headers: { "Content-Type": "application/json" },
                    body: JSON.stringify({
                        id: parseInt(productID),
                        name,
                        quantity: parseInt(quantity),
                        price: parseInt(price),
                    }),
                });

                const data = await response.json();

                if (!response.ok) {
                    throw new Error(data.message || "⚠ Failed to add product. Database may be unavailable.");
                }

                alert("✅ Product Added Successfully!");
                addProductForm.reset();
                addProductForm.classList.remove("was-validated");
                loadProducts();
            } catch (error) {
                console.error("Error adding product:", error);
                alert(error.message || "⚠ Failed to add product. Database may be unavailable.");
            }
        });
    }

    // ✅ Check if the database is available before loading products
    async function checkDatabaseStatus() {
        try {
            const response = await fetch("http://localhost:8080/products");
            if (!response.ok) {
                throw new Error("Database unavailable");
            }
        } catch (error) {
            console.warn("Database is not available:", error);
            document.getElementById("products").innerHTML = `
                <div class="alert alert-warning text-center">
                    ⚠ Database is not available. Products cannot be loaded.
                </div>
            `;
        }
    }

    // ✅ Load products into the table
    async function loadProducts() {
        try {
            const response = await fetch("http://localhost:8080/products");
            if (!response.ok) {
                throw new Error("⚠ Database unavailable.");
            }
            const products = await response.json();

            const tbody = document.querySelector("#productTable tbody");
            tbody.innerHTML = "";

            products.forEach((product) => {
                const row = document.createElement("tr");
                row.innerHTML = `
                    <td>${product.id}</td>
                    <td>${product.name}</td>
                    <td>${product.quantity === 0 ? '<span style="color: red; font-weight: bold;">SOLD OUT!</span>' : product.quantity}</td>
                    <td>${product.price}</td>
                    <td><button class="btn btn-danger btn-sm deleteBtn" data-id="${product.id}">Delete</button></td>
                `;
                tbody.appendChild(row);
            });
        } catch (error) {
            console.error("Error loading products:", error);
            document.getElementById("products").innerHTML = `
                <div class="alert alert-warning text-center">
                    ⚠ Database is not available. Products cannot be loaded.
                </div>
            `;
        }
    }

    // ✅ Handle product deletion
    document.querySelector("#productTable tbody").addEventListener("click", async function (event) {
        if (event.target.classList.contains("deleteBtn")) {
            const productId = event.target.getAttribute("data-id");

            try {
                const response = await fetch(`http://localhost:8080/product/delete/${productId}`, {
                    method: "DELETE",
                });

                if (!response.ok) {
                    throw new Error("⚠ Failed to delete product. Database may be unavailable.");
                }

                loadProducts();
            } catch (error) {
                console.error("Error deleting product:", error);
                alert(error.message || "⚠ Failed to delete product. Database may be unavailable.");
            }
        }
    });

    // ✅ Handle order placement
    document.getElementById("placeOrder").addEventListener("click", async function () {
        const productId = parseInt(document.getElementById("productID").value);
        const quantity = parseInt(document.getElementById("orderQuantity").value);

        try {
            // ✅ Fetch product details first
            const productResponse = await fetch(`http://localhost:8080/product/${productId}`);
            if (!productResponse.ok) {
                throw new Error("⚠ Failed to fetch product details.");
            }

            const product = await productResponse.json();
            
            // ✅ Check if requested quantity is available
            if (quantity > product.quantity) {
                alert("⚠ Not enough stock available!");
                return;
            }

            // ✅ Proceed with order placement
            const orderResponse = await fetch("http://localhost:8080/order/place", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify([{ productId, quantity }]),
            });

            const data = await orderResponse.json();

            if (!orderResponse.ok) {
                throw new Error(data.error || "⚠ Failed to place order.");
            }

            alert("✅ Order placed successfully!");
            loadProducts(); // ✅ Refresh stock after order placement
        } catch (error) {
            console.error("Error placing order:", error);
            alert(error.message || "⚠ Failed to place order.");
        }
    });
});

// ✅ Additional Click Event Handler for Order Placement

document.addEventListener("click", function(event) {
    if (event.target.classList.contains("place-order-btn")) {
        let productId = event.target.getAttribute("data-id");

        fetch(`/api/products/${productId}`)
            .then(response => response.json())
            .then(product => {
                if (product.quantity < 1) {
                    alert("This product is out of stock!");
                    return;
                }

                fetch("/api/orders", {
                    method: "POST",
                    headers: { "Content-Type": "application/json" },
                    body: JSON.stringify([{ product_id: productId, quantity: 1 }])
                    
                })
                .then(response => response.json())
                .then(result => {
                    if (result.error) {
                        alert("Failed to place order! " + result.error);
                    } else {
                        alert("Order placed successfully!");
                        loadProducts(); // Refresh UI
                    }
                })
                .catch(error => console.error("Error placing order:", error));
            })
            .catch(error => console.error("Error checking product stock:", error));
    }
});
