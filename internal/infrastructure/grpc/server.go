package gapi

import (
	"context"
	"errors"
	"time"
	"tunelport-svc-devices/internal/domain/persistence"
	"tunelport-svc-devices/internal/domain/services"
	"tunelport-svc-devices/internal/domain/types"

	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

type Server struct {
	UnimplementedDeviceServiceServer
	repository persistence.DeviceRepository
}

func mapDomainError(err error, internalMsg string) error {
	switch {
	case errors.Is(err, types.ErrDeviceAlreadyExists):
		return status.Errorf(codes.AlreadyExists, "device already exists: %v", err)
	case errors.Is(err, types.ErrDeviceNotFound):
		return status.Errorf(codes.NotFound, "device not found: %v", err)
	case errors.Is(err, types.ErrInvalidStatus):
		return status.Errorf(codes.InvalidArgument, "invalid status: %v", err)
	default:
		return status.Errorf(codes.Internal, "%s: %v", internalMsg, err)
	}
}
func NewServer(repository persistence.DeviceRepository) *Server {
	return &Server{
		repository: repository,
	}
}

func (s *Server) CreateDevice(ctx context.Context, req *CreateDeviceRequest) (*CreateDeviceResponse, error) {
	if req == nil || req.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "name is required")
	}

	service := services.NewSaveDeviceUseCase(s.repository)
	id, err := service.Do(req.Name, time.Now(), time.Now())
	if err != nil {
		return nil, mapDomainError(err, "failed to create device")
	}

	return &CreateDeviceResponse{
		Id:   id,
		Name: req.Name,
	}, nil
}

func (s *Server) GetDevice(ctx context.Context, req *GetDeviceRequest) (*GetDeviceResponse, error) {
	if req == nil || req.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}

	service := services.NewGetDeviceUseCase(s.repository)
	device, err := service.Do(req.Id)
	if err != nil {
		return nil, mapDomainError(err, "failed to get device")
	}

	return &GetDeviceResponse{
		Id:   device.ID.Value(),
		Name: device.Name.Value(),
	}, nil
}

func (s *Server) UpdateDevice(ctx context.Context, req *UpdateDeviceRequest) (*UpdateDeviceResponse, error) {
	if req == nil || req.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}
	if req.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "name is required")
	}

	service := services.NewUpdateDeviceUseCase(s.repository)
	err := service.Do(req.Id, req.Name, time.Now(), time.Now())
	if err != nil {
		return nil, mapDomainError(err, "failed to update device")
	}

	return &UpdateDeviceResponse{
		Updated: time.Now().Format(time.RFC3339),
	}, nil
}

func (s *Server) DeleteDevice(ctx context.Context, req *DeleteDeviceRequest) (*DeleteDeviceResponse, error) {
	if req == nil || req.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}

	service := services.NewDeleteDeviceUseCase(s.repository)
	err := service.Do(req.Id)
	if err != nil {
		return nil, mapDomainError(err, "failed to delete device")
	}

	return &DeleteDeviceResponse{
		Deleted: time.Now().Format(time.RFC3339),
	}, nil
}

func (s *Server) ListDevices(ctx context.Context, req *ListDevicesRequest) (*ListDevicesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request is required")
	}

	service := services.NewListDeviceUseCase(s.repository)
	devices, err := service.Do()
	if err != nil {
		return nil, mapDomainError(err, "failed to list devices")
	}

	items := make([]*Device, 0, len(devices))
	for _, t := range devices {
		items = append(items, &Device{
			Id:   t.ID.Value(),
			Name: t.Name.Value(),
		})
	}

	return &ListDevicesResponse{
		Devices: items,
	}, nil
}
