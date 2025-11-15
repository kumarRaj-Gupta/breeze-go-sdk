package breeze

// Place Order
type PlaceOrderRequest struct {
	StockCode    string `json:"stock_code"`
	ExchangeCode string `json:"exchange_code"`
	Product      string `json:"product"`
	Action       string `json:"action"`
	OrderType    string `json:"order_type"`
	Quantity     string `json:"quantity"`
	Price        string `json:"price"`
	ExpiryDate   string `json:"expiry_date"`
	Right        string `json:"right"`
	StrikePrice  string `json:"strike_price"`
	// Not Mandatory
	StopLoss          float64 `json:"stoploss"`
	Validity          string  `json:"validity"`
	ValidityDate      string  `json:"validity_date"`
	DisclosedQuantity string  `json:"disclosed_quantity"`
	UserRemark        string  `json:"user_remark"`
}

type PlaceOrderSuccess struct {
	Datetime     string `json:"datetime"`
	StockCode    string `json:"stock_code"`
	ExchangeCode string `json:"exchange_code"`
	ProductType  string `json:"product_type"`
	ExpiryDate   string `json:"expiry_date"`
	Right        string `json:"right"`
	StrikePrice  string `json:"strike_price"`
	Open         string `json:"open"`
	High         string `json:"high"`
	Low          string `json:"low"`
	Close        string `json:"close"`
	Volume       string `json:"volume"`
	OpenInterest string `json:"open_interest"`
	Count        int    `json:"count"`
}

type PlaceOrderResponse struct {
	Success []PlaceOrderSuccess `json:"Success"`
	Status  string              `json:"Status"`
	Error   string              `json:"Error"`
}

func (bc *BreezeClient) PlaceOrder(order PlaceOrderRequest) (PlaceOrderResponse, error) {

	return nil, nil
}
