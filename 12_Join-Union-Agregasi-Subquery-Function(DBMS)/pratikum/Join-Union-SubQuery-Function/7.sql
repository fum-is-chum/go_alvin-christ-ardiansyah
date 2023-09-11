DELIMITER $$
CREATE TRIGGER update_transaction_qty
AFTER DELETE ON transaction_details FOR EACH ROW
BEGIN
    DECLARE deleted_qty INT;
    SET deleted_qty = OLD.qty;
    UPDATE `transactions`
    SET total_qty = total_qty - deleted_qty
    WHERE id = OLD.transaction_id;
END;
$$
DELIMITER ;