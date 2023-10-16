package service

import "app_megawish_core/internal/domain/services/entities/service"

type IServiceRepository interface {
	getAll() ([]*service.Service, error)
	getByID(id int64) (*service.Service, error)
	create(service *service.Service) (*service.Service, error)
	update(service *service.Service) (*service.Service, error)
	delete(id int64) error
	getAllByClientID(clientID int64) ([]*service.Service, error)
}
