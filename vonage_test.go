package vonage_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/tenkoh/vonage-go-sdk"
)

func TestNewVonageClient(t *testing.T) {
	// with input
	client, err := vonage.NewClient(
		vonage.ApiKey("foo"),
		vonage.ApiSecret("bar"),
	)
	if err != nil {
		t.Error(err)
		return
	}
	if got := client.GetApiKey(); got != "foo" {
		t.Errorf("want foo, got %s", got)
	}
	if got := client.GetApiSecret(); got != "bar" {
		t.Errorf("want bar, got %s", got)
	}

	// without input, use env
	client, err = vonage.NewClient()
	if err != nil {
		t.Error(err)
		return
	}
	if got := client.GetApiKey(); got != apiKey {
		t.Errorf("want %s, got %s", apiKey, got)
	}
	if got := client.GetApiSecret(); got != apiSecret {
		t.Errorf("want %s, got %s", apiSecret, got)
	}

}

func TestMakeRequest(t *testing.T) {
	client, err := vonage.NewClient()
	if err != nil {
		t.Error(err)
		return
	}
	req, err := client.MakeAuthRequest("GET", "foo", "bar", strings.NewReader("test"))
	if err != nil {
		t.Error(err)
		return
	}
	header := req.Header
	uas, ok := header["User-Agent"]
	if !ok {
		t.Error("no user-agent in request header")
		return
	}
	if len(uas) != 1 {
		t.Error("not expected user-agent in request header")
		return
	}
	ua := uas[0]
	if want := client.GetUserAgent(); ua != want {
		t.Errorf("want %s, got %s", want, ua)
	}
	fmt.Printf("%+v\n", req)
}
