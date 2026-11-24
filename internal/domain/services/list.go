package services

import (
	"tunelport-svc-devices/internal/domain/models"
	"tunelport-svc-devices/internal/domain/persistence"
)

type ListDeviceUseCase struct {
	repository persistence.DeviceRepository
}

func NewListDeviceUseCase(repo persistence.DeviceRepository) *ListDeviceUseCase {
	return &ListDeviceUseCase{repository: repo}
}

func (uc *ListDeviceUseCase) Do() ([]*models.Device, error) {
	devices, err := uc.repository.List()
	if err != nil {
		return nil, err
	}

	return devices, nil
}
