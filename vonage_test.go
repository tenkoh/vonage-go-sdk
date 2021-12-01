package vonage_test

import (
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
