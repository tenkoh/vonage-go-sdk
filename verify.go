package vonage

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"strconv"
)

const (
	// status
	VerifyStatusOK = 0
)

var verifyEndpoints = map[string]string{
	"verify": "/verify/json",
}

type VerifyClient struct {
	client *VonageClient
}

type VerifyRequest struct {
	ApiKey    string `json:"api_key"`
	ApiSecret string `json:"api_secret"`
	Number    string `json:"number,omitempty"`
	Brand     string `json:"brand,omitempty"`
}

type VerifyResponse struct {
	RequestID string `json:"request_id"`
	Status    string `json:"status"`
}

type VerifyOption func(*VerifyRequest)

func (vc *VerifyClient) Verify(options ...VerifyOption) (*VerifyResponse, error) {
	client := vc.client
	vreq := new(VerifyRequest)
	for _, option := range options {
		option(vreq)
	}
	// validate. Add methods when options are added.
	if vreq.Number == "" || vreq.Brand == "" {
		return nil, ErrInvalidVerifyParameters
	}
	// temp
	vreq.ApiKey = client.apiKey
	vreq.ApiSecret = client.apiSecret

	b, err := json.Marshal(vreq)
	if err != nil {
		return nil, err
	}
	req, err := client.MakeAuthRequest(
		"POST",
		client.apiHost,
		verifyEndpoints["verify"],
		bytes.NewBuffer(b),
	)
	if err != nil {
		return nil, err
	}
	resp, err := client.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// using TeeReader for debug. Reference: https://mattn.kaoriya.net/software/lang/go/20171026101727.htm
	var r io.Reader = resp.Body
	r = io.TeeReader(r, os.Stderr)

	var vres VerifyResponse
	if err := json.NewDecoder(r).Decode(&vres); err != nil {
		return nil, err
	}
	return &vres, nil
}

func VerifyNumber(number string) VerifyOption {
	return func(vr *VerifyRequest) {
		vr.Number = number
	}
}

func VerifyBrand(brand string) VerifyOption {
	return func(vr *VerifyRequest) {
		vr.Brand = brand
	}
}

func (vr *VerifyResponse) GetStatus() (int, error) {
	i, err := strconv.Atoi(vr.Status)
	if err != nil {
		return -1, err
	}
	return i, nil
}

func (vr *VerifyResponse) GetRequestID() string {
	return vr.RequestID
}
