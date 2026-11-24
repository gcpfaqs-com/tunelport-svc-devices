package types

import (
	"net/netip"
	"strings"
)

type Ip struct {
	value string
}

func NewIp(s string) (Ip, error) {
	clean := strings.TrimSpace(s)

	addr, err := netip.ParseAddr(clean)
	if err != nil {
		return Ip{}, ErrInvalidIp
	}

	if !addr.Is4() {
		return Ip{}, ErrInvalidIp
	}

	return Ip{
		value: addr.String(),
	}, nil
}

func (ip Ip) Value() string {
	return ip.value
}
