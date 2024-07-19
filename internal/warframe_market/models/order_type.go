package models

type OrderType string

func (orderType OrderType) IsSell() bool {
	return orderType == "sell"
}
