package usecase

import "github.com/leandrosouzasnts/go-intensivo/internal/entity"

type OrderInput struct {
	Id    string
	Price float64
	Tax   float64
}

type OrderOutput struct {
	Id         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

type CalculateFinalPrice struct {
	OrderRepositoryInterface entity.OrderRepositoryInterface
}

func (c *CalculateFinalPrice) Execute(input OrderInput) (*OrderOutput, error) {
	order, err := entity.NewOrder(input.Id, input.Price, input.Tax)
	if err != nil {
		return nil, err
	}

	err = order.CalculateFinalPrice()
	if err != nil {
		return nil, err
	}

	err = c.OrderRepositoryInterface.Save(order)
	if err != nil {
		return nil, err
	}

	return &OrderOutput{
		Id:         order.Id,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
	}, nil
}
