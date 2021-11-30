package vonage

import (
	"reflect"
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
			if !reflect.DeepEqual(got.key, tt.wantKey) {
				t.Errorf("NewClient() = %v, want %v", got.key, tt.wantKey)
			}
		})
	}
}
