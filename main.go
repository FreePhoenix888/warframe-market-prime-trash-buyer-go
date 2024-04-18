package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/ztrue/tracerr"
)

type ItemsItem struct {
	ID       string `json:"id"`
	URLName  string `json:"url_name"`
	Thumb    string `json:"thumb"`
	ItemName string `json:"item_name"`
}

type ItemsResponse struct {
	Payload struct {
		Items []ItemsItem `json:"items"`
	} `json:"payload"`
}

type UserStatus string

func (s UserStatus) IsOnline() bool {
	return s == "online"
}

func (s UserStatus) IsOffline() bool {
	return s == "offline"
}

func (s UserStatus) IsIngame() bool {
	return s == "ingame"
}

type Platform string

func (p Platform) IsPC() bool {
	return p == "pc"
}

type UserInfoInOrder struct {
	ID         string     `json:"id"`
	IngameName string     `json:"ingame_name"`
	Status     UserStatus `json:"status"`
	Region     string     `json:"region"`
	Reputation int        `json:"reputation"`
	Avatar     string     `json:"avatar"`
	LastSeen   time.Time  `json:"last_seen"`
}

type Order struct {
	ID           string          `json:"id"`
	Platinum     int             `json:"platinum"`
	Quantity     int             `json:"quantity"`
	OrderType    string          `json:"order_type"`
	Platform     Platform        `json:"platform"`
	Region       string          `json:"region"`
	CreationDate time.Time       `json:"creation_date"`
	LastUpdate   time.Time       `json:"last_update"`
	Subtype      string          `json:"subtype"`
	Visible      bool            `json:"visible"`
	User         UserInfoInOrder `json:"user"`
}

type FullItem struct {
	ID         string `json:"id"`
	ItemsInSet []struct {
		ID             string   `json:"id"`
		URLName        string   `json:"url_name"`
		Icon           string   `json:"icon"`
		IconFormat     string   `json:"icon_format"`
		Thumb          string   `json:"thumb"`
		SubIcon        string   `json:"sub_icon"`
		ModMaxRank     int      `json:"mod_max_rank"`
		Subtypes       []string `json:"subtypes"`
		Tags           []string `json:"tags"`
		Ducats         int      `json:"ducats"`
		QuantityForSet int      `json:"quantity_for_set"`
		SetRoot        bool     `json:"set_root"`
		MasteryLevel   int      `json:"mastery_level"`
		Rarity         string   `json:"rarity"`
		TradingTax     int      `json:"trading_tax"`
		En             struct {
			ItemName    string `json:"item_name"`
			Description string `json:"description"`
			WikiLink    string `json:"wiki_link"`
			Drop        []struct {
				Name string `json:"name"`
				Link string `json:"link"`
			} `json:"drop"`
		} `json:"en"`
		Ru struct {
			ItemName    string `json:"item_name"`
			Description string `json:"description"`
			WikiLink    string `json:"wiki_link"`
			Drop        []struct {
				Name string `json:"name"`
				Link string `json:"link"`
			} `json:"drop"`
		} `json:"ru"`
		Ko struct {
			ItemName    string `json:"item_name"`
			Description string `json:"description"`
			WikiLink    string `json:"wiki_link"`
			Drop        []struct {
				Name string `json:"name"`
				Link string `json:"link"`
			} `json:"drop"`
		} `json:"ko"`
		Fr struct {
			ItemName    string `json:"item_name"`
			Description string `json:"description"`
			WikiLink    string `json:"wiki_link"`
			Drop        []struct {
				Name string `json:"name"`
				Link string `json:"link"`
			} `json:"drop"`
		} `json:"fr"`
		De struct {
			ItemName    string `json:"item_name"`
			Description string `json:"description"`
			WikiLink    string `json:"wiki_link"`
			Drop        []struct {
				Name string `json:"name"`
				Link string `json:"link"`
			} `json:"drop"`
		} `json:"de"`
		Sv struct {
			ItemName    string `json:"item_name"`
			Description string `json:"description"`
			WikiLink    string `json:"wiki_link"`
			Drop        []struct {
				Name string `json:"name"`
				Link string `json:"link"`
			} `json:"drop"`
		} `json:"sv"`
		ZhHant struct {
			ItemName    string `json:"item_name"`
			Description string `json:"description"`
			WikiLink    string `json:"wiki_link"`
			Drop        []struct {
				Name string `json:"name"`
				Link string `json:"link"`
			} `json:"drop"`
		} `json:"zh_hant"`
		ZhHans struct {
			ItemName    string `json:"item_name"`
			Description string `json:"description"`
			WikiLink    string `json:"wiki_link"`
			Drop        []struct {
				Name string `json:"name"`
				Link string `json:"link"`
			} `json:"drop"`
		} `json:"zh_hans"`
		Pt struct {
			ItemName    string `json:"item_name"`
			Description string `json:"description"`
			WikiLink    string `json:"wiki_link"`
			Drop        []struct {
				Name string `json:"name"`
				Link string `json:"link"`
			} `json:"drop"`
		} `json:"pt"`
		Es struct {
			ItemName    string `json:"item_name"`
			Description string `json:"description"`
			WikiLink    string `json:"wiki_link"`
			Drop        []struct {
				Name string `json:"name"`
				Link string `json:"link"`
			} `json:"drop"`
		} `json:"es"`
		Pl struct {
			ItemName    string `json:"item_name"`
			Description string `json:"description"`
			WikiLink    string `json:"wiki_link"`
			Drop        []struct {
				Name string `json:"name"`
				Link string `json:"link"`
			} `json:"drop"`
		} `json:"pl"`
	} `json:"items_in_set"`
}

type OrdersResponse struct {
	Payload struct {
		Orders []Order `json:"orders"`
	} `json:"payload"`
	Include struct {
		Item FullItem `json:"item"`
	} `json:"include"`
}

type GoodOrder struct {
	Order Order
	Item  ItemsItem
}

func GetGoodOrders() ([]GoodOrder, error) {
	log.Println("Fetching items")
	itemsResponseEncoded, err := http.Get("https://api.warframe.market/v1/items")
	if err != nil {
		return nil, fmt.Errorf("failed to get items info: %s", err)
	}
	defer itemsResponseEncoded.Body.Close()

	log.Println("Decoding items response")
	var itemsResponse ItemsResponse
	if err := json.NewDecoder(itemsResponseEncoded.Body).Decode(&itemsResponse); err != nil {
		return nil, fmt.Errorf("failed to decode JSON: %s", err)
	}
	items := itemsResponse.Payload.Items

	goodOrders := make([]GoodOrder, 0, 10)
	for _, item := range items {
		log.Println("Fetching orders for " + item.ItemName)
		ordersResponseEncoded, err := http.Get(fmt.Sprintf("https://api.warframe.market/v1/items/%s/orders", item.URLName))
		if err != nil {
			return nil, fmt.Errorf("error getting orders for %s:%s", item.ItemName, err)
		}
		defer ordersResponseEncoded.Body.Close()

		log.Println("Decoding orders response for " + item.ItemName)
		var ordersResponse OrdersResponse
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

func GenerateMessages(goodOrders []GoodOrder) ([]string, error) {
	messages := make([]string, len(goodOrders))
	for _, goodOrder := range goodOrders {
		log.Println("Generating message for " + goodOrder.Order.ID)
		userName := goodOrder.Order.User.IngameName
		itemName := goodOrder.Item.ItemName
		price := goodOrder.Order.Platinum
		quantity := goodOrder.Order.Quantity
		sum := min(3, price) * quantity
		message := "/w " + userName + " Hi, " + userName + "! You have WTS order: " + itemName + " for " + strconv.Itoa(price) + ". I would like to buy all " + strconv.Itoa(quantity) + " for " + strconv.Itoa(sum) + " if you are interested :)"
		messages = append(messages, message)
	}
	return messages, nil
}

func main() {
	env := os.Getenv("ENVIRONMENT")

	if env == "production" {
		file, err := os.OpenFile("error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Fatal("Failed to open error log file")
		}
		defer file.Close()
		log.SetOutput(file)
		log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	} else {
		log.SetOutput(os.Stdout)
		log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
	}

	goodOrders, err := GetGoodOrders()
	if err != nil {
		log.Fatal("Error getting good orders:", err)
	}

	messages, err := GenerateMessages(goodOrders)
	if err != nil {
		log.Fatal("Failed to generate messages:", err)

	}

	for _, message := range messages {
		fmt.Println(message)
	}

}
