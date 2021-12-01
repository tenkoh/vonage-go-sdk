package vonage_test

import (
	"fmt"
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

	resp, err := client.GenerateVerifyClient().Verify()
	if err != nil {
		t.Error(err)
		return
	}
	if resp.GetStatus() != vonage.VerifyStatusOK {
		t.Error("bad response status")
	}
	fmt.Printf("request id: %s\n", resp.GetRequestID())
}
