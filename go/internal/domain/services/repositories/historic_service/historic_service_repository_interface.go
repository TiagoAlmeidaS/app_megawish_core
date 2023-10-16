package historic_service

import "app_megawish_core/internal/domain/services/entities/historic_service"

type IHistoricServiceRepository interface {
	getAll() ([]*historic_service.HistoricService, error)
	getByID(id int64) (*historic_service.HistoricService, error)
	create(historicService *historic_service.HistoricService) (*historic_service.HistoricService, error)
	update(historicService *historic_service.HistoricService) (*historic_service.HistoricService, error)
	delete(id int64) error
	getAllByClientID(clientID int64) ([]*historic_service.HistoricService, error)
}
