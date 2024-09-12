//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/EleyOliveira/go-clean-arch/internal/entity"
	"github.com/EleyOliveira/go-clean-arch/internal/infra/database"
	"github.com/EleyOliveira/go-clean-arch/internal/infra/web"
	"github.com/EleyOliveira/go-clean-arch/internal/usecase"
	"github.com/google/wire"
)

var setOrderRepositoryDependency = wire.NewSet(
	database.NewOrderRepository,
	wire.Bind(new(entity.OrderRepositoryInterface), new(*database.OrderRepository)),
)

func NewListOrderUseCase(db *sql.DB) *usecase.ListOrderUseCase {
	wire.Build(setOrderRepositoryDependency, usecase.NewListOrderUseCase)
	return &usecase.ListOrderUseCase{}
}

func NewWebOrderHandler(db *sql.DB) *web.WebOrderHandler {
	wire.Build(
		setOrderRepositoryDependency,
		web.NewWebOrderHandler,
	)
	return &web.WebOrderHandler{}
}
