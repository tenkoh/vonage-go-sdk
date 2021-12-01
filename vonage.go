package vonage

import (
	"fmt"
	"os"
)

const (
	SDK_VERSION       = "0.1.0"
	AUTH_EXP_DURATION = 60
	HOST              = "rest.nexmo.com"
	API_HOST          = "api.nexmo.com"
	HOST_PATTERN      = `^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$|^(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]*[a-zA-Z0-9])\.)+([A-Za-z]|[A-Za-z][A-Za-z0-9\-]*[A-Za-z0-9])$`
)

type VonageClient struct {
	apiKey    string
	apiSecret string
}

type Option func(*VonageClient)

// NewClient returns a *VonageClient.
// In basic usage, apiKey and apiSecret have to be passed as options.
// When apiKey and apiSecret are not passed,
// this constructor tries to get environment variables named as VONAGE_API_KEY and VONAGE_API_SECRET instead.
func NewClient(options ...Option) (*VonageClient, error) {
	client := new(VonageClient)
	for _, option := range options {
		option(client)
	}
	if client.apiKey == "" {
		client.SetEnvApiKey()
	}
	if client.apiSecret == "" {
		client.SetEnvApiSecret()
	}
	if err := validateAuthParameters(client.apiKey, client.apiSecret); err != nil {
		return nil, fmt.Errorf("fail to create a new client; %w", err)
	}
	return client, nil
}

func (vc *VonageClient) GenerateVerifyClient() *VerifyClient {
	return nil
}

func (vc *VonageClient) GetApiKey() string {
	return vc.apiKey
}

func (vc *VonageClient) SetApiKey(key string) {
	vc.apiKey = key
}

func (vc *VonageClient) SetEnvApiKey() {
	key := os.Getenv("VONAGE_API_KEY")
	if key == "" {
		return
	}
	vc.apiKey = key
}

func (vc *VonageClient) GetApiSecret() string {
	return vc.apiSecret
}

func (vc *VonageClient) SetApiSecret(secret string) {
	vc.apiSecret = secret
}

func (vc *VonageClient) SetEnvApiSecret() {
	secret := os.Getenv("VONAGE_API_SECRET")
	if secret == "" {
		return
	}
	vc.apiSecret = secret
}

func ApiKey(key string) Option {
	return func(vc *VonageClient) {
		vc.SetApiKey(key)
	}
}

func ApiSecret(secret string) Option {
	return func(vc *VonageClient) {
		vc.SetApiSecret(secret)
	}
}

func validateAuthParameters(key, secret string) error {
	if key == "" || secret == "" {
		return ErrInvalidAuthParameters
	}
	return nil
}
