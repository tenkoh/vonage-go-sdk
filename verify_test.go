package vonage_test

import (
	"fmt"
	"testing"

	"github.com/tenkoh/vonage-go-sdk"
)

// Note This test sends SMS, then wastes vonage credit.
// Invalidate here, if you are concerned about cost.

// func TestVerify_VerifyAndCheck(t *testing.T) {
// 	client, err := vonage.NewClient(
// 		vonage.ApiKey(apiKey),
// 		vonage.ApiSecret(apiSecret),
// 	)
// 	if err != nil {
// 		t.Error(err)
// 		return
// 	}
// 	resp, err := client.GenerateVerifyClient().Verify(
// 		vonage.VerifyNumber(recipientNumber),
// 		vonage.VerifyBrand(brandName),
// 	)
// 	if err != nil {
// 		t.Error(err)
// 		return
// 	}
// 	log.Printf("%+v\n", resp)
// 	status, err := resp.GetStatus()
// 	if err != nil {
// 		t.Error(err)
// 		return
// 	}
// 	if status != vonage.VerifyStatusOK {
// 		t.Error("bad response status")
// 	}
// }

func Test_Cancel(t *testing.T) {
	client, _ := vonage.NewClient()
	id := ""
	resp, _ := client.GenerateVerifyClient().Cancel(id)
	fmt.Println(resp)
}
