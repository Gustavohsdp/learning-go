package entity

import "errors"

type Order struct {
	Id         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

func NewOrder(id string, price float64, tax float64) (*Order, error) {
	order := &Order{
		Id:    id,
		Price: price,
		Tax:   tax,
	}

	err := order.validate()

	if err != nil {
		return nil, err
	}

	return order, nil
}

func (order *Order) validate() error {
	if order.Id == "" {
		return errors.New("Id is required")
	}

	if order.Price <= 0 {
		return errors.New("Invalid price")
	}

	if order.Tax <= 0 {
		return errors.New("Invalid tax")
	}

	return nil
}

func (order *Order) CalculateFinalPrice() error {
	order.FinalPrice = order.Price + order.Tax

	err := order.validate()

	if err != nil {
		return err
	}

	return nil
}
