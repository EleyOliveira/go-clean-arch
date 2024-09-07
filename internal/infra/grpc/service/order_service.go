package service

import (
	"github.com/EleyOliveira/go-clean-arch/internal/infra/grpc/pb"
	"github.com/EleyOliveira/go-clean-arch/internal/usecase"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	ListOrderUseCase usecase.ListOrderUseCase
}
