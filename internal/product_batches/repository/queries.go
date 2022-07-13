package repository

const (
	createQuery              = "INSERT INTO product_batches (batch_number, current_quantity, current_temperature, due_date, initial_quantity, manufacturing_date, manufacturing_hour, minimum_temperature, product_id, section_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	getQuery                 = "SELECT id, batch_number, current_quantity, current_temperature, due_date, initial_quantity, manufacturing_date, manufacturing_hour, minimum_temperature, product_id, section_id FROM product_batches"
	singleSectionReportQuery = "SELECT product_batches.current_quantity, sections.section_id, sections.section_number FROM product_batches INNER JOIN sections ON product_batches.section_id = sections.id WHERE product_batches.section_id = ?"
	allSectionsReportQuery   = "SELECT product_batches.current_quantity, sections.section_id, sections.section_number FROM product_batches INNER JOIN sections ON product_batches.section_id = sections.id"
)
