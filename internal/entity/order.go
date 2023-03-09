package entity

import (
	"errors"
)

type Order struct {
	Id         string
	Price      float64
	Tax        float64
	FinalPrice float64
	Discount   float64
}

func NewOrder(id string, price float64, tax float64) (*Order, error) {
	order := &Order{
		Id:    id,
		Price: price,
		Tax:   tax,
	}

	err := order.Validate()

	if err != nil {
		return nil, err
	}

	return order, nil
}

func (order *Order) Validate() error {
	if order.Id == "" {
		return errors.New("id is required")
	}

	if order.Price <= 0 {
		return errors.New("invalid price")
	}

	if order.Tax <= 0 {
		return errors.New("invalid tax")
	}

	return nil
}

func (order *Order) CalculateFinalPrice() error {
	order.FinalPrice = order.Price + order.Tax

	err := order.Validate()

	if err != nil {
		return err
	}

	return nil
}

func (order *Order) CalculateFinalPriceWithDiscount(discount float64) {
	finalPrice := order.Price + order.Tax

	finalPriceWithDiscount := (finalPrice * discount) / 100

	order.Discount = finalPriceWithDiscount

}
