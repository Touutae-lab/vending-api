package repository

import (
	"database/sql"
	"github.com/go-jet/jet/v2/postgres"
	"github.com/touutae-lab/vending-api/internal/database/vending_machine/vending_machine_service/model"
	"github.com/touutae-lab/vending-api/internal/database/vending_machine/vending_machine_service/table"
	"golang.org/x/net/context"
)

type ProductRepositoryImpl struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepositoryImpl {
	return &ProductRepositoryImpl{
		db: db,
	}
}

func (pr *ProductRepositoryImpl) GetAllProduct(ctx context.Context) ([]model.Product, error) {
	query := postgres.SELECT(
		table.Product.AllColumns,
	).FROM(
		table.Product,
	)

	queryString, args := query.Sql()

	rows, err := pr.db.QueryContext(ctx, queryString, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []model.Product
	for rows.Next() {
		var p model.Product
		if err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.Price,
			&p.ImgURL,
			&p.Details,
		); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func (pr *ProductRepositoryImpl) GetProductByID(ctx context.Context, id int32) (model.Product, error) {
	query := postgres.SELECT(
		table.Product.AllColumns,
	).FROM(
		table.Product,
	).WHERE(
		table.Product.ID.EQ(postgres.Int32(id)),
	)

	queryString, args := query.Sql()

	row := pr.db.QueryRowContext(ctx, queryString, args...)

	var p model.Product
	if err := row.Scan(
		&p.ID,
		&p.Name,
		&p.Price,
		&p.ImgURL,
		&p.Details,
	); err != nil {
		return model.Product{}, err
	}

	return p, nil
}

func (pr *ProductRepositoryImpl) CreateProduct(ctx context.Context, product *model.Product) (int64, error) {
	query := table.Product.INSERT(
		table.Product.Name,
		table.Product.Price,
		table.Product.ImgURL,
		table.Product.Details,
	).VALUES(
		product.Name,
		product.Price,
		product.ImgURL,
		product.Details,
	)

	queryString, args := query.Sql()

	result, err := pr.db.ExecContext(ctx, queryString, args...)
	if err != nil {
		return -1, err
	}

	return result.LastInsertId()
}

func (pr *ProductRepositoryImpl) UpdateProduct(ctx context.Context, product *model.Product) (int64, error) {
	query := table.Product.UPDATE(
		table.Product.Name,
		table.Product.Price,
		table.Product.ImgURL,
		table.Product.Details,
	).SET(
		table.Product.Name.SET(postgres.String(product.Name)),
		table.Product.Price.SET(postgres.Float(product.Price)),
		table.Product.ImgURL.SET(postgres.String(product.ImgURL)),
		table.Product.Details.SET(postgres.String(product.Details)),
	).WHERE(
		table.Product.ID.EQ(postgres.Int32(product.ID)),
	)

	queryString, args := query.Sql()

	result, err := pr.db.ExecContext(ctx, queryString, args...)
	if err != nil {
		return -1, err
	}

	return result.RowsAffected()
}

func (pr *ProductRepositoryImpl) DeleteProduct(ctx context.Context, id int32) (int64, error) {
	query := table.Product.DELETE().WHERE(
		table.Product.ID.EQ(postgres.Int32(id)),
	)

	queryString, args := query.Sql()

	result, err := pr.db.ExecContext(ctx, queryString, args...)
	if err != nil {
		return -1, err
	}

	return result.RowsAffected()
}
