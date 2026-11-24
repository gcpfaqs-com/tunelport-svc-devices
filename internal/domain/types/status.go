package types

import (
	"fmt"
	"strings"
)

type Status struct {
	value string
}

const (
	StatusCreatedValue = "created"
	StatusActiveValue  = "active"
	StatusBlockedValue = "blocked"
)

var (
	StatusCreated = Status{value: StatusCreatedValue}
	StatusActive  = Status{value: StatusActiveValue}
	StatusBlocked = Status{value: StatusBlockedValue}
)

func NewStatus(s string) (Status, error) {
	normalized := strings.ToLower(strings.TrimSpace(s))

	if normalized == StatusCreatedValue {
		return StatusCreated, nil
	}
	if normalized == StatusActiveValue {
		return StatusActive, nil
	}
	if normalized == StatusBlockedValue {
		return StatusBlocked, nil
	}

	return Status{}, fmt.Errorf("%w: %q", ErrInvalidStatus, s)
}

func (st Status) Value() string {
	return st.value
}
