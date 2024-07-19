package models

import "time"

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
	OrderType    OrderType       `json:"order_type"`
	Platform     Platform        `json:"platform"`
	Region       string          `json:"region"`
	CreationDate time.Time       `json:"creation_date"`
	LastUpdate   time.Time       `json:"last_update"`
	Subtype      string          `json:"subtype"`
	Visible      bool            `json:"visible"`
	User         UserInfoInOrder `json:"user"`
}

type OrdersResponse struct {
	Payload struct {
		Orders []Order `json:"orders"`
	} `json:"payload"`
	Include struct {
		Item FullItem `json:"item"`
	} `json:"include"`
}
