USE alta_online_shop;

SELECT p.*, pt.name AS 'Product Type' FROM products p
INNER JOIN product_types pt ON p.product_type_id = pt.id;