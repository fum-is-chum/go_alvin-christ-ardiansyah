USE alta_online_shop;

SELECT u.`id` AS 'User ID', u.`name`, SUM(t.total_price) AS 'Total Transaksi' FROM users u
INNER JOIN transactions t ON u.id = t.user_id AND t.user_id = 1;

