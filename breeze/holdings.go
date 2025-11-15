package breeze

import (
	"encoding/json"
	"fmt"
)

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

func (bc *BreezeClient) GetHoldings() ([]DematHolding, error) {
	res, err := bc.request("GET", "dematholdings", GetHoldingsRequest{})
	if err != nil {
		return nil, fmt.Errorf("Error in getting response: %v", err)
	}
	resBody := &GetHoldingsResponse{}
	byteSlice, ok := res.([]byte)
	if ok {
		err = json.Unmarshal(byteSlice, resBody)
		if err != nil {
			return nil, fmt.Errorf("Response is not a byte slice")
		}
	}
	if resBody.Error != nil {
		return nil, fmt.Errorf("There was an error in receiving the response: Code:%v. Error Message:%v", resBody.Status, resBody)
	}
	return resBody.Success, nil

}
