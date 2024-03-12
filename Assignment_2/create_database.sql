CREATE DATABASE orders_by;

USE orders_by;

CREATE TABLE items (
    item_id INT AUTO_INCREMENT PRIMARY KEY,
    item_code VARCHAR(255),
    description VARCHAR(255),
    quantity INT,
    order_id INT
);

CREATE TABLE orders (
    order_id INT AUTO_INCREMENT PRIMARY KEY,
    customer_name VARCHAR(255),
    ordered_at DATETIME
);
