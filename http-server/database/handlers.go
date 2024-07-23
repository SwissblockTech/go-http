package database

import (
	"context"
	"database/sql"
)

func GetProducts(db *sql.DB, start, count int, ctx context.Context) ([]*Product, error) {
	rows, err := db.QueryContext(ctx, getProductsQuery, count, start)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := make([]*Product, 0)
	for rows.Next() {
		var prod Product
		if err := rows.Scan(&prod.ID, &prod.Name, &prod.Price); err != nil {
			return nil, err
		}
		products = append(products, &prod)
	}

	return products, nil
}

func GetProduct(db *sql.DB, product *Product, ctx context.Context) error {
	return db.QueryRowContext(ctx, getProductQuery, product.ID).
		Scan(&product.Name, &product.Price)
}

func CreateProduct(db *sql.DB, product *Product, ctx context.Context) error {
	err := db.QueryRowContext(ctx, createProductQuery, product.Name, product.Price).Scan(&product.ID)
	if err != nil {
		return err
	}
	return nil
}

func UpdateProduct(db *sql.DB, product *Product, ctx context.Context) error {
	_, err := db.ExecContext(ctx, updateProductQuery, product.Name, product.Price, product.ID)
	return err
}

func DeleteProduct(db *sql.DB, productId int, ctx context.Context) error {
	_, err := db.ExecContext(ctx, deleteProductQuery, productId)
	return err
}
