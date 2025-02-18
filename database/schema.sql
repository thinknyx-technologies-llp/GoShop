CREATE DATABASE shop;
USE shop;

CREATE TABLE products (id INT AUTO_INCREMENT PRIMARY KEY,name VARCHAR(100),quantity INT, price INT);

CREATE TABLE orders (id INT AUTO_INCREMENT PRIMARY KEY,product_id INT,quantity INT,price INT);

