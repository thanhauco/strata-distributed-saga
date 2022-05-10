package order

import (
	"context"
	"fmt"
)

type Service struct {
	orders map[string]*Order
}

func NewService() *Service {
	return &Service{orders: make(map[string]*Order)}
}

func (s *Service) CreateOrder(ctx context.Context, o *Order) error {
	if o.Amount <= 0 {
		return fmt.Errorf("order amount must be positive")
	}
	o.Status = StatusSubmitted
	o.Version = 1
	s.orders[o.ID] = o
	return nil
}
