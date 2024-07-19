package warframe_market_prime_trash_buyer

import (
	"os"
	"testing"

	"go.uber.org/zap"
)

var logger *zap.Logger

func TestMain(m *testing.M) {
	var err error
	testLogger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	code := m.Run()
	testLogger.Sync()
	os.Exit(code)
}

func TestGetProfitableOrders(t *testing.T) {
	profitableOrders, err := GetProfitableOrders(logger)

	if err != nil {
		t.Errorf("GetProfitableOrders() returned an error: %v", err)
	}

	if profitableOrders == nil {
		t.Error("GetProfitableOrders() returned nil profitableOrders slice")
	}

}
