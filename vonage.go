package vonage

const HOST_PATTERN = `^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$|^(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]*[a-zA-Z0-9])\.)+([A-Za-z]|[A-Za-z][A-Za-z0-9\-]*[A-Za-z0-9])$`

var version = "0.1.0"

type Client struct {
	key              string
	secret           string
	signature_secret string
	signature_method string
	application_id   string
	private_key      string
	app_name         string
	app_version      string
}

type Option func(*Client)

func Key(k string) Option {
	return func(c *Client) {
		c.key = k
	}
}

func Secret(s string) Option {
	return func(c *Client) {
		c.secret = s
	}
}

func SignatureSecret(s string) Option {
	return func(c *Client) {
		c.signature_secret = s
	}
}

func SignatureMethod(s string) Option {
	return func(c *Client) {
		c.signature_method = s
	}
}

func ApplicationID(id string) Option {
	return func(c *Client) {
		c.application_id = id
	}
}

func PrivateKey(pk string) Option {
	return func(c *Client) {
		c.private_key = pk
	}
}

func AppName(name string) Option {
	return func(c *Client) {
		c.app_name = name
	}
}

func AppVersion(ver string) Option {
	return func(c *Client) {
		c.app_version = ver
	}
}

func NewClient(options ...Option) (*Client, error) {
	c := new(Client)
	for _, option := range options {
		option(c)
	}
	// if nil option passed, set a ENV value.
	// setEnvValues(c)

	return c, nil
}
