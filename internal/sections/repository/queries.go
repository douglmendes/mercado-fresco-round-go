package repository

const (
	GetAllQuery = `
		SELECT
			id,
			section_number,
			current_temperature,
			minimum_temperature,
			current_capacity,
			minimum_capacity,
			maximum_capacity,
			warehouse_id,
			product_type_id
		FROM
			sections`
	GetByIdQuery = `
		SELECT
			id,
			section_number,
			current_temperature,
			minimum_temperature,
			current_capacity,
			minimum_capacity,
			maximum_capacity,
			warehouse_id,
			product_type_id
		FROM
			sections
		WHERE
			id = ?`
	CreateQuery = `
		INSERT INTO sections (
			section_number,
			current_temperature,
			minimum_temperature,
			current_capacity,
			minimum_capacity,
			maximum_capacity,
			warehouse_id,
			product_type_id
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	UpdateQuery = `
		UPDATE sections
		SET
			section_number = ?,
			current_temperature = ?,
			minimum_temperature = ?,
			current_capacity = ?,
			minimum_capacity = ?,
			maximum_capacity = ?,
			warehouse_id = ?,
			product_type_id = ?
		WHERE id = ?`
	DeleteQuery = `
		DELETE
		FROM
			sections
		WHERE
			id = ?`
)
