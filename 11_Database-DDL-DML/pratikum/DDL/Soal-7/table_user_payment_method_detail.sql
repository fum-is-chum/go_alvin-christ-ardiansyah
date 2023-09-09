USE alta_online_shop;

CREATE TABLE user_payment_method_detail (
	id INT AUTO_INCREMENT PRIMARY KEY,
	user_id INT NOT NULL,
	payment_method_id INT NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY (user_id) REFERENCES user(id),
	FOREIGN KEY (payment_method_id) REFERENCES payment_method(id)
)