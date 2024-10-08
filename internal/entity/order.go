package entity

import "errors"

func NewOrder(id string, price float64, tax float64) (*Order, error) {
	order := &Order{
		ID:    id,
		Price: price,
		Tax:   tax,
	}

	err := order.IsValid()
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (o *Order) IsValid() error {

	if o.ID == "" {
		return errors.New("invalid id")
	}

	if o.Price <= 0 {
		return errors.New("invalid price")
	}

	if o.Tax <= 0 {
		return errors.New("invalid tax")
	}
	return nil
}

func (o *Order) CalculateFinalPrice() error {
	err := o.IsValid()
	if err != nil {
		return err
	}
	o.Finalprice = o.Price + o.Tax
	return nil
}
