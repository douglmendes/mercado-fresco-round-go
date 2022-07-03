package mariadb

const (
	GetAllQuery = `
		SELECT
			id,
			product_code,
			description,
			width,
			height,
			length,
			net_weight,
			expiration_rate,
			recommended_freezing_temperature,
			freezing_rate,
			product_type_id,
			seller_id
		FROM
			products`
	GetByIdQuery = `
		SELECT
			id,
			product_code,
			description,
			width,
			height,
			length,
			net_weight,
			expiration_rate,
			recommended_freezing_temperature,
			freezing_rate,
			product_type_id,
			seller_id
		FROM
			products
		WHERE
			id = ?`
	CreateQuery = `
		INSERT INTO products (
			product_code,
			description,
			width,
			height,
			length,
			net_weight,
			expiration_rate,
			recommended_freezing_temperature,
			freezing_rate,
			product_type_id,
			seller_id
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	UpdateQuery = `
		UPDATE products
		SET
			product_code = ?,
			description = ?,
			width = ?,
			height = ?,
			length = ?,
			net_weight = ?,
			expiration_rate = ?,
			recommended_freezing_temperature = ?,
			freezing_rate = ?,
			product_type_id = ?,
			seller_id = ?
		WHERE id = ?`
)
