package models

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