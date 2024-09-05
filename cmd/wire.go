//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/EleyOliveira/go-clean-arch/internal/entity"
	"github.com/EleyOliveira/go-clean-arch/internal/infra/database"
	"github.com/EleyOliveira/go-clean-arch/internal/usecase"
	"github.com/google/wire"
)

var setRepositoryDependency = wire.NewSet(
	database.NewOrderRepository,
	wire.Bind(new(entity.OrderRepositoryInterface), new(*database.OrderRepository)),
)

func NewListOrderUseCase(db *sql.DB) *usecase.ListOrderUseCase {
	wire.Build(setRepositoryDependency, usecase.NewListOrderUseCase)
	return &usecase.ListOrderUseCase{}
}
