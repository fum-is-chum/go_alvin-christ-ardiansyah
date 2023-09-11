USE alta_online_shop;

# delete from product desc
DELETE FROM product_descriptions WHERE product_id = 1;
# delete from transaction details
DELETE FROM transaction_details WHERE product_id = 1;
DELETE FROM products WHERE id = 1;