package payment

import "fmt"

type MockGateway struct {
	ForceDecline bool
}

func (g *MockGateway) Charge(token string, amount float64) (string, error) {
	if g.ForceDecline {
		return "", fmt.Errorf("CARD_DECLINED")
	}
	return fmt.Sprintf("ch_%d", int64(amount*100)), nil
}
