package types

import (
	"encoding/base64"
	"strings"
)

const wireGuardPublicKeyBase64Len = 44

type PublicKey struct {
	value string
}

func NewPublicKey(s string) (PublicKey, error) {
	clean := strings.TrimSpace(s)
	if clean == "" {
		return PublicKey{}, ErrInvalidPublicKey
	}

	if len(clean) != wireGuardPublicKeyBase64Len {
		return PublicKey{}, ErrInvalidPublicKey
	}

	decoded, err := base64.StdEncoding.DecodeString(clean)
	if err != nil {
		return PublicKey{}, ErrInvalidPublicKey
	}

	if len(decoded) != 32 {
		return PublicKey{}, ErrInvalidPublicKey
	}

	return PublicKey{
		value: clean,
	}, nil
}

func (pk PublicKey) Value() string {
	return pk.value
}
