package vonage_test

import (
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
	apiKey = os.Getenv("VONAGE_API_KEY")
	apiSecret = os.Getenv("VONAGE_API_SECRET")
	recipientNumber = os.Getenv("RECIPIENT_NUMBER")
	brandName = os.Getenv("BRAND_NUMBER")

	m.Run()
}
