package services_test

import (
	"testing"
	"tunelport-svc-devices/internal/domain/models"
	"tunelport-svc-devices/internal/domain/services"
	"tunelport-svc-devices/internal/domain/test"
	"tunelport-svc-devices/internal/domain/types"

	"github.com/stretchr/testify/assert"
)

func TestItShouldResultAResource(t *testing.T) {
	id := "43b14e10-e18a-4f88-9bcf-cd83698e5ba7"
	repository := test.NewStubRepository()

	sut := services.NewGetDeviceUseCase(repository)
	resource, err := sut.Do(id)

	assert.Nil(t, err)
	assert.IsType(t, &models.Device{}, resource)
	assert.True(t, repository.GetHaveBeenCalled(), sut)
}

func TestGetItShouldAnErrorOnInvalidId(t *testing.T) {
	id := "43b14e10-e18a-4f88-9bcf-cd83698e5ba"
	repository := test.NewStubRepository()

	sut := services.NewGetDeviceUseCase(repository)
	_, err := sut.Do(id)

	assert.IsType(t, err, types.ErrInvalidID)
}
