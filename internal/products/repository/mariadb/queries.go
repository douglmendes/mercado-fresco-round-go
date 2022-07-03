package mariadb

const (
	GetAllQuery = `
		SELECT
			id,
			description,
			expiration_rate,
			freezing_rate,
			height,
			length,
			net_weight,
			product_code,
			recommended_freezing_temperature,
			width,
			product_type_id,
			seller_id
		FROM
			products`
)
