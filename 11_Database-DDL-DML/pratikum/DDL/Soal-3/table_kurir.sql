USE alta_online_shop;

CREATE TABLE kurir (
	id INT AUTO_INCREMENT PRIMARY KEY,
	nama VARCHAR(255) NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP
)