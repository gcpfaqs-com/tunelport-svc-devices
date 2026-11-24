package services_test

import (
	"time"

	"testing"
	"tunelport-svc-devices/internal/domain/services"
	"tunelport-svc-devices/internal/domain/test"
	"tunelport-svc-devices/internal/domain/types"

	"github.com/stretchr/testify/assert"
)

var newName string = "any-valid-name"

func TestItShouldSaveAResource(t *testing.T) {
	repository := test.NewStubRepository()

	sut := services.NewSaveDeviceUseCase(repository)
	_, err := sut.Do(newName, time.Now(), time.Now())

	assert.Nil(t, err)
	assert.True(t, repository.SaveHaveBeenCalled(), sut)
}

func TestItShouldResultAnErrorOnInvalidName(t *testing.T) {
	newName = ""
	repository := test.NewStubRepository()

	sut := services.NewSaveDeviceUseCase(repository)
	_, err := sut.Do(newName, time.Now(), time.Now())

	assert.NotNil(t, err)
	assert.IsType(t, types.ErrInvalidName, err)
}
