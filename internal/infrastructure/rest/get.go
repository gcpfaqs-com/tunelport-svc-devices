package rest

import (
	"encoding/json"
	"net/http"
	"tunelport-svc-devices/internal/domain/persistence"
	"tunelport-svc-devices/internal/domain/services"
)

type RESTGetUseCase struct {
	repository persistence.DeviceRepository
}

func NewRESTGetUseCase(repository persistence.DeviceRepository) *RESTGetUseCase {
	return &RESTGetUseCase{
		repository: repository,
	}
}

func (uc *RESTGetUseCase) Get(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	getUC := services.NewGetDeviceUseCase(uc.repository)
	device, err := getUC.Do(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := deviceDTO{
		ID:        device.ID.Value(),
		Name:      device.Name.Value(),
		CreatedAt: device.CreatedAt,
		UpdatedAt: device.UpdatedAt,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
