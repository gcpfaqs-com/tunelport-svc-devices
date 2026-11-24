package rest

import (
	"encoding/json"
	"net/http"
	"tunelport-svc-devices/internal/domain/persistence"
	"tunelport-svc-devices/internal/domain/services"
)

type RESTListUseCase struct {
	repository persistence.DeviceRepository
}

func NewRESTListUseCase(repository persistence.DeviceRepository) *RESTListUseCase {
	return &RESTListUseCase{
		repository: repository,
	}
}

func (uc *RESTListUseCase) List(w http.ResponseWriter, r *http.Request) {
	listUC := services.NewListDeviceUseCase(uc.repository)
	devices, err := listUC.Do()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := make([]deviceDTO, 0, len(devices))
	for _, a := range devices {
		response = append(response, deviceDTO{
			ID:        a.ID.Value(),
			Name:      a.Name.Value(),
			CreatedAt: a.CreatedAt,
			UpdatedAt: a.UpdatedAt,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
