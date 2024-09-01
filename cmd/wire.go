//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/EleyOliveira/go-clean-arch/internal/entity"
	"github.com/EleyOliveira/go-clean-arch/internal/usecase"
	"github.com/google/wire"
)

func NewListOrderUseCase(db *sql.DB) *usecase.ListOrderUseCase {
	wire.Build(entity.NewOrderRepository, usecase.NewListOrderUseCase)
	return &usecase.ListOrderUseCase{}
}
