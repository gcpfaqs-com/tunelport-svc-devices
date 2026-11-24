package models

import (
	"time"
	"tunelport-svc-devices/internal/domain/types"
)

type Device struct {
	ID        types.ID   `json:"id"`
	Name      types.Name `json:"name"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

func NewDevice(
	name string,
	createdAt time.Time,
	updatedAt time.Time) (*Device, error) {
	deviceName, err := types.NewName(name)
	if err != nil {
		return nil, err
	}

	deviceID, err := types.NewID()
	if err != nil {
		return nil, err
	}

	return &Device{
		ID:        deviceID,
		Name:      deviceName,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}, nil
}
