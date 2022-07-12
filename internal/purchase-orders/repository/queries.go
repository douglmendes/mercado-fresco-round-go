package repository

const (
	queryCreate = "insert into purchase_orders (order_number, order_date, tracking_code, buyer_id, order_status_id, warehouse_id) values(?,?,?,?,?,?,?)"
	queryGetAll = "SELECT id, order_number , order_date , tracking_code , buyer_id, order_status_id , warehouse_id  from purchase_orders "
)
