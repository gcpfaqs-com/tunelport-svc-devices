package types

import "errors"

var (
	ErrDeviceNotFound      = errors.New("device not found")
	ErrDeviceAlreadyExists = errors.New("device already exists")
	ErrInvalidStatus       = errors.New("invalid device status")
	ErrInvalidID           = errors.New("invalid id must be an UUID v4 string")
	ErrInvalidName         = errors.New("must be 5-50 chars, start with a letter, end with letter or digit, and contain only letters, digits, or '-'")
	ErrInvalidIp           = errors.New("invalid ip address. must be an ipv4 address")
	ErrInvalidPublicKey    = errors.New("invalid public key")
)
