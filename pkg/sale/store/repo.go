package store

import (
	"github.com/Bailey30/accounts/pkg/sale"
	"log"

	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	selectSale      = `SELECT * FROM sale WHERE id=$1`
	selectManySales = `SELECT * FROM sale`
	insertSale      = `INSERT INTO sale (amount, payment_date, created_at, updated_at) VALUES ($1, $2, now(), now()) RETURNING id`
	updateSale      = `UPDATE sale SET amount = $1, payment_date = $2, updated_at = now() WHERE id = $3`
)

type saleRepo struct {
	DB *sqlx.DB
}

func New(db *sqlx.DB) sale.Repo {
	return &saleRepo{db}
}

func (r *saleRepo) Get(id int) (sale.Sale, error) {
	var sale sale.Sale

	err := r.DB.QueryRow(selectSale, id).Scan(&sale.Id, &sale.Amount, &sale.PaymentDate, &sale.CreatedAt, &sale.UpdatedAt)
	if err != nil {
		return sale, fmt.Errorf("Unable to query db:%s", err.Error())
	}

	return sale, nil
}

func (r *saleRepo) GetAll() ([]sale.Sale, error) {
	allSales := make([]sale.Sale, 0)

	rows, err := r.DB.Query(selectManySales)
	if err != nil {
		return allSales, fmt.Errorf("Unable to query db: %s", err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var s sale.Sale
		if err := rows.Scan(&s.Id, &s.Amount, &s.PaymentDate, &s.CreatedAt, &s.UpdatedAt); err != nil {
			return allSales, fmt.Errorf("Unable to scan db rows: %s", err.Error())
		}

		allSales = append(allSales, s)
	}

	return allSales, nil
}

func (r *saleRepo) Create(sale sale.Sale) (string, error) {
	var id string

	if err := r.DB.QueryRow(insertSale, sale.Amount, sale.PaymentDate).Scan(&id); err != nil {
		return "", fmt.Errorf("unable to create sale: %v", err.Error())
	}

	log.Printf("Created sale with id: %s", id)
	return id, nil

}

func (r *saleRepo) Update(sale sale.Sale) error {
	return nil
}

func (r *saleRepo) Delete(id string) error {
	return nil
}
