package warframe_market_prime_trash_buyer

import (
	"strconv"
)

func GeneratePurchaseMessages(profitableOrders []OrderWithItem) ([]string, error) {
	messages := make([]string, 0, len(profitableOrders))
	for _, orderWithItem := range profitableOrders {
		userName := orderWithItem.Order.User.IngameName
		itemName := orderWithItem.Item.ItemName
		price := orderWithItem.Order.Platinum
		quantity := orderWithItem.Order.Quantity
		sum := min(3, price) * quantity
		message := "/w " + userName + " Hi, " + userName + "! You have WTS order: " + itemName + " for " + strconv.Itoa(price) + ". I would like to buy all " + strconv.Itoa(quantity) + " for " + strconv.Itoa(sum) + " if you are interested :)"
		messages = append(messages, message)
	}
	return messages, nil
}
