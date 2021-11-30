package vonage

import "errors"

var (
	ErrInvalidPrivateKey = errors.New("invalid private key")
	ErrInvalidHostName   = errors.New("invalid host name")
)
