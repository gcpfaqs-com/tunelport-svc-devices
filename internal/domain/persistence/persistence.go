package persistence

import (
	"tunelport-svc-devices/internal/domain/models"
	"tunelport-svc-devices/internal/domain/types"
)

type DeviceRepository interface {
	Save(device *models.Device) error
	Get(id *types.ID) (*models.Device, error)
	Delete(id *types.ID) error
	Update(device *models.Device) error
	List() ([]*models.Device, error)
}
