USE alta_online_shop;

CREATE TABLE product (
	id INT AUTO_INCREMENT PRIMARY KEY,
	type_id INT NOT NULL,
	payment_method_id INT NOT NULL,
	nama VARCHAR(255) NOT NULL,
	deskripsi VARCHAR(255) NOT NULL,
	operator VARCHAR(255) NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP,
	FOREIGN KEY (type_id) REFERENCES product_type(id),
	FOREIGN KEY (payment_method_id) REFERENCES payment_method(id)
);