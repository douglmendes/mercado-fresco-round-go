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
			section`
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
			section
		WHERE
			id = ?`
	CreateQuery = `
		INSERT INTO section (
			section_number,
			current_temperature,
			minimum_temperature,
			current_capacity,
			minimum_capacity,
			maximum_capacity,
			warehouse_id,
			product_type_id
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	UpdateQuery = `
		UPDATE section
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
			section
		WHERE
			id = ?`
)
