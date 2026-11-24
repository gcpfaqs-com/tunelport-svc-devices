package services_test

import (
	"testing"
	"tunelport-svc-devices/internal/domain/models"
	"tunelport-svc-devices/internal/domain/services"
	"tunelport-svc-devices/internal/domain/test"

	"github.com/stretchr/testify/assert"
)

func TestItShouldResultAList(t *testing.T) {
	repository := test.NewStubRepository()

	sut := services.NewListDeviceUseCase(repository)
	list, err := sut.Do()

	assert.Nil(t, err)
	assert.IsType(t, []*models.Device{}, list)
	assert.True(t, repository.ListHaveBeenCalled(), sut)
}
