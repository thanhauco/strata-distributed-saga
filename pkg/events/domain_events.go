package events

type OrderCreatedEvent struct {
	OrderID    string  `json:"order_id"`
	CustomerID string  `json:"customer_id"`
	Amount     float64 `json:"amount"`
	Currency   string  `json:"currency"`
	SKU        string  `json:"sku"`
	Quantity   int     `json:"quantity"`
}

type PaymentReservedEvent struct {
	SagaID        string  `json:"saga_id"`
	TransactionID string  `json:"transaction_id"`
	Amount        float64 `json:"amount"`
}

type InventoryReservedEvent struct {
	SagaID      string `json:"saga_id"`
	WarehouseID string `json:"warehouse_id"`
	SKU         string `json:"sku"`
	Quantity    int    `json:"quantity"`
}

func IsHighValueOrder(amount float64) bool {
	return amount >= 5000.0
}
