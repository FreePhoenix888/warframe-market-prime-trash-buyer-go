package models

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