package types

import (
	"regexp"
)

var nameSlugRegex = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9-]{3,48}[a-zA-Z0-9]$`)

type Name struct {
	value string
}

func NewName(s string) (Name, error) {
	if !nameSlugRegex.MatchString(s) {
		return Name{}, ErrInvalidName
	}

	return Name{
		value: s,
	}, nil
}

func (n Name) Value() string {
	return n.value
}
