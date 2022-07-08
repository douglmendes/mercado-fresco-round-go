package repository

const (
	queryGetAll  = "SELECT id, cid, company_name, address, telephone FROM sellers"
	queryGetById = "SELECT id, cid, company_name, address, telephone FROM sellers where id = ?"
	queryCreate  = "INSERT INTO sellers (cid, company_name, address, telephone) VALUES (?, ?, ?, ?)"
	queryUpdate  = "UPDATE sellers SET cid = ?, company_name = ?, address = ?, telephone = ? WHERE id = ?"
	queryDelete  = "DELETE FROM sellers WHERE id = ?"
)
