package service

import "app_megawish_core/internal/domain/services/entities/service"

type IServiceUseCases interface {
	GetAll() ([]*service.Service, error)
	GetByID(id int64) (*service.Service, error)
	Create(service *service.Service) (*service.Service, error)
	Update(service *service.Service) (*service.Service, error)
	Delete(id int64) error
	GetAllByClientID(clientID int64) ([]*service.Service, error)
}
