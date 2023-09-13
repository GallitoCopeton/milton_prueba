package product

import (
	"challenge/database"
	"fmt"
	"github.com/jmoiron/sqlx"
)

const TableName = "products"

type Repository struct {
	db *database.Database
}

func NewRepository(db *database.Database) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetProducts() (Products, error) {
	query := fmt.Sprintf(
		"SELECT * FROM %v", TableName,
	)

	updatedQuery, args, err := sqlx.In(query)
	if err != nil {
		return nil, err
	}

	updatedQuery = r.db.Rebind(updatedQuery)

	var products Products
	if err = r.db.Select(&products, updatedQuery, args...); err != nil {
		return nil, fmt.Errorf("unable to query products %w", err)
	}

	return products, nil
}
