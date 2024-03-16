package repository

import (
	"database/sql"
	"errors"

	"are.moe/testtask/internal/domain"
)

type ProductRepo interface {
	Create(input domain.ProductInput) (domain.Product, error)
	GetByID(id domain.ProductID) (domain.Product, error)
	Update(id domain.ProductID, input domain.ProductInput) (domain.Product, error)
	Delete(id domain.ProductID) error
	List(offset domain.ProductID, limit domain.ProductPageSize) ([]domain.Product, error)
	FindByName(name string, offset domain.ProductID, limit domain.ProductPageSize) ([]domain.Product, error)
}

type Product struct {
	db *sql.DB
}

func NewProductRepo(db *sql.DB) *Product {
	return &Product{db}
}

func rowsToProducts(rows *sql.Rows, length domain.ProductPageSize) ([]domain.Product, error) {
	res := make([]domain.Product, 0, length)

	for rows.Next() {
		var product domain.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.CreatedAt); err != nil {
			return nil, err
		}
		res = append(res, product)
	}

	return res, nil
}

func (r *Product) Create(input domain.ProductInput) (domain.Product, error) {
	var product domain.Product
	row := r.db.QueryRow(
		`INSERT INTO "Products" ("Name", "Description", "Price")
			VALUES ($1, $2, $3)
			RETURNING "ID", "Name", "Description", "Price", "CreatedAt"`,
		input.Name, input.Description, input.Price,
	)

	if err := row.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.CreatedAt); err != nil {
		return product, err
	}

	return product, nil
}

func (r *Product) GetByID(id domain.ProductID) (domain.Product, error) {
	var product domain.Product
	row := r.db.QueryRow(
		`SELECT "ID", "Name", "Description", "Price", "CreatedAt"
			FROM "Products"
			WHERE "ID" = $1`,
		id)

	err := row.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.CreatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return product, domain.ErrProductNotFound
	}
	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *Product) Update(id domain.ProductID, input domain.ProductInput) (domain.Product, error) {
	var product domain.Product
	row := r.db.QueryRow(
		`UPDATE "Products"
			SET "Name"=$1, "Description"=$2, "Price"=$3
			WHERE "ID" = $4
			RETURNING "ID", "Name", "Description", "Price", "CreatedAt"`,
		input.Name, input.Description, input.Price, id)

	err := row.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.CreatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return product, domain.ErrProductNotFound
	}
	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *Product) Delete(id domain.ProductID) error {
	var product domain.Product
	row := r.db.QueryRow(
		`DELETE FROM "Products"
			WHERE "ID" = $1
			RETURNING "ID", "Name", "Description", "Price", "CreatedAt"`,
		id)

	err := row.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.CreatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return domain.ErrProductNotFound
	}
	return err
}

func (r *Product) List(offset domain.ProductID, limit domain.ProductPageSize) ([]domain.Product, error) {
	rows, err := r.db.Query(
		`SELECT "Products"."ID", "Products"."Name", "Products"."Description", "Products"."Price", "Products"."CreatedAt"
			FROM "Products" JOIN (
				SELECT "ID" FROM "Products"
				ORDER BY "ID"
				OFFSET $1 LIMIT $2
			) AS ids ON ids."ID" = "Products"."ID"`,
		offset, limit,
	)
	if err != nil {
		return nil, err
	}

	return rowsToProducts(rows, limit)
}

func (r *Product) FindByName(name string, offset domain.ProductID, limit domain.ProductPageSize) ([]domain.Product, error) {
	rows, err := r.db.Query(
		`SELECT "Products"."ID", "Products"."Name", "Products"."Description", "Products"."Price", "Products"."CreatedAt"
			FROM "Products" JOIN (
			SELECT "ID" FROM "Products"
				WHERE "Name" = $1
				ORDER BY "ID"
				OFFSET $2 LIMIT $3
			) AS ids ON ids."ID" = "Products"."ID"`,
		name, offset, limit,
	)
	if err != nil {
		return nil, err
	}

	return rowsToProducts(rows, limit)
}
