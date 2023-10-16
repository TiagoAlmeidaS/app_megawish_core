package quick_service

import "app_megawish_core/internal/domain/services/entities/quick_service"

type IQuickServiceUseCases interface {
	GetAll() ([]*quick_service.QuickService, error)
	GetByID(id int64) (*quick_service.QuickService, error)
	Create(quickService *quick_service.QuickService) (*quick_service.QuickService, error)
	Update(quickService *quick_service.QuickService) (*quick_service.QuickService, error)
	Delete(id int64) error
}
