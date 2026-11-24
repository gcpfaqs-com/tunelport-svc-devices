package test

import (
	"tunelport-svc-devices/internal/domain/models"
	"tunelport-svc-devices/internal/domain/types"
)

type StubRepository struct {
	getCalled    bool
	deleteCalled bool
	listCalled   bool
	saveCalled   bool
	updateCalled bool
}

func NewStubRepository() *StubRepository {
	return &StubRepository{
		getCalled:    false,
		deleteCalled: false,
		listCalled:   false,
		saveCalled:   false,
		updateCalled: false,
	}
}

func (r *StubRepository) GetHaveBeenCalled() bool {
	result := r.getCalled

	r.getCalled = false

	return result
}

func (r *StubRepository) Get(id *types.ID) (*models.Device, error) {
	r.getCalled = true

	return &models.Device{}, nil
}

func (r *StubRepository) SaveHaveBeenCalled() bool {
	result := r.saveCalled

	r.saveCalled = false

	return result
}

func (r *StubRepository) Save(device *models.Device) error {
	r.saveCalled = true

	return nil
}

func (r *StubRepository) DeleteHaveBeenCalled() bool {
	result := r.deleteCalled

	r.deleteCalled = false

	return result
}

func (r *StubRepository) Delete(id *types.ID) error {
	r.deleteCalled = true

	return nil
}

func (r *StubRepository) UpdateHaveBeenCalled() bool {
	result := r.updateCalled

	r.updateCalled = false

	return result
}

func (r *StubRepository) Update(device *models.Device) error {
	r.updateCalled = true

	return nil
}

func (r *StubRepository) ListHaveBeenCalled() bool {
	result := r.listCalled

	r.listCalled = false

	return result
}

func (r *StubRepository) List() ([]*models.Device, error) {
	r.listCalled = true

	return []*models.Device{}, nil
}
