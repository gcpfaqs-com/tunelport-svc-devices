package rest

import (
	"net/http"
	"tunelport-svc-devices/internal/domain/persistence"
	"tunelport-svc-devices/internal/domain/services"
)

type RESTDeleteUseCase struct {
	repository persistence.DeviceRepository
}

func NewRESTDeleteUseCase(repository persistence.DeviceRepository) *RESTDeleteUseCase {
	return &RESTDeleteUseCase{
		repository: repository,
	}
}

func (uc *RESTDeleteUseCase) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	deleteUC := services.NewDeleteDeviceUseCase(uc.repository)
	if err := deleteUC.Do(idStr); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}
