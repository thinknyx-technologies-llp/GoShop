# GoShop

![GoShop](https://socialify.git.ci/thinknyx-technologies-llp/GoShop/image?font=KoHo&issues=1&logo=https%3A%2F%2Fcdn-icons-png.flaticon.com%2F128%2F18902%2F18902474.png&name=1&owner=1&pattern=Plus&pulls=1&stargazers=1&theme=Light)

## Go Capstone Project

**Go-Shop: E-commerce Application**
This is a simple E-commerce application built with Golang that allows users to place orders, view products, and lets admins manage products and orders. This app demonstrates basic CRUD operations and authentication functionalities.
Prerequisites
Before setting up this project, ensure you have the following:
1.	Go Programming Language
o	To install Go, visit the official Go Installation Guide. Download and follow the instructions specific to your operating system.
2.	IDE for Go Development
o	A popular choice is Visual Studio Code (VS Code).
o	To download VS Code, visit the official download page.
o	After installing VS Code, search for the Go extension in the Extensions Marketplace to enable features like IntelliSense, debugging, testing, and more.

Project Setup
**1. Clone the Repository**
Clone this repository to your local machine:
git clone 
cd go-shop
**2. Directory Structure**
Ensure your project directory structure is as follows:

https://www.thinknyx.com/wp-content/uploads/2025/02/directory_structure_golang.png 

**3. Initialize the Go Module**
In the terminal, navigate to your project folder and initialize the Go module:
go mod init go-shop              # Initialize the Go module
go mod tidy                      # Install all dependencies

**4. Run the Application**
To run the project, execute:
go run main.go                  # Start the Go server
Once the server is running, open your browser and visit the following URL:
http://localhost:8082

**App Features**
**User Page**
•	Click on the User button to view available products.
•	Place orders by adding products to the cart and confirming the order.
**Admin Page**
•	Click on the Admin button to access the admin login page.
•	Admin credentials:
o	Username: admin
o	Password: password
•	Once logged in, the admin can add, view, and delete products.

**Database Setup****
1.	In the go-shop/database/schema.sql, the SQL schema is defined to create the necessary tables for orders and products.
2.	Ensure your database is set up and running before using the app.

**Troubleshooting**
1.	Cannot connect to database: Ensure your database credentials in db.go are correct and that the database is running.
2.	404 Error: Make sure you are navigating to the correct routes (/user and /admin).

Contributing
If you'd like to contribute to this project, feel free to fork it and submit pull requests. We appreciate any improvements or bug fixes!



