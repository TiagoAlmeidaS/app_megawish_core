package service_type

import "app_megawish_core/internal/domain/services/entities/service_type"

type IServiceTypeRepository interface {
	GetAll() ([]*service_type.ServiceType, error)
	GetByID(id int64) (*service_type.ServiceType, error)
	Create(serviceType *service_type.ServiceType) (*service_type.ServiceType, error)
	Update(serviceType *service_type.ServiceType) (*service_type.ServiceType, error)
	Delete(id int64) error
}
