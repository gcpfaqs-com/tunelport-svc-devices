package rest

import (
	"encoding/json"
	"net/http"
	"tunelport-svc-devices/internal/domain/persistence"
	"tunelport-svc-devices/internal/domain/services"
)

type RESTSaveUseCase struct {
	repository persistence.DeviceRepository
}

func NewRESTSaveUseCase(repository persistence.DeviceRepository) *RESTSaveUseCase {
	return &RESTSaveUseCase{
		repository: repository,
	}
}

func (uc *RESTSaveUseCase) Save(w http.ResponseWriter, r *http.Request) {
	var dto deviceDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	saveUC := services.NewSaveDeviceUseCase(uc.repository)
	id, err := saveUC.Do(
		dto.Name,
		dto.CreatedAt,
		dto.UpdatedAt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	dto.ID = id
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(dto)
}
