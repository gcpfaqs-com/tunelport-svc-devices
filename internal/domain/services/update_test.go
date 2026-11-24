package services_test

import (
	"time"

	"testing"
	"tunelport-svc-devices/internal/domain/services"
	"tunelport-svc-devices/internal/domain/test"
	"tunelport-svc-devices/internal/domain/types"

	"github.com/stretchr/testify/assert"
)

var updateId string = "43b14e10-e18a-4f88-9bcf-cd83698e5ba7"
var updateName string = "any valid name"

func TestItShouldUpdateAResource(t *testing.T) {
	repository := test.NewStubRepository()

	sut := services.NewUpdateDeviceUseCase(repository)
	err := sut.Do(updateId, updateName, time.Now(), time.Now())

	assert.Nil(t, err)
	assert.True(t, repository.UpdateHaveBeenCalled(), sut)
}

func TestUpdateItShouldResultAnErrorOnInvalidName(t *testing.T) {
	name := ""
	repository := test.NewStubRepository()

	sut := services.NewUpdateDeviceUseCase(repository)
	err := sut.Do(updateId, name, time.Now(), time.Now())

	assert.NotNil(t, err)
	assert.IsType(t, types.ErrInvalidName, err)
}
