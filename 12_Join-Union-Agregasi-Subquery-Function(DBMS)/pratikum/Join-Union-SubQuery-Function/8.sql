USE alta_online_shop;

SELECT * FROM products WHERE id NOT IN (SELECT product_id FROM transaction_details)