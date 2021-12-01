package vonage

import "testing"

func Test_ValidateAuthParameters(t *testing.T) {
	type args struct {
		key    string
		secret string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{"both ok", args{"foo", "bar"}, false},
		{"empty key", args{"", "bar"}, true},
		{"empty secret", args{"foo", ""}, true},
		{"both empty", args{"", ""}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateAuthParameters(tt.args.key, tt.args.secret); (err != nil) != tt.wantErr {
				t.Errorf("validateAuthParameters() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
