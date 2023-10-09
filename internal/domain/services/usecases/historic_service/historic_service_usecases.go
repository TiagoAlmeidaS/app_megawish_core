package historic_service

import "app_megawish_core/internal/domain/services/entities/historic_service"

type IHistoricServiceUseCases interface {
	GetAll() ([]*historic_service.HistoricService, error)
	GetByID(id int64) (*historic_service.HistoricService, error)
	Create(historicService *historic_service.HistoricService) (*historic_service.HistoricService, error)
	Update(historicService *historic_service.HistoricService) (*historic_service.HistoricService, error)
	Delete(id int64) error
	GetAllByClientID(clientID int64) ([]*historic_service.HistoricService, error)
}
