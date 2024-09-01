// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"database/sql"
	"github.com/EleyOliveira/go-clean-arch/internal/entity"
	"github.com/EleyOliveira/go-clean-arch/internal/usecase"
)

import (
	_ "github.com/go-sql-driver/mysql"
)

// Injectors from wire.go:

func NewListOrderUseCase(db *sql.DB) *usecase.ListOrderUseCase {
	orderRepository := entity.NewOrderRepository(db)
	listOrderUseCase := usecase.NewListOrderUseCase(orderRepository)
	return listOrderUseCase
}
