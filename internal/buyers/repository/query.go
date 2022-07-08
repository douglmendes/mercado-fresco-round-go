package repository

const (
	queryCreate  = "insert into buyers (id_card_number, first_name, last_name) values (?,?,?)"
	queryGetAll  = "SELECT id, id_card_number, first_name, last_name FROM buyers"
	queryGetById = "SELECT id, id_card_number, first_name, last_name FROM buyers where id = ?"
	queryUpdate  = "update buyers set id_card_number = ?, first_name  = ?, last_name  = ? where id = ?"
	queryDelete  = "DELETE FROm buyers WHERE id = ?"
)
