package web

import (
	"encoding/json"
	"net/http"

	"github.com/EleyOliveira/go-clean-arch/internal/entity"
	"github.com/EleyOliveira/go-clean-arch/internal/usecase"
)

type WebOrderHandler struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewWebOrderHandler(OrderRepository entity.OrderRepositoryInterface) *WebOrderHandler {
	return &WebOrderHandler{
		OrderRepository: OrderRepository,
	}
}

func (h *WebOrderHandler) List(w http.ResponseWriter, r *http.Request) {
	var dto usecase.OrderOutputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	listorder := usecase.NewListOrderUseCase(h.OrderRepository)
	output, err := listorder.ListOrders()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
