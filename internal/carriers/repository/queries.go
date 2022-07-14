package repository

const (
	sqlCreateCarrier  = "INSERT INTO carries (cid, company_name, address, telephone, locality_id) VALUES (?, ?, ?, ?, ?)"
	sqlGetAllCarriers = "SELECT id, cid, company_name, address, telephone, locality_id FROM carries"
)
