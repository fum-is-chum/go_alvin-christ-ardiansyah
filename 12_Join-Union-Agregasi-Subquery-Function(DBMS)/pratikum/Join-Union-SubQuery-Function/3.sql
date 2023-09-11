USE alta_online_shop;

SELECT p.name AS 'Product Name', pt.name AS 'Product Type', SUM(t.price) AS 'Total Transaksi'
FROM transaction_details t
INNER JOIN products p ON p.id = t.product_id
INNER JOIN product_types pt ON pt.id = p.product_type_id
WHERE pt.id = 2
GROUP BY p.name;