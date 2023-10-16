package quick_service

import "app_megawish_core/internal/domain/services/entities/quick_service"

type IQuickServiceRepository interface {
	getAll() ([]*quick_service.QuickService, error)
	getByID(id int64) (*quick_service.QuickService, error)
	create(quickService *quick_service.QuickService) (*quick_service.QuickService, error)
	update(quickService *quick_service.QuickService) (*quick_service.QuickService, error)
	delete(id int64) error
}
