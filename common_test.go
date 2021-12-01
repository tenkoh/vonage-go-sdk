package vonage_test

import (
	"io"
	"log"
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
	brandName = os.Getenv("BRAND_NAME")

	f, err := os.OpenFile("debug.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0755)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	multi := io.MultiWriter(f, os.Stdout)
	log.SetOutput(multi)

	m.Run()
}
