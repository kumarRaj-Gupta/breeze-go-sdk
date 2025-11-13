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

func (bc *BreezeClient) request(method, path string, payload any) (any, error) {
	// We won't simply json.Marshl() our payload because we have to make sure that extra whitespaces
	// dont' seep into our request body. That will make Checksum fail on the API side in complex scenarios.
	// Instead, we ought to create a buffer, write our payload into it and encode it into a JSON using json.encode()
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf) // This will write the JSON Encoding directly to our buffer buf
	// encoder.SetIndent("", "") // Disabled for now.
	err := encoder.Encode(payload)
	if err != nil {
		return nil, fmt.Errorf("Error in encoding our payload: %v", err)
	}

	payloadString := buf.Bytes()
	// Generating Timestamp and Checksum
	timestamp := bc.generateTimestamp()
	checksum := bc.generateChecksum(timestamp, string(payloadString))
	// Full Path
	fullPath := baseURL + path
	// Creating the Request
	req, err := http.NewRequest(method, fullPath, &buf)
	if err != nil {
		return nil, fmt.Errorf("Error creating request :%v", err)
	}
	req.Header.Add("X-Checksum", "token"+checksum)
	req.Header.Add("X-Timestamp", timestamp)
	req.Header.Add("X-AppKey", bc.appKey)
	req.Header.Add("X-SessionToken", bc.sessionToken)
	// Creating the HTTP Client and sending the request
	httpClient := &http.Client{}
	res, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error occured in Response: %v", err)
	}
	responseBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("Error in reading the response: %v", err)
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("The response status is not 200 OK. Error Code:%v, Body: %v", res.StatusCode, string(responseBytes))
	}
	defer res.Body.Close()
	return responseBytes, nil
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
	baseLoginURL := "https://api.icicidirect.com/apiuser/login?api_key="
	baseLoginURL += bc.appKey
	return baseLoginURL
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
	fmt.Println("Session Token:", bc.sessionToken)

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
