package repository

const (
	queryGetAll = "SELECT id, locality_name, province_name, country_name FROM localities"
	queryGetById = "SELECT id, locality_name, province_name, country_name FROM localities WHERE id = ?"
	queryCreate = "INSERT INTO localities (locality_name, province_name, country_name) VALUES (?, ?, ?)"
	queryUpdate = "UPDATE localities SET locality_name = ?, province_name = ?, country_name = ? WHERE id = ?"
	queryDelete = "DELETE FROM localities WHERE id = ?"
)