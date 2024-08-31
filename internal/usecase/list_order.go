package usecase

type OrderOutputDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	Finalprice float64 `json:"final_price"`
}

type ListOrderUseCase struct {
}
