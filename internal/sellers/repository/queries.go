package repository

const (
	queryGetAll  = "SELECT id, cid, company_name, address, telephone, locality_id FROM sellers"
	queryGetById = "SELECT id, cid, company_name, address, telephone, locality_id FROM sellers where id = ?"
	queryCreate  = "INSERT INTO sellers (cid, company_name, address, telephone, locality_id) VALUES (?, ?, ?, ?, ?)"
	queryUpdate  = "UPDATE sellers SET cid = ?, company_name = ?, address = ?, telephone = ?, locality_id = ? WHERE id = ?"
	queryDelete  = "DELETE FROM sellers WHERE id = ?"
	// queryGetLocality = "SELECT id, locality_name, province_name, country_name FROM localities WHERE id = ?"
)