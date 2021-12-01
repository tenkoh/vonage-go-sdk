package vonage

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
)

const (
	SDK_VERSION     = "0.0.1"
	API_HOST        = "https://api.nexmo.com"
	DEFAULT_TIMEOUT = 30
)

type VonageClient struct {
	apiKey    string
	apiSecret string
	userAgent string
	apiHost   string
	client    *http.Client
}

type VonageRequest struct {
	ApiKey    string `json:"api_key"`
	ApiSecret string `json:"api_secret"`
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
		client.setEnvApiKey()
	}
	if client.apiSecret == "" {
		client.setEnvApiSecret()
	}
	if err := validateAuthParameters(client.apiKey, client.apiSecret); err != nil {
		return nil, fmt.Errorf("fail to create a new client; %w", err)
	}
	client.setUserAgent()
	client.apiHost = API_HOST
	client.client = new(http.Client)
	return client, nil
}

func (vc *VonageClient) GenerateVerifyClient() *VerifyClient {
	verify := new(VerifyClient)
	verify.client = vc
	return verify
}

func (vc *VonageClient) authBody(body interface{}) io.Reader {
	auth := &VonageRequest{vc.apiKey, vc.apiSecret}
	merged := mergeStructsAsJson(auth, body)
	return merged
}

func (vc *VonageClient) MakeAuthRequest(method, host, endpoint string, body interface{}) (*http.Request, error) {
	uri, err := uriJoin(host, endpoint)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(method, uri, vc.authBody(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", vc.userAgent)
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(vc.apiKey, vc.apiSecret)
	return req, nil
}

func (vc *VonageClient) GetApiKey() string {
	return vc.apiKey
}

func (vc *VonageClient) setApiKey(key string) {
	vc.apiKey = key
}

func (vc *VonageClient) setEnvApiKey() {
	key := os.Getenv("VONAGE_API_KEY")
	if key == "" {
		return
	}
	vc.apiKey = key
}

func (vc *VonageClient) GetApiSecret() string {
	return vc.apiSecret
}

func (vc *VonageClient) setApiSecret(secret string) {
	vc.apiSecret = secret
}

func (vc *VonageClient) setEnvApiSecret() {
	secret := os.Getenv("VONAGE_API_SECRET")
	if secret == "" {
		return
	}
	vc.apiSecret = secret
}

func (vc *VonageClient) setUserAgent() {
	ua := fmt.Sprintf("vonage-go/%s go/%s", SDK_VERSION, runtime.Version())
	vc.userAgent = ua
}

func (vc *VonageClient) GetUserAgent() string {
	return vc.userAgent
}

// Constructor methods
func ApiKey(key string) Option {
	return func(vc *VonageClient) {
		vc.setApiKey(key)
	}
}

func ApiSecret(secret string) Option {
	return func(vc *VonageClient) {
		vc.setApiSecret(secret)
	}
}

func validateAuthParameters(key, secret string) error {
	if key == "" || secret == "" {
		return ErrInvalidAuthParameters
	}
	return nil
}
