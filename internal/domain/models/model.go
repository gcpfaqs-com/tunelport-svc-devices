package models

import (
	"time"
	"tunelport-svc-devices/internal/domain/types"
)

type Device struct {
	ID        types.ID        `json:"id"`
	Name      types.Name      `json:"name"`
	Ip        types.Ip        `json:"ip"`
	PublicKey types.PublicKey `json:"public_key"`
	Status    types.Status    `json:"status"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}

func NewDevice(
	name string,
	ip string,
	publicKey string,
	status string,
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

	devideIp, err := types.NewIp(ip)
	if err != nil {
		return nil, err
	}

	devicePublicKey, err := types.NewPublicKey(publicKey)
	if err != nil {
		return nil, err
	}

	deviceStatus, err := types.NewStatus(status)
	if err != nil {
		return nil, err
	}

	return &Device{
		ID:        deviceID,
		Name:      deviceName,
		Ip:        devideIp,
		PublicKey: devicePublicKey,
		Status:    deviceStatus,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}, nil
}
