package graph

import "github.com/EleyOliveira/go-clean-arch/internal/usecase"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	ListOrderUseCase   usecase.ListOrderUseCase
	CreateOrderUseCase usecase.CreateOrderUseCase
}