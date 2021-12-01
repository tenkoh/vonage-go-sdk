package vonage

import "errors"

var (
	ErrInvalidAuthParameters   = errors.New("invalid auth parameters")
	ErrInvalidVerifyParameters = errors.New("invalid verify parameters")
)
