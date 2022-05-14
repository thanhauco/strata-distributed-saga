package payment

import (
	"context"
	"fmt"
)

type Service struct {
	reservations map[string]float64
}

func NewService() *Service {
	return &Service{reservations: make(map[string]float64)}
}

func (s *Service) ReservePayment(ctx context.Context, sagaID string, amount float64) error {
	if amount > 10000.0 {
		return fmt.Errorf("transaction limit exceeded: %.2f", amount)
	}
	s.reservations[sagaID] = amount
	return nil
}

func (s *Service) RefundPayment(ctx context.Context, sagaID string) error {
	delete(s.reservations, sagaID)
	return nil
}
