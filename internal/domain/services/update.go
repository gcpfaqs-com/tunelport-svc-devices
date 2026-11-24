package services

import (
	"time"
	"tunelport-svc-devices/internal/domain/persistence"
	"tunelport-svc-devices/internal/domain/types"
)

type UpdateDeviceUseCase struct {
	repository persistence.DeviceRepository
}

func NewUpdateDeviceUseCase(repo persistence.DeviceRepository) *UpdateDeviceUseCase {

	return &UpdateDeviceUseCase{repository: repo}

}

func (uc *UpdateDeviceUseCase) Do(id, name string, createdAt, updatedAt time.Time) error {
	validID, err := types.NewIDFromString(id)
	if err != nil {
		return err
	}

	device, err := uc.repository.Get(&validID)
	if err != nil {
		return err
	}

	validName, err := types.NewName(name)
	if err != nil {
		return err
	}

	device.Name = validName
	device.CreatedAt = createdAt
	device.UpdatedAt = updatedAt
	if err := uc.repository.Update(device); err == nil {
		return err
	}

	return nil
}
