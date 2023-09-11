USE alta_online_shop;

SET @product_id = (SELECT id FROM products WHERE product_type_id = 1);
# Delete from product desc
DELETE FROM product_descriptions WHERE product_id = @product_id;
# Delete from transaction detail
DELETE FROM transaction_details WHERE product_id = @product_id;
DELETE FROM products WHERE product_type_id = 1;