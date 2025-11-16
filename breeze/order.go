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
	Stoploss          string `json:"stoploss"`
	Validity          string `json:"validity"`
	ValidityDate      string `json:"validity_date"`
	DisclosedQuantity string `json:"disclosed_quantity"`
	UserRemark        string `json:"user_remark"`
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

// Get Order Details
type OrderDetailsRequest struct {
	ExchangeCode string `json:"exchange_code"`
	OrderId      string `json:"order_id"`
}
type OrderDetailsResponse struct {
	Success []OrderDetailsSuccess `json:"Success"`
	Status  string                `json:"Status"`
	Error   any
}
type OrderDetailsSuccess struct {
	OrderId                       string `json:"order_id"`
	ExchangeOrderId               string `json:"exchange_order_id"`
	ExchangeCode                  string `json:"exchange_code"`
	StockCode                     string `json:"stock_code"`
	ProdcutType                   string `json:"product_type"`
	Action                        string `json:"action"`
	OrderType                     string `json:"order_type"`
	Stoploss                      string `json:"stoploss"`
	Quantity                      string `json:"quantity"`
	Price                         string `json:"price"`
	Validity                      string `json:"validity"`
	DisclosedQuantity             string `json:"disclosed_quantity"`
	ExpiryDate                    string `json:"expiry_date"`
	Right                         string `json:"right"`
	StrikePrice                   string `json:"strike_price"`
	AveragePrice                  string `json:"average_price"`
	CancelledQuantity             string `json:"cancelled_quantity"`
	PendingQuantity               string `json:"pending_quantity"`
	Status                        string `json:"status"`
	UserRemark                    string `json:"user_remark"`
	OrderDatetime                 string `json:"order_datetime"`
	ParentOrderId                 string `json:"parent_order_id"`
	ModiificationNumber           string `json:"modification_number"`
	ExchangeAcknowledgementDate   string `json:"exchange_acknowledgement_date"`
	ExchangeAcknowledgementNumber string `json:"exchange_acknowledge_number"`
	SLIPPrice                     string `json:"SLTP_price"`
	InitialLimit                  string `json:"initial_limit"`
	LTP                           string `json:"LTP"`
	LimitOffset                   string `json:"limit_offset"`
	MbcFlag                       string `json:"mbc_flag"`
	CutoffPrice                   string `json:"cutoff_price"`
	ValidityDate                  string `json:"validity_date"`
}

func (bc *BreezeClient) GetOrderDetails(order OrderDetailsRequest) ([]OrderDetailsSuccess, error) {
	res, err := bc.request("GET", "order", order)
	if err != nil {
		return nil, fmt.Errorf("Error in request: %v", err)
	}
	// resBytes := &OrderDetailsResponse{}
	resBytes, ok := res.([]byte)
	if !ok {
		return nil, fmt.Errorf("Error, the response from request() is not a byte slice. res:%v", res)
	}
	resBody := &OrderDetailsResponse{}
	err = json.Unmarshal(resBytes, resBody)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshalling the response body: %v", err)
	}
	if resBody.Error != nil {
		return nil, fmt.Errorf("Error in Response Body: %v", err)
	}
	return resBody.Success, nil
}

// Order List
type OrderListRequest struct {
	ExchangeCode string `json:"exchange_code"`
	FromDate     string `json:"from_date"`
	ToDate       string `json:"to_date"`
}

type OrderListSuccess struct {
	OrderId                       string `json:"order_id"`
	ExchangeOrderId               string `json:"exchange_order_id"`
	ExchangeCode                  string `json:"exchange_code"`
	StockCode                     string `json:"stock_code"`
	ProdcutType                   string `json:"product_type"`
	Action                        string `json:"action"`
	OrderType                     string `json:"order_type"`
	Stoploss                      string `json:"stoploss"`
	Quantity                      string `json:"quantity"`
	Price                         string `json:"price"`
	Validity                      string `json:"validity"`
	DisclosedQuantity             string `json:"disclosed_quantity"`
	ExpiryDate                    string `json:"expiry_date"`
	Right                         string `json:"right"`
	StrikePrice                   string `json:"strike_price"`
	AveragePrice                  string `json:"average_price"`
	CancelledQuantity             string `json:"cancelled_quantity"`
	PendingQuantity               string `json:"pending_quantity"`
	Status                        string `json:"status"`
	UserRemark                    string `json:"user_remark"`
	OrderDatetime                 string `json:"order_datetime"`
	ParentOrderId                 string `json:"parent_order_id"`
	ModiificationNumber           string `json:"modification_number"`
	ExchangeAcknowledgementDate   string `json:"exchange_acknowledgement_date"`
	ExchangeAcknowledgementNumber string `json:"exchange_acknowledge_number"`
	SLIPPrice                     string `json:"SLTP_price"`
	InitialLimit                  string `json:"initial_limit"`
	InitialSLTP                   string `json:"intial_sltp"`
	LTP                           string `json:"LTP"`
	LimitOffset                   string `json:"limit_offset"`
	MbcFlag                       string `json:"mbc_flag"`
	CutoffPrice                   string `json:"cutoff_price"`
	ValidityDate                  string `json:"validity_date"`
}
type OrderListResponse struct {
	Success []OrderListSuccess
	Status  string
	Error   any
}

func (bc *BreezeClient) GetOrderList(order OrderListRequest) ([]OrderListSuccess, error) {
	res, err := bc.request("GET", "order", order)
	if err != nil {
		return nil, fmt.Errorf("Error in request: %v", err)
	}
	// resBytes := &OrderDetailsResponse{}
	resBytes, ok := res.([]byte)
	if !ok {
		return nil, fmt.Errorf("Error, the response from request() is not a byte slice. res:%v", res)
	}
	resBody := &OrderListResponse{}
	err = json.Unmarshal(resBytes, resBody)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshalling the response body: %v", err)
	}
	if resBody.Error != nil {
		return nil, fmt.Errorf("Error in Response Body: %v", err)
	}
	return resBody.Success, nil
}
