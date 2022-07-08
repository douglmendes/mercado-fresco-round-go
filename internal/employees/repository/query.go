package repository

const (
	queryGetAll  = "SELECT id, id_card_number, first_name, last_name, warehouse_id FROM employees"
	queryGetById = "SELECT id, id_card_number, first_name, last_name, warehouse_id FROM employees where id = ?"
	queryCreate  = "insert into employees (id_card_number , first_name, last_name, warehouse_id) values (?,?,?,?)"
	queryUpdate  = "UPDATE employees  SET id_card_number  =  ? , first_name= ? , last_name = ? , warehouse_id  = ? WHERE id=?"
	queryDelete  = "DELETE FROM employees  WHERE id =?"
)
