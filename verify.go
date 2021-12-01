package vonage

const (
	VerifyStatusOK = 0
)

type VerifyClient struct {
}

type VerifyResponse struct {
}

func (vc *VerifyClient) Verify() (*VerifyResponse, error) {
	return nil, nil
}

func (vr *VerifyResponse) GetStatus() int {
	return 0
}

func (vr *VerifyResponse) GetRequestID() string {
	return ""
}
