package vonage_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/tenkoh/vonage-go-sdk"
)

func TestVerify_Verify(t *testing.T) {
	client, err := vonage.NewClient(
		vonage.ApiKey(apiKey),
		vonage.ApiSecret(apiSecret),
	)
	if err != nil {
		t.Error(err)
		return
	}

	resp, err := client.GenerateVerifyClient().Verify(
		vonage.VerifyNumber(recipientNumber),
		vonage.VerifyBrand(brandName),
	)
	if err != nil {
		t.Error(err)
		return
	}
	log.Printf("%+v\n", resp)
	status, err := resp.GetStatus()
	if err != nil {
		t.Error(err)
		return
	}
	if status != vonage.VerifyStatusOK {
		t.Error("bad response status")
	}
	fmt.Printf("request id: %s\n", resp.GetRequestID())
}
