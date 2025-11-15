package breeze

import (
	"encoding/json"
	"fmt"
)

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
	OrderId    string `json:"order_id"`
	Message    string `json:"message"`
	UserRemark string `json:"user_remark"`
}

type PlaceOrderResponse struct {
	Success PlaceOrderSuccess `json:"Success"`
	Status  string            `json:"Status"`
	Error   any               `json:"Error"`
}

func (bc *BreezeClient) PlaceOrder(order PlaceOrderRequest) (PlaceOrderSuccess, error) {
	res, err := bc.request("POST", "order", order)
	if err != nil {
		return PlaceOrderSuccess{}, fmt.Errorf("There was an error getting response: %v", err)
	}
	resBody := &PlaceOrderResponse{}
	byteResponse, ok := res.([]byte)
	if ok {
		err = json.Unmarshal(byteResponse, resBody)
		if err != nil {
			return PlaceOrderSuccess{}, fmt.Errorf("There was error unmarshalling the response: %v", err)
		}
	}
	if resBody.Error != nil {
		return PlaceOrderSuccess{}, fmt.Errorf("Some Error Occured Server Side. Status: %v, Error: %v", resBody.Status, resBody.Error)
	}
	return resBody.Success, nil
}
