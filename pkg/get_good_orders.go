package warframe_market_prime_trash_buyer

import (
	"encoding/json"
	"fmt"
	"net/http"
	"slices"
	"time"

	warframe_market "github.com/freephoenix888/warframe-market-prime-trash-buyer-go-lib/internal/warframe_market"
	warframe_market_models "github.com/freephoenix888/warframe-market-prime-trash-buyer-go-lib/internal/warframe_market/models"
)

func GetProfitableOrders() ([]OrderWithItem, error) {
	itemsResponseEncoded, err := http.Get("https://api.warframe.market/v1/items")
	if err != nil {
		return nil, fmt.Errorf("failed to get items info: %s", err)
	}
	defer itemsResponseEncoded.Body.Close()

	var itemsResponse warframe_market_models.ItemsResponse
	if err := json.NewDecoder(itemsResponseEncoded.Body).Decode(&itemsResponse); err != nil {
		return nil, fmt.Errorf("failed to decode JSON: %s", err)
	}
	items := itemsResponse.Payload.Items

	var itemsToBuy []warframe_market_models.ItemsItem
	for _, item := range items {
		if slices.Contains(ProfitableItemNames, item.ItemName) {
			itemsToBuy = append(itemsToBuy, item)
		}
	}

	profitableOrders := make([]OrderWithItem, 0, 10)

	rateLimiter := time.NewTicker(time.Second / warframe_market.MaxRequestsPerSecond)
	defer rateLimiter.Stop()

	for _, itemToBuy := range itemsToBuy {
		<-rateLimiter.C

		ordersResponseEncoded, err := http.Get(fmt.Sprintf("https://api.warframe.market/v1/items/%s/orders", itemToBuy.URLName))
		if err != nil {
			return nil, fmt.Errorf("error getting orders for %s: %s", itemToBuy.ItemName, err)
		}
		defer ordersResponseEncoded.Body.Close()

		var ordersResponse warframe_market_models.OrdersResponse
		if err := json.NewDecoder(ordersResponseEncoded.Body).Decode(&ordersResponse); err != nil {
			return nil, fmt.Errorf("failed to decode JSON: %s", err)
		}

		orders := ordersResponse.Payload.Orders

		for _, order := range orders {
			isSellOrder := order.OrderType.IsSell()
			if !isSellOrder {
				continue
			}

			isIngame := order.User.Status.IsIngame()
			if !isIngame {
				continue
			}

			isPc := order.Platform.IsPC()
			if !isPc {
				continue
			}

			isGoodQuantity := order.Quantity >= 3
			if !isGoodQuantity {
				continue
			}

			isGoodPrice := order.Platinum <= 4
			if !isGoodPrice {
				continue
			}

			orderWithItem := OrderWithItem{
				Order: order,
				Item:  itemToBuy,
			}
			profitableOrders = append(profitableOrders, orderWithItem)
		}
	}

	return profitableOrders, nil
}
