DELIMITER $$
CREATE TRIGGER delete_transaction_detail
AFTER DELETE ON transactions FOR EACH ROW
BEGIN
    DELETE FROM transaction_details WHERE transaction_id = OLD.id;
END;
$$
DELIMITER ;