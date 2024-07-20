package warframe_market_prime_trash_buyer

import (
	"testing"
)

func TestGetProfitableOrders(t *testing.T) {
	profitableOrders, err := GetProfitableOrders()

	if err != nil {
		t.Errorf("GetProfitableOrders() returned an error: %v", err)
	}

	if profitableOrders == nil {
		t.Error("GetProfitableOrders() returned nil profitableOrders slice")
	}

}
