package breeze

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type BreezeClient struct {
	appKey       string
	secretKey    string
	sessionToken string
}

// Constructor for BreezeClient
func NewBreezeClient(AppKey, SecretKey string) *BreezeClient {
	return &BreezeClient{
		appKey:    AppKey,
		secretKey: SecretKey,
	}
}

// returns the current UTC time formatted exactly as required by the API
// : ISO8601 with 0 milliseconds, ending in Z (e.g., 2024-06-01T10:23:56.000Z).
func (bc *BreezeClient) generateTimestamp() string {
	currentTime := time.Now().UTC()
	formattedTime := currentTime.Format("2006-01-02T15:04:05.000Z")
	return formattedTime
}

// takes the generated timestamp and the JSON payload string,
// combines them with the SecretKey, and returns the SHA256 hexadecimal hash of the result
func (bc *BreezeClient) generateChecksum(timestamp string, payload string) string {
	combinedString := timestamp + payload + bc.secretKey
	hasher := sha256.New()
	hasher.Write([]byte(combinedString))
	return hex.EncodeToString(hasher.Sum(nil))

}

// login url
func (bc *BreezeClient) GetLoginURL() string {
	baseURL := "https://api.icicidirect.com/apiuser/login?api_key="
	baseURL += bc.appKey
	return baseURL
}

func (bc *BreezeClient) ObtainSessionToken(apiSessionKey string) error {
	url := "https://api.icicidirect.com/breezeapi/api/v1/customerdetails"

	CustomerDetails := CustomerDetails{
		AppKey:       bc.appKey,
		SessionToken: apiSessionKey,
	}

	payload, err := json.Marshal(CustomerDetails)
	if err != nil {
		return fmt.Errorf("Error Marshalling CustomerDetails struct in ObtainSessionToken:%v", err)
	}

	req, err := http.NewRequest("GET", url, bytes.NewBuffer(payload))
	if err != nil {
		return fmt.Errorf("Error creating new request body in ObtainSessionToken:%v", err)
	}

	// Attaching Headers
	req.Header.Add("Content-Type", "application/json")

	// Make the Request
	// 1. Create a http client
	// 2. Make the request do()
	httpClient := &http.Client{}
	res, err := httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("Error getting response in ObtainSessionToken:%v", err)
	}
	defer res.Body.Close()
	// Checking if the error code is not 200 OK
	if res.StatusCode != 200 {
		return fmt.Errorf("Unsuccessful call to API. StatusCode:%v", res.StatusCode)
	}
	cdr := CustomerDetailsResponse{}
	responseBody, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("Error reading the response body in ObtainSessionToken:%v", err)
	}
	err = json.Unmarshal(responseBody, &cdr)
	if err != nil {
		return fmt.Errorf("Error Unmarshalling the response boyd in ObtainSessionToken:%v", err)
	}

	// Checking if the response has not null error
	if cdr.Error != nil {
		return fmt.Errorf("Response has error. Error:%v", cdr.Error)
	}
	// saving the sessionToken
	bc.sessionToken = cdr.Success.Session_token

	return nil
}

// completing the login
func (bc *BreezeClient) CompleteLogin(redirectURL string) error {
	u, err := url.Parse(redirectURL)
	if err != nil {
		return fmt.Errorf("Error parsing the redirect URL:%v", err)
	}
	queryParams := u.Query()
	session_key := queryParams.Get("session_key")
	if session_key == "" {
		return fmt.Errorf("Session Key returned was empty")
	}
	err = bc.ObtainSessionToken(session_key)
	if err != nil {
		return fmt.Errorf("Error in obtaining Session Token in CompleteLogin %v", err)
	}
	return nil
}
