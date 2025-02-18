# GoShop

![GoShop](https://socialify.git.ci/thinknyx-technologies-llp/GoShop/image?font=KoHo&logo=https%3A%2F%2Fcdn-icons-png.flaticon.com%2F512%2F2998%2F2998251.png&name=1&owner=1&pattern=Plus&theme=Light)

## Go Capstone Project

**GoShop: E-commerce Application**

This is a simple E-commerce application built with Golang that allows users to place orders, view products, and lets admins manage products and orders. This app demonstrates basic CRUD operations and authentication functionalities.

## Prerequisites

Before setting up this project, ensure you have the following:

1. **Go Programming Language**
    - To install Go, visit the official [Go Installation Guide](https://golang.org/doc/install). Download and follow the instructions specific to your operating system.

2. **IDE for Go Development**
    - A popular choice is Visual Studio Code (VS Code).
    - To download VS Code, visit the official [download page](https://code.visualstudio.com/Download).
    - After installing VS Code, search for the Go extension in the Extensions Marketplace to enable features like IntelliSense, debugging, testing, and more.

3. **MySQL Database**
    - Install MySQL 8 or above version.

## Project Setup

### 1. Clone the Repository
Clone this repository to your local machine:

```
git clone https://github.com/thinknyx-technologies-llp/GoShop.git
cd GoShop
```

### 2. Directory Structure
Ensure your project directory structure is as follows:

![GoShop-Directory-Structure](https://www.thinknyx.com/wp-content/uploads/2025/02/directory_structure_golang.png)

### 3. Database Setup

a) In the `GoShop/database/schema.sql`, the SQL schema is defined to create the necessary tables for orders and products.

b) Execute the commands written in `schema.sql` to set up your database schema.

c) Change the database credentials (username & password) in `db.go` according to your database configuration.

d) Ensure your MySQL database is set up and running before using the app.

### 4. Initialize the Go Module
In the terminal, navigate to your project folder and initialize the Go module:

```
go mod init GoShop    # Initialize the Go module
go mod tidy            # Install all dependencies
```

### 5. Run the Application
To run the project, execute:

```
go run main.go    # Start the Go server
```

Once the server is running, open your browser and visit the following URL: http://localhost:8082

## App Features

### 1. User Page
- Click on the **`User`** button to view available products.
- Place orders by adding products to the cart and confirming the order.

### 2. Admin Page
- Click on the **`Admin`** button to access the admin login page.
- Admin credentials:
  - **Username**: `admin`
  - **Password**: `password`
- Once logged in, the admin can:
  - **`Add`**, **`view`**, and **`delete`** products.

## Troubleshooting

- **Cannot connect to database:**
  - Ensure your database credentials in db.go are correct and that the database is running.
- **404 Error:**
  - Make sure you are navigating to the correct routes (`/user` and `/admin`).

## Contributing

Feel free to fork this repository and submit pull requests with enhancements and bug fixes.

## Powered By
*This project is Powered by [Thinknyx Technologies LLP](www.thinknyx.com)*
