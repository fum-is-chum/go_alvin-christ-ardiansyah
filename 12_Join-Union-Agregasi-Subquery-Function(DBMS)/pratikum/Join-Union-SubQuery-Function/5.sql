USE alta_online_shop;

SELECT t.*, p.name AS 'Product Name', u.name AS 'User Name' FROM transaction_details td
INNER JOIN products p ON p.id = td.product_id
INNER JOIN transactions t ON t.id = td.transaction_id
INNER JOIN users u ON u.id = t.user_id;