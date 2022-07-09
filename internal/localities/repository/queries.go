package repository

const (
	queryGetAll        = "SELECT id, locality_name, province_name, country_name FROM localities"
	queryGetById       = "SELECT id, locality_name, province_name, country_name FROM localities WHERE id = ?"
	queryCreate        = "INSERT INTO localities (locality_name, province_name, country_name) VALUES (?, ?, ?)"
	queryUpdate        = "UPDATE localities SET locality_name = ?, province_name = ?, country_name = ? WHERE id = ?"
	queryDelete        = "DELETE FROM localities WHERE id = ?"
	queryGetBySeller   = "SELECT l.id, l.locality_name, count(s.id) AS sellers_count FROM localities l INNER JOIN sellers s ON l.id = s.locality_id WHERE l.id = ? GROUP BY l.id"
	queryGetBySellers  = "SELECT l.id, l.locality_name, count(s.id) AS sellers_count FROM localities l INNER JOIN sellers s ON l.id = s.locality_id GROUP BY l.id"
	queryGetByCarrier  = "SELECT l.id, l.locality_name, count(c.id) AS carriers_count FROM localities l INNER JOIN carries c ON l.id = c.locality_id WHERE l.id = ? GROUP BY l.id"
	queryGetByCarriers = "SELECT l.id, l.locality_name, count(s.id) AS carriers_count FROM localities l INNER JOIN carries s ON l.id = s.locality_id GROUP BY l.id"
)
