USE alta_online_shop;

SELECT * FROM users WHERE created_at >= DATE_SUB(NOW(), INTERVAL 7 DAY) AND
`name` LIKE '%a%';
