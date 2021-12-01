package vonage

import (
	"crypto"
	"fmt"
	"hash"
	"io"
	"os"
	"regexp"
	"runtime"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

const (
	AUTH_EXP_DURATION = 60
	HOST              = "rest.nexmo.com"
	API_HOST          = "api.nexmo.com"
	HOST_PATTERN      = `^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$|^(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]*[a-zA-Z0-9])\.)+([A-Za-z]|[A-Za-z][A-Za-z0-9\-]*[A-Za-z0-9])$`
)

var version = "0.1.0"

type Client struct {
	apiKey              string
	apiSecret           string
	signatureSecret     string
	signatureMethodName string
	signatureMethod     hash.Hash
	applicationID       string
	privateKey          string // In Python SDK, this field accept string or byte.
	appName             string
	appVersion          string
	hostPattern         string
	host                string
	apiHost             string
	headers             map[string]string
	auth                map[string]interface{}
}

type Option func(*Client)

func Key(k string) Option {
	return func(c *Client) {
		c.apiKey = k
	}
}

func Secret(s string) Option {
	return func(c *Client) {
		c.apiSecret = s
	}
}

func SignatureSecret(s string) Option {
	return func(c *Client) {
		c.signatureSecret = s
	}
}

func SignatureMethod(s string) Option {
	return func(c *Client) {
		c.signatureMethodName = s
	}
}

func ApplicationID(id string) Option {
	return func(c *Client) {
		c.applicationID = id
	}
}

// [FIXME] []byte input is required?
func PrivateKey(pk interface{}) Option {
	switch pk := pk.(type) {
	case string:
		return func(c *Client) {
			c.privateKey = pk
		}
	case []byte:
		_pk := string(pk)
		return func(c *Client) {
			c.privateKey = _pk
		}
	default:
		return func(c *Client) {}
	}
}

func AppName(name string) Option {
	return func(c *Client) {
		c.appName = name
	}
}

func AppVersion(ver string) Option {
	return func(c *Client) {
		c.appVersion = ver
	}
}

func (c *Client) setEnvValues() {
	if c.apiKey == "" {
		c.apiKey = os.Getenv("VONAGE_API_KEY")
	}
	if c.apiSecret == "" {
		c.apiSecret = os.Getenv("VONAGE_API_SECRET")
	}
	if c.signatureSecret == "" {
		c.signatureSecret = os.Getenv("VONAGE_SIGNATURE_SECRET")
	}
	if c.signatureMethodName == "" {
		c.signatureMethodName = os.Getenv("VONAGE_SIGNATURE_METHOD")
	}
}

func (c *Client) setSignatureMethod() {
	switch c.signatureMethodName {
	case "md5":
		c.signatureMethod = crypto.MD5.New()
	case "sha1":
		c.signatureMethod = crypto.SHA1.New()
	case "sha256":
		c.signatureMethod = crypto.SHA256.New()
	case "sha512":
		c.signatureMethod = crypto.SHA512.New()
	}
}

func (c *Client) loadExternalPrivateKey() error {
	if c.privateKey == "" {
		return nil
	}
	if strings.Contains(c.privateKey, "\n") {
		return ErrInvalidPrivateKey
	}
	f, err := os.Open(c.privateKey)
	if err != nil {
		return fmt.Errorf("fail to open private key; %w", err)
	}
	defer f.Close()
	b, err := io.ReadAll(f)
	if err != nil {
		return fmt.Errorf("fail to read private key; %w", err)
	}
	c.privateKey = string(b)
	return nil
}

func (c *Client) setConstants() error {
	c.hostPattern = HOST_PATTERN
	if err := c.SetHost(HOST); err != nil {
		return err
	}
	c.apiHost = API_HOST
	return nil
}

func (c *Client) setStringLiterals() {
	ua := fmt.Sprintf("vonage-go/%s go/%s", version, runtime.Version())
	if c.appName != "" && c.appVersion != "" {
		ua += fmt.Sprintf(" %s/%s", c.appName, c.appVersion)
	}
	c.headers["User-Agent"] = ua
}

// NewClient returns a pointer of vonage-client.
// Required options varies due to usage. Check README.
// An example:
// client, err := NewClient(Key("YOUR_KEY"), Secret("YOUR_SECRET"))
func NewClient(options ...Option) (*Client, error) {
	c := new(Client)
	for _, option := range options {
		option(c)
	}
	// prevent nil pointer
	c.headers = map[string]string{}
	c.auth = map[string]interface{}{}

	if err := c.loadExternalPrivateKey(); err != nil {
		return nil, err
	}
	if err := c.setConstants(); err != nil {
		return nil, err
	}
	c.setEnvValues()
	c.setSignatureMethod()
	c.setStringLiterals()
	return c, nil
}

func (c *Client) SetHost(host string) error {
	reg := regexp.MustCompile(c.hostPattern)
	if !reg.MatchString(host) {
		return ErrInvalidHostName
	}
	c.host = host
	return nil
}

func (c *Client) GetHost() string {
	return c.host
}

func (c *Client) SetApiHost(host string) error {
	reg := regexp.MustCompile(c.hostPattern)
	if !reg.MatchString(host) {
		return ErrInvalidHostName
	}
	c.apiHost = host
	return nil
}

func (c *Client) GetApiHost() string {
	return c.apiHost
}

func (c *Client) Auth(auth map[string]interface{}) {
	c.auth = auth
}

// [FIXME] this implementation is ambiguous
func (c *Client) generateApplicationJwt() (string, error) {
	claims := jwt.MapClaims{}
	claims["application_id"] = c.applicationID
	claims["iat"] = int(time.Now().Unix())
	claims["exp"] = int(time.Now().Unix() + AUTH_EXP_DURATION)
	claims["jti"] = func() string {
		u := new(uuid.UUID)
		return u.String()
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	// overwrite claim with user's configuration
	for k, v := range c.auth {
		claims[k] = v
	}
	tokenString, err := token.SignedString(c.privateKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
