package repository

import (
	"database/sql"
	"fmt"
	"github.com/douglmendes/mercado-fresco-round-go/internal/buyers/domain"
	"log"
)

type repository struct {
	db *sql.DB
}

func (r *repository) GetAll() ([]domain.Buyer, error) {
	rows, err := r.db.Query(queryGetAll)
	if err != nil {
		log.Println("Error while quering buyers table" + err.Error())
		return nil, err
	}
	buyers := make([]domain.Buyer, 0)
	for rows.Next() {
		var b domain.Buyer
		err := rows.Scan(&b.Id, &b.CardNumberId, &b.FirstName, &b.LastName)
		if err != nil {
			log.Println("Error while scanning buyers" + err.Error())
			return nil, err
		}
		buyers = append(buyers, b)
	}
	return buyers, nil
}

func (r *repository) GetById(id int) (*domain.Buyer, error) {
	row := r.db.QueryRow(queryGetById, id)
	var b domain.Buyer
	err := row.Scan(&b.Id, &b.CardNumberId, &b.FirstName, &b.LastName)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Buyer %d not found", id)
		} else {
			log.Println("Error while scanning customer " + err.Error())
			return nil, fmt.Errorf("unexpected database error")
		}
	}
	return &b, nil
}

func (r *repository) Create(cardNumberId, firstName, lastName string) (*domain.Buyer, error) {
	result, err := r.db.Exec(queryCreate, cardNumberId, firstName, lastName)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("error retrieving id %d ", id)
	}
	e := domain.Buyer{int(id), cardNumberId, firstName, lastName}
	return &e, nil
}

func (r *repository) Update(id int, cardNumberId, firstName, lastName string) (*domain.Buyer, error) {

	buy, err := r.GetById(id)
	if err != nil {
		return nil, fmt.Errorf("employee %d not found", id)
	}
	if cardNumberId != "" {
		buy.CardNumberId = cardNumberId
	}
	if firstName != "" {
		buy.FirstName = firstName
	}
	if lastName != "" {
		buy.LastName = lastName
	}
	result, err := r.db.Exec(queryUpdate, buy.CardNumberId, buy.FirstName, buy.LastName, id)
	log.Println(result.RowsAffected())

	return buy, nil

}

func (r *repository) Delete(id int) error {
	_, err := r.db.Exec(queryDelete, id)
	if err != nil {
		return fmt.Errorf("error when deleting buyer %d", id)
	}
	return nil
}

func NewRepository(db *sql.DB) domain.Repository {
	return &repository{
		db: db,
	}

}