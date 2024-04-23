package lib

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"slices"
	"strconv"

	"github.com/freephoenix888/warframe-market-prime-trash-buyer-go/data"
	warframemarket "github.com/freephoenix888/warframe-market-prime-trash-buyer-go/warframe-market"
	"github.com/ztrue/tracerr"
)

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

	goodOrders := make([]GoodOrder, 0, 10)
	for _, item := range items {
		if !slices.Contains(data.ItemNamesToBuy, item.ItemName) {
			log.Println("Ignoring " + item.ItemName + " because it is not in the list of items we should buy")
			continue
		}

		log.Println("Fetching orders for " + item.ItemName)
		ordersResponseEncoded, err := http.Get(fmt.Sprintf("https://api.warframe.market/v1/items/%s/orders", item.URLName))
		if err != nil {
			return nil, fmt.Errorf("error getting orders for %s: %s", item.ItemName, err)
		}
		defer ordersResponseEncoded.Body.Close()

		log.Println("Decoding orders response for " + item.ItemName)
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
				Item:  item,
			}
			log.Println("Found good order!")
			goodOrders = append(goodOrders, goodOrder)
		}
	}

	log.Println("Found " + strconv.Itoa(len(goodOrders)) + " good orders")
	return goodOrders, nil
}
