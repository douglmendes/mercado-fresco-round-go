package repository

const (
	queryCreate  = "insert into buyers (id_card_number, first_name, last_name) values (?,?,?)"
	queryGetAll  = "SELECT id, id_card_number, first_name, last_name FROM buyers"
	queryGetById = "SELECT id, id_card_number, first_name, last_name FROM buyers where id = ?"
	queryUpdate  = "update buyers set id_card_number = ?, first_name  = ?, last_name  = ? where id = ?"
	queryDelete  = "DELETE FROm buyers WHERE id = ?"
	queryGetOrdersByBuyer = "SELECT b.id, b.id_card_number, b.first_name, b.last_name, count(p.id) AS purchase_orders_count FROM buyers b INNER JOIN purchase_orders p ON b.id = p.buyer_id WHERE b.id = ? GROUP BY b.id"
	queryGetOrdersByBuyers = "SELECT b.id, b.id_card_number, b.first_name, b.last_name, count(p.id) AS purchase_orders_count FROM buyers b INNER JOIN purchase_orders p ON b.id = p.buyer_id GROUP BY b.id"
)
