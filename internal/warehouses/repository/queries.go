package repository

const (
	sqlCreate  = "INSERT INTO warehouse (address, telephone, warehouse_code, locality_id) VALUES (?, ?, ?, ?)"
	sqlGetAll  = "SELECT id, address, telephone, warehouse_code, locality_id FROM warehouse"
	sqlGetById = "SELECT id, address, telephone, warehouse_code, locality_id FROM warehouse WHERE id = ?"
	sqlDelete  = "DELETE FROM warehouse WHERE id = ?"
	sqlUpdate  = "UPDATE warehouse SET address = ?, telephone = ?, warehouse_code = ?, locality_id = ? WHERE id = ?"
)
