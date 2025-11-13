package breeze

import (
	"encoding/json"
	"fmt"
)

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
