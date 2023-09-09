USE alta_online_shop;

CREATE TABLE transaction_detail (
	id INT AUTO_INCREMENT PRIMARY KEY,
	transaction_id INT NOT NULL,
	product_id INT NOT NULL,
	qty INT NOT NULL,
	price INT NOT NULL,
	STATUS ENUM('cancel','pending','success') NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP,
	FOREIGN KEY (transaction_id) REFERENCES TRANSACTION(id),
	FOREIGN KEY (product_id) REFERENCES product(id)
);