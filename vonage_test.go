package vonage

import (
	"reflect"
	"strings"
	"testing"
)

func TestNewClient(t *testing.T) {
	type args struct {
		options []Option
	}
	op1 := Key("foo")
	tests := []struct {
		name    string
		args    args
		wantKey string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"no key", args{}, "", false},
		{"no key", args{[]Option{op1}}, "foo", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewClient(tt.args.options...)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.apiKey, tt.wantKey) {
				t.Errorf("NewClient() = %v, want %v", got.apiKey, tt.wantKey)
			}
		})
	}
}

func TestNewClientByEnv(t *testing.T) {
	envs := []string{"VONAGE_API_KEY", "VONAGE_API_SECRET", "VONAGE_SIGNATURE_SECRET", "VONAGE_SIGNATURE_METHOD"}
	for _, e := range envs {
		t.Setenv(e, strings.ToLower(e))
	}
	client, err := NewClient()
	if err != nil {
		t.Error(err)
		return
	}
	for i, e := range envs {
		var got string
		switch i {
		case 0:
			got = client.apiKey
		case 1:
			got = client.apiSecret
		case 2:
			got = client.signatureSecret
		case 3:
			got = client.signatureMethodName
		}
		want := strings.ToLower(e)
		if got != want {
			t.Errorf("want %s, got %s", want, got)
		}
	}
}

func TestNewClientPrivateKey(t *testing.T) {
	client, err := NewClient(PrivateKey("testdata/private.key"))
	if err != nil {
		t.Error(err)
	}
	if got := client.privateKey; got != "foobarkey" {
		t.Errorf("got %s", got)
	}
}
