USE alta_online_shop;

INSERT INTO transactions(`user_id`, `payment_method_id`, `status`, `total_qty`, `total_price`)
VALUES (1, 1, 'Completed', 13, 88500),
		 (2, 2, 'Completed', 10, 100000),
		 (3, 3, 'Completed', 8, 150000)

INSERT INTO transaction_details(transaction_id, product_id, `status`, `qty`, `price`)
VALUES
	(1,1,'Completed', 2, 25000),
	(1,2,'Completed', 1, 13500),
	(1,3,'Completed', 10, 50000),
	(2,3,'Completed', 3, 30000),
	(2,4,'Completed', 3, 30000),
	(3,5,'Completed', 4, 40000),
	(3,1,'Completed', 2, 18750),
	(3,2,'Completed', 2, 18750),
	(3,3,'Completed', 2, 18750)

