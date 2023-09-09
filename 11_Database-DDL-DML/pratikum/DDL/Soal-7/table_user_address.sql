USE alta_online_shop;

CREATE TABLE user_address (
	id INT AUTO_INCREMENT PRIMARY KEY,
	user_id INT NOT NULL,
	alamat VARCHAR(255) NOT NULL,
	FOREIGN KEY (user_id) REFERENCES user(id)
);

ALTER TABLE user
DROP COLUMN alamat;