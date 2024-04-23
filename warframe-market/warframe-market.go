package warframemarket

import "time"

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
