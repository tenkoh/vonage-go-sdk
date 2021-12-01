package vonage

import (
	"encoding/json"
	"strconv"
)

const (
	// status
	VerifyStatusOK = 0
)

var verifyEndpoints = map[string]string{
	"verify": "/verify/json",
	"check":  "/verify/check/json",
}

type VerifyClient struct {
	client *VonageClient
}

type VerifyRequest struct {
	Number string `json:"number"`
	Brand  string `json:"brand"`
}

type VerifyCheckRequest struct {
	RequestID string `json:"request_id"`
	Code      string `json:"code"`
}

type VerifyResponse struct {
	RequestID string `json:"request_id"`
	Status    string `json:"status"`
}

type VerifyOption func(*VerifyRequest)
type VerifyCheckOption func(*VerifyCheckRequest)

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
	req, err := client.MakeAuthRequest("POST", client.apiHost, verifyEndpoints["verify"], vreq)
	if err != nil {
		return nil, err
	}
	resp, err := client.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var vres VerifyResponse
	if err := json.NewDecoder(resp.Body).Decode(&vres); err != nil {
		return nil, err
	}
	return &vres, nil
}

func (vc *VerifyClient) Check(options ...VerifyCheckOption) (*VerifyResponse, error) {
	client := vc.client
	vreq := new(VerifyCheckRequest)
	for _, option := range options {
		option(vreq)
	}
	// validate. Add methods when options are added.
	if vreq.RequestID == "" || vreq.Code == "" {
		return nil, ErrInvalidVerifyParameters
	}
	req, err := client.MakeAuthRequest("POST", client.apiHost, verifyEndpoints["check"], vreq)
	if err != nil {
		return nil, err
	}
	resp, err := client.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var vres VerifyResponse
	if err := json.NewDecoder(resp.Body).Decode(&vres); err != nil {
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

func VerifyCheckRequestID(id string) VerifyCheckOption {
	return func(vcr *VerifyCheckRequest) {
		vcr.RequestID = id
	}
}

func VerifyCheckCode(code string) VerifyCheckOption {
	return func(vcr *VerifyCheckRequest) {
		vcr.Code = code
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
