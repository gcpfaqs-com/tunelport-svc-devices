package services

import (
	"tunelport-svc-devices/internal/domain/models"
	"tunelport-svc-devices/internal/domain/persistence"
	"tunelport-svc-devices/internal/domain/types"
)

type GetDeviceUseCase struct {
	repository persistence.DeviceRepository
}

func NewGetDeviceUseCase(repo persistence.DeviceRepository) *GetDeviceUseCase {
	return &GetDeviceUseCase{repository: repo}
}

func (uc *GetDeviceUseCase) Do(id string) (*models.Device, error) {
	validID, err := types.NewIDFromString(id)
	if err != nil {
		return nil, err
	}

	device, err := uc.repository.Get(&validID)
	if err != nil {
		return nil, err
	}

	return device, nil
}
