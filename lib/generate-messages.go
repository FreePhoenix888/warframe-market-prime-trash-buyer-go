package lib

import (
	"log"
	"strconv"
)

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
