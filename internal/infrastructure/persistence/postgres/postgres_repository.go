package postgres

import (
	"errors"
	"time"
	"tunelport-svc-devices/internal/domain/models"
	"tunelport-svc-devices/internal/domain/persistence"
	"tunelport-svc-devices/internal/domain/types"

	"gorm.io/gorm"
)

type PostgresDeviceRepository struct {
	db *gorm.DB
}

type deviceRecord struct {
	ID        string    `gorm:"primaryKey;type:uuid;not null"`
	Name      string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}

func (deviceRecord) TableName() string {
	return "devices"
}

func NewPostgresDeviceRepository(db *gorm.DB) persistence.DeviceRepository {
	if db == nil {
		panic("postgres repository requires a valid gorm db instance")
	}

	if err := db.AutoMigrate(&deviceRecord{}); err != nil {
		panic(err)
	}

	return &PostgresDeviceRepository{db: db}
}

func (r *PostgresDeviceRepository) Save(device *models.Device) error {
	if device == nil {
		return errors.New("device is nil")
	}
	if r.db == nil {
		return errors.New("database connection is nil")
	}
	if device.ID.Value() == "" ||
		device.Name.Value() == "" ||
		device.CreatedAt.IsZero() ||
		device.UpdatedAt.IsZero() {
		return errors.New("device is incomplete")
	}

	payload := deviceRecord{
		ID:        device.ID.Value(),
		Name:      device.Name.Value(),
		CreatedAt: device.CreatedAt,
		UpdatedAt: device.UpdatedAt,
	}

	return r.db.Save(&payload).Error
}

func (r *PostgresDeviceRepository) Get(id *types.ID) (*models.Device, error) {
	if id == nil {
		return nil, errors.New("device id is nil")
	}
	if r.db == nil {
		return nil, errors.New("database connection is nil")
	}

	var payload deviceRecord
	result := r.db.Where("id = ?", id.Value()).First(&payload)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, result.Error
		}
		return nil, result.Error
	}

	deviceID, err := types.NewIDFromString(payload.ID)
	if err != nil {
		return nil, err
	}

	deviceName, err := types.NewName(payload.Name)
	if err != nil {
		return nil, err
	}

	return &models.Device{
		ID:        deviceID,
		Name:      deviceName,
		CreatedAt: payload.CreatedAt,
		UpdatedAt: payload.UpdatedAt,
	}, nil
}

func (r *PostgresDeviceRepository) Delete(id *types.ID) error {
	if id == nil {
		return errors.New("device id is nil")
	}
	if r.db == nil {
		return errors.New("database connection is nil")
	}

	result := r.db.Delete(&deviceRecord{}, "id = ?", id.Value())
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *PostgresDeviceRepository) Update(device *models.Device) error {
	if device == nil {
		return errors.New("device is nil")
	}
	if r.db == nil {
		return errors.New("database connection is nil")
	}
	if device.ID.Value() == "" ||
		device.Name.Value() == "" ||
		device.CreatedAt.IsZero() ||
		device.UpdatedAt.IsZero() {
		return errors.New("device is incomplete")
	}

	payload := deviceRecord{
		ID:        device.ID.Value(),
		Name:      device.Name.Value(),
		CreatedAt: device.CreatedAt,
		UpdatedAt: device.UpdatedAt,
	}

	result := r.db.Model(&deviceRecord{}).Where("id = ?", payload.ID).Updates(payload)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *PostgresDeviceRepository) List() ([]*models.Device, error) {
	if r.db == nil {
		return nil, errors.New("database connection is nil")
	}

	var payloads []deviceRecord
	if err := r.db.Find(&payloads).Error; err != nil {
		return nil, err
	}

	devices := make([]*models.Device, 0, len(payloads))
	for _, payload := range payloads {
		deviceID, err := types.NewIDFromString(payload.ID)
		if err != nil {
			return nil, err
		}

		deviceName, err := types.NewName(payload.Name)
		if err != nil {
			return nil, err
		}

		devices = append(devices, &models.Device{
			ID:        deviceID,
			Name:      deviceName,
			CreatedAt: payload.CreatedAt,
			UpdatedAt: payload.UpdatedAt,
		})
	}

	return devices, nil
}
