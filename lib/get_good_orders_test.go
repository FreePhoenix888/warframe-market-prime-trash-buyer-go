package lib

import (
	"testing"
)

func TestGetGoodOrders(t *testing.T) {
	goodOrders, err := GetGoodOrders()

	if err != nil {
		t.Errorf("GetGoodOrders() returned an error: %v", err)
	}

	if goodOrders == nil {
		t.Error("GetGoodOrders() returned nil goodOrders slice")
	}

}
