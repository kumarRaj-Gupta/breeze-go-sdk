package breeze

import (
	"testing"
	"time"
)

var TestClient *BreezeClient = NewBreezeClient("app_key", "secret_key")

// var sessionToken string = "session_token"

// This is not the right way to do this. I should instead be checking for pattern of the
// output format. The time may not be equal.
func TestGenerateTimestamp(t *testing.T) {
	expectedTime := time.Now().UTC()
	result := TestClient.generateTimestamp()
	expected := expectedTime.Format("2006-01-02T15:04:05.000Z")

	if result != expected {
		t.Errorf("Expected %v Got %v", expected, result)
	}
}

func TestGenerateChecksum(t *testing.T) {
	expected := "ef7155927e0fd51171a8b95c1b5f0c2d1fb7a6a3f2f23a13c8445a9833398b1a"
	result := TestClient.generateChecksum("timestamp", "{}")
	if result != expected {
		t.Errorf("Expected %v Got %v", expected, result)
	}

}
