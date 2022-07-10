package mariadb

const (
	CreateQuery = `
		INSERT INTO product_records (
			last_update_date,
			purchase_price,
			sale_price,
			product_id
		) VALUES (?, ?, ?, ?)
	`
	GetAllGroupByProductIdQuery = `
		SELECT
			product_records.product_id,
			products.description,
			count(*) AS records_count
		FROM product_records
		INNER JOIN products ON product_records.product_id = products.id
		GROUP BY product_records.product_id;
	`
	GetAllGroupByProductIdWhereIdQuery = `
		SELECT
			product_records.product_id,
			products.description,
			count(*) AS records_count
		FROM product_records
		INNER JOIN products ON product_records.product_id = products.id
		WHERE product_records.product_id = ?
		GROUP BY product_records.product_id;
	`
)
