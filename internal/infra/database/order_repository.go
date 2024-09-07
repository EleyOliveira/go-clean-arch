package database

import (
	"context"
	"database/sql"

	"github.com/EleyOliveira/go-clean-arch/internal/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{Db: db}
}

func (o *OrderRepository) ListOrders() ([]entity.Order, error) {

	ctx := context.Background()
	queries := entity.New(o.Db)

	orders, err := queries.ListOrders(ctx)
	if err != nil {
		return nil, err
	}

	return orders, nil
}
