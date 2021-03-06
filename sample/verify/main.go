package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/tenkoh/vonage-go-sdk"
)

var (
	recipientNumber string
	brandName       string
)

// load environment variables. must include below.
// VONAGE_API_KEY
// VONAGE_API_SECRET
// RECIPIENT_NUMBER
// BRAND_NAME
func init() {
	err := godotenv.Load(filepath.Clean("../../.env"))
	if err != nil {
		panic(err)
	}
	recipientNumber = os.Getenv("RECIPIENT_NUMBER")
	brandName = os.Getenv("BRAND_NAME")
}

func main() {
	// logger setting
	f, err := os.OpenFile("debug.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0755)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	log.SetOutput(f)

	// send verify code
	client, err := vonage.NewClient()
	if err != nil {
		panic(err)
	}
	verify := client.GenerateVerifyClient()
	resp, err := verify.Verify(
		vonage.VerifyNumber(recipientNumber),
		vonage.VerifyBrand(brandName),
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v\n", resp)
	reqID := resp.GetRequestID()

	// check verify coce
	fmt.Println("enter code. hit 'n' to cancel the process.")
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	code := sc.Text()

	if code == "n" {
		log.Println("start cancel process")
		resp, err := verify.Cancel(reqID)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("%+v\n", resp)
		return
	}

	log.Println("start check process")
	resp, err = verify.Check(
		vonage.VerifyCheckCode(code),
		vonage.VerifyCheckRequestID(reqID),
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v\n", resp)
}
