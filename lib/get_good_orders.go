package lib

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"slices"
	"strconv"
	"time"

	"github.com/freephoenix888/warframe-market-prime-trash-buyer-go/data"
	warframemarket "github.com/freephoenix888/warframe-market-prime-trash-buyer-go/warframe_market"
	"github.com/ztrue/tracerr"
)

// RateLimit defines the number of requests allowed per second.
const RateLimit = 3

func GetGoodOrders() ([]GoodOrder, error) {
	log.Println("Fetching items")
	itemsResponseEncoded, err := http.Get("https://api.warframe.market/v1/items")
	if err != nil {
		return nil, fmt.Errorf("failed to get items info: %s", err)
	}
	defer itemsResponseEncoded.Body.Close()

	log.Println("Decoding items response")
	var itemsResponse warframemarket.ItemsResponse
	if err := json.NewDecoder(itemsResponseEncoded.Body).Decode(&itemsResponse); err != nil {
		return nil, fmt.Errorf("failed to decode JSON: %s", err)
	}
	items := itemsResponse.Payload.Items

	var itemsToBuy []warframemarket.ItemsItem
	for _, item := range items {
		if slices.Contains(data.ItemNamesToBuy, item.ItemName) {
			itemsToBuy = append(itemsToBuy, item)
		}
	}

	goodOrders := make([]GoodOrder, 0, 10)

	// Create a ticker to limit requests
	ticker := time.Tick(time.Second / RateLimit)

	for _, itemToBuy := range itemsToBuy {

		log.Println("Fetching orders for " + itemToBuy.ItemName)

		// Wait for the next tick to limit requests
		<-ticker

		ordersResponseEncoded, err := http.Get(fmt.Sprintf("https://api.warframe.market/v1/items/%s/orders", itemToBuy.URLName))
		if err != nil {
			return nil, fmt.Errorf("error getting orders for %s: %s", itemToBuy.ItemName, err)
		}
		defer ordersResponseEncoded.Body.Close()

		log.Println("Decoding orders response for " + itemToBuy.ItemName)
		var ordersResponse warframemarket.OrdersResponse
		if err := json.NewDecoder(ordersResponseEncoded.Body).Decode(&ordersResponse); err != nil {
			return nil, tracerr.Errorf("failed to decode JSON: %s", err)
		}

		orders := ordersResponse.Payload.Orders

		log.Println("Looking for good orders...")
		for _, order := range orders {
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

			goodOrder := GoodOrder{
				Order: order,
				Item:  itemToBuy,
			}
			log.Println("Found good order!")
			goodOrders = append(goodOrders, goodOrder)
		}
	}

	log.Println("Found " + strconv.Itoa(len(goodOrders)) + " good orders")
	return goodOrders, nil
}
