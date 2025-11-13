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
