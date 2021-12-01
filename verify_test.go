package vonage

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

var (
	apiKey          string
	apiSecret       string
	recipientNumber string
	brandName       string
)

func TestMain(m *testing.M) {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	apiKey = os.Getenv("API_KEY")
	apiSecret = os.Getenv("API_SECRET")
	recipientNumber = os.Getenv("RECIPIENT_NUMBER")
	brandName = os.Getenv("BRAND_NUMBER")

	m.Run()
}

func TestVerify_Verify(t *testing.T) {
	client, err := NewClient(
		ApiKey(apiKey),
		ApiSecret(apiSecret),
	)
	if err != nil {
		t.Error(err)
		return
	}

	resp, err := client.GenerateVerifyClient().Verify()
	if err != nil {
		t.Error(err)
		return
	}
	if resp.GetStatus() != VerifyStatusOK {
		t.Error("bad response status")
	}
	fmt.Printf("request id: %s\n", resp.GetRequestID())
}
