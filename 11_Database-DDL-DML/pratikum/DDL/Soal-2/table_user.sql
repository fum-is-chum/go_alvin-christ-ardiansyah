USE alta_online_shop;

CREATE TABLE user (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nama VARCHAR(255) NOT NULL,
    alamat VARCHAR(255) NOT NULL,
    tanggal_lahir DATE NOT NULL,
    status VARCHAR(50) NOT NULL,
    gender ENUM('Male', 'Female') NOT NULL,
    PASSWORD VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);