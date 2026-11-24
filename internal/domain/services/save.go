package services

import (
	"time"
	"tunelport-svc-devices/internal/domain/models"
	"tunelport-svc-devices/internal/domain/persistence"
)

type SaveDeviceUseCase struct {
	repository persistence.DeviceRepository
}

func NewSaveDeviceUseCase(repo persistence.DeviceRepository) *SaveDeviceUseCase {

	return &SaveDeviceUseCase{repository: repo}

}

func (uc *SaveDeviceUseCase) Do(name string, created_at, updated_at time.Time) (string, error) {
	device, err := models.NewDevice(name, created_at, updated_at)
	if err != nil {
		return "", err
	}

	err = uc.repository.Save(device)
	if err != nil {
		return "", err
	}

	return device.ID.Value(), nil
}
