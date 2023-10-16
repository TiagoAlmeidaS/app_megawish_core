package service_type

import "app_megawish_core/internal/domain/services/entities/service_type"

type IServiceTypeRepository interface {
	getAll() ([]*service_type.ServiceType, error)
	getByID(id int64) (*service_type.ServiceType, error)
	create(serviceType *service_type.ServiceType) (*service_type.ServiceType, error)
	update(serviceType *service_type.ServiceType) (*service_type.ServiceType, error)
	delete(id int64) error
}
