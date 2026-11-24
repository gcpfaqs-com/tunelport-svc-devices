package services

import (
	"tunelport-svc-devices/internal/domain/persistence"
	"tunelport-svc-devices/internal/domain/types"
)

type DeleteDeviceUseCase struct {
	repository persistence.DeviceRepository
}

func NewDeleteDeviceUseCase(repo persistence.DeviceRepository) *DeleteDeviceUseCase {
	return &DeleteDeviceUseCase{repository: repo}
}

func (uc *DeleteDeviceUseCase) Do(id string) error {
	validID, err := types.NewIDFromString(id)
	if err != nil {
		return err
	}

	err = uc.repository.Delete(&validID)
	if err != nil {
		return err
	}

	return nil
}
