USE alta_online_shop;

INSERT INTO product_descriptions (product_id, `description`)
SELECT id, CONCAT('Description for product with id ', id) FROM products;
