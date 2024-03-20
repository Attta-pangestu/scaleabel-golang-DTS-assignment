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


INSERT INTO items (item_code, description, quantity, order_id) VALUES
('123', 'Beras', 2, 1),
('456', 'Minyak Goreng', 3, 1),
('789', 'Gula', 1, 2),
('101', 'Telur', 1, 2),
('202', 'Sabun Mandi', 2, 3),
('303', 'Pasta Gigi', 1, 3);


INSERT INTO orders (customer_name, ordered_at) VALUES
('Budi', '2024-03-11 10:00:00'),
('Siti', '2024-03-12 11:30:00'),
('Joko', '2024-03-13 09:45:00');


