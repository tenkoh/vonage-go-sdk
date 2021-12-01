package vonage

import (
	"fmt"
	"io"
	"net/http"
)

type BasicAuthenticatedServer struct {
	host      string
	client    *http.Client
	timeout   int
	apiKey    string
	apiSecret string
	userAgent string
}

// In python implementation, this class contains session to reuse basic-authenticated request.
// However, Go does not support session natively.
// So, this Go class has a field; http.Client, and has a method; AuthRequest instead.
func NewBasicAuthenticatedServer(host, agent, key, secret string) *BasicAuthenticatedServer {
	server := new(BasicAuthenticatedServer)
	server.host = host
	server.client = &http.Client{}
	server.apiKey = key
	server.apiSecret = secret
	server.userAgent = agent
	return server
}

func (server *BasicAuthenticatedServer) AuthRequest(method, uri string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, uri, body)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(server.apiKey, server.apiSecret)
	req.Header.Add("User-Agent", server.userAgent)
	resp, err := server.client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (server *BasicAuthenticatedServer) Uri(path string) string {
	return fmt.Sprintf("%s%s", server.host, path)
}

func (server *BasicAuthenticatedServer) Get()

type ApplicationV2 struct {
	apiServer *BasicAuthenticatedServer
}

func (a *ApplicationV2) Create()
