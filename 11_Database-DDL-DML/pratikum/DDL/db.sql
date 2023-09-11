USE alta_online_shop;

CREATE TABLE product_types (
	id INT AUTO_INCREMENT PRIMARY KEY,
	name VARCHAR(255) NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP
);

CREATE TABLE operators (
	id INT AUTO_INCREMENT PRIMARY KEY,
	NAME VARCHAR(255) NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP
);

CREATE TABLE products (
	id INT AUTO_INCREMENT PRIMARY KEY,
	product_type_id INT NOT NULL,
	operator_id INT NOT NULL,
	CODE VARCHAR(50) NOT NULL,
	NAME VARCHAR(100) NOT NULL,
	STATUS SMALLINT DEFAULT 0,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP,
	
	FOREIGN KEY (product_type_id) REFERENCES product_types(id),
	FOREIGN KEY (operator_id) REFERENCES operators(id)
);

CREATE TABLE product_descriptions (
	id INT AUTO_INCREMENT PRIMARY KEY,
	description TEXT,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP
);

CREATE TABLE payment_methods (
	id INT AUTO_INCREMENT PRIMARY KEY,
	NAME VARCHAR(255) NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP
);

CREATE TABLE users (
	id INT AUTO_INCREMENT PRIMARY KEY,
	STATUS SMALLINT DEFAULT 0,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP
)

CREATE TABLE transactions (
	id INT AUTO_INCREMENT PRIMARY KEY,
	user_id INT NOT NULL,
	payment_method_id INT NOT NULL,
	STATUS VARCHAR(10),
	total_qty INT NOT NULL,
	total_price INT NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP,
	
	FOREIGN KEY (user_id) REFERENCES users(id),
	FOREIGN KEY (payment_method_id) REFERENCES payment_methods(id)
);

CREATE TABLE transaction_details (
	transaction_id INT NOT NULL,
	product_id INT NOT NULL,
	STATUS VARCHAR(10),
	qty INT NOT NULL,
	price NUMERIC(25,2) NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP,
	
	FOREIGN KEY (transaction_id) REFERENCES transactions(id),
	FOREIGN KEY (product_id) REFERENCES products(id)
);


