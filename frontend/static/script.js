document.addEventListener("DOMContentLoaded", function () {
    loadProducts();
});

async function loadProducts() {
    const response = await fetch("http://localhost:8080/products/");
    const products = await response.json();
    const productList = document.getElementById("productList");
    productList.innerHTML = "";

    products.forEach(product => {
        const item = document.createElement("li");
        item.textContent = `${product.name} - $${product.price}`;
        productList.appendChild(item);
    });
}

document.getElementById("addProductForm")?.addEventListener("submit", async function (e) {
    e.preventDefault();
    const name = document.getElementById("name").value;
    const price = document.getElementById("price").value;

    await fetch("http://localhost:8080/products/", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ name, price })
    });

    alert("Product Added!");
    loadProducts();
});

document.getElementById("orderForm")?.addEventListener("submit", async function (e) {
    e.preventDefault();
    const productId = document.getElementById("productId").value;
    const quantity = document.getElementById("quantity").value;

    await fetch("http://localhost:8082/orders/", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ product_id: productId, quantity })
    });

    alert("Order Placed!");
});
