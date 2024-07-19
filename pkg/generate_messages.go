package warframe_market_prime_trash_buyer

import (
	"strconv"

	"go.uber.org/zap"
)

func GeneratePurchaseMessages(profitableOrders []OrderWithItem, logger *zap.Logger) ([]string, error) {
	messages := make([]string, 0, len(profitableOrders))
	for _, orderWithItem := range profitableOrders {
		logger.Info("Generating message for order", zap.String("order_id", orderWithItem.Order.ID))
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
