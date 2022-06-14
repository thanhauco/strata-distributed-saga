package inventory

import (
	"context"
	"fmt"
	"sync"
)

type StockService struct {
	mu    sync.Mutex
	stock map[string]int
}

func NewStockService() *StockService {
	s := &StockService{stock: make(map[string]int)}
	s.stock["LAPTOP-PRO-16"] = 50
	s.stock["HEADPHONES-ANC"] = 100
	s.stock["PHONE-MAX"] = 0 // Out of stock trigger
	return s
}

func (s *StockService) Reserve(ctx context.Context, sku string, qty int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	avail, exists := s.stock[sku]
	if !exists || avail < qty {
		return fmt.Errorf("INSUFFICIENT_STOCK: sku %s only has %d units", sku, avail)
	}
	s.stock[sku] -= qty
	return nil
}

func (s *StockService) Release(ctx context.Context, sku string, qty int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.stock[sku] += qty
}
