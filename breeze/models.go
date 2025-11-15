package breeze

const baseURL string = "https://api.icicidirect.com/breezeapi/api/v1/"

type CustomerDetails struct {
	AppKey       string `json:"AppKey"`
	SessionToken string `json:"SessionToken"`
}
type SuccessResponse struct {
	Session_token string `json:"session_token"`
}
type CustomerDetailsResponse struct {
	Success SuccessResponse `json:"Success"`
	Status  int             `json:"Status"`
	Error   any             `json:"Error"`
}

// PUBLIC API STRUCTS

// get holdings
type GetHoldingsRequest struct{}
type DematHolding struct {
	StockCode               string `json:"stock_code"`
	StockISIN               string `json:"stock_ISIN"`
	Quantity                string `json:"quantity"`
	DemantTotalBulkQuantity string `json:"demat_total_bulk_quantity"`
	DemantAvailQuantity     string `json:"demat_avail_quantity"`
	BlockedQuantity         string `json:"blocked_quantity"`
	DematAllocatedQuantity  string `json:"demat_allocated_quantity"`
}
type GetHoldingsResponse struct {
	Success []DematHolding
	Status  int
	Error   any
}

// Place Order
// stock_code 	String 	"AXIBAN", "TATMOT" 	Yes
// exchange_code 	String 	"NSE", "NFO" 	Yes
// product 	String 	"futures","options","optionplus","cash","btst","margin" 	Yes
// action 	String 	"buy", "sell" 	Yes
// order_type 	String 	"limit","market","stoploss" 	Yes
// stoploss 	Double 	Numeric Currency 	No
// quantity 	String 	Number of quantity to place the order 	Yes
// price 	String 	Numeric Currency 	Yes
// validity 	String 	"day","ioc" 	No
// validity_date 	String 	ISO 8601 	No
// disclosed_quantity 	String 	Number of quantity to be disclosed 	No
// expiry_date 	String 	ISO 8601 	Yes
// right 	String 	"call","put","others" 	Yes
// strike_price 	String 	Numeric Currency 	Yes
// user_remark 	String 	Users are free to add their comment/tag to the order 	No
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
