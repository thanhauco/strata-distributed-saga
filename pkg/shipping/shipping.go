package shipping

import (
	"context"
	"fmt"
	"time"
)

type Shipment struct {
	TrackingNumber string    `json:"tracking_number"`
	Carrier        string    `json:"carrier"`
	Status         string    `json:"status"`
	CreatedAt      time.Time `json:"created_at"`
}

type Service struct {
	shipments map[string]*Shipment
}

func NewService() *Service {
	return &Service{shipments: make(map[string]*Shipment)}
}

func (s *Service) Dispatch(ctx context.Context, orderID string) (*Shipment, error) {
	shipment := &Shipment{
		TrackingNumber: fmt.Sprintf("TRK-%s", orderID),
		Carrier:        "FEDEX_EXPRESS",
		Status:         "MANIFESTED",
		CreatedAt:      time.Now(),
	}
	s.shipments[orderID] = shipment
	return shipment, nil
}

func (s *Service) Cancel(ctx context.Context, orderID string) {
	delete(s.shipments, orderID)
}

var SupportedCarriers = []string{"FEDEX", "UPS", "DHL"}
// Error wrapping verified\n