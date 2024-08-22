package types

type ItemDetails struct {
	Name     string  `json:"item_name"`
	FloatVal float64 `json:"float_value"`
	Wear     string  `json:"wear_name"`
	Stickers []any   `json:"stickers"`
}

type SearchConfig struct {
	MaxPrice         int
	MinPrice         int
	MinDiscountValue int
	Category         int //1 normal
	Gun              string
	Keyfile          string
}

type FilterConfig struct {
	MinDiscountPercentage float64
	Stickers              bool
	Top                   int
}

func NewFilterConfig(config *InputConfig) FilterConfig {
	return FilterConfig{
		MinDiscountPercentage: config.MinDiscountPercentage,
		Stickers:              config.Stickers,
		Top:                   config.Top,
	}
}

type InputConfig struct {
	SearchConfig
	FilterConfig
}

type Ref struct {
	BasePrice      int64   `json:"base_price"`
	FloatFactor    float64 `json:"float_factor"`
	PredictedPrice int64   `json:"predicted_price"`
}

type Gun struct {
	Name     string
	Wear     string
	Stickers []any
}

type ItemFilter[B BuyNowItem | AuctionItem] interface {
	Filter(b []B, filterConfig FilterConfig) []FilteredItem
}

type FilteredItem struct {
	Item
	TerminatesIn string
}
type BuyNowItem struct {
	Item
}

type Item struct {
	Price       int64  `json:"price"`
	Id          string `json:"id"`
	ItemDetails `json:"item"`
	Ref         `json:"reference"`
}

type AuctionItem struct {
	Item
	AuctionDetails `json:"auction_details"`
}

type Writer interface {
	WriteOut(result []OutputItem)
}

type AuctionDetails struct {
	ReservePrice int `json:"reserve_price"`
	TopBid       `json:"top_bid"`
	ExpiresAt    string `json:"expires_at"`
	MinNextBid   int    `json:"min_next_bid"`
}

type TopBid struct {
	Price int    `json:"price"`
	State string `json:"state"`
}

type OutputItem struct {
	Type         string
	Name         string
	Wear         string
	Price        string
	Discount     string
	DiscountPerc string
	MarketPrice  string
	TerminatesIn string
	Link         string
	Image        any
}
