package repository

const (
	queryGetAll       = "SELECT id,order_date,order_number,employee_id,product_batch_id, warehouse_id FROM inbound_orders"
	queryGetByEmplyee = "Select e.id ,e.id_card_number , e.first_name , e.last_name , e.warehouse_id ,count(*) as inbound_orders_count from inbound_orders io inner join employees e on e.id = io.employee_id where employee_id = ?"
	queryCreate       = "insert into inbound_orders(order_date, order_number, employee_id, product_batch_id, warehouse_id) values(?,?,?,?,?)"
)
