package historic_service

import (
	"errors"
	"time"
)

var (
	ErrDataInicioIsInvalid = errors.New("date init don't to be empty")
	ErrDataFimIsInvalid    = errors.New("date end don't to be before date init")
	ErrValorIsInvalid      = errors.New("currency value don't to be empty")
	ErrServiceIDIsInvalid  = errors.New("serviceID type don't to be empty")
	ErrClientIDIsInvalid   = errors.New("clientID type don't to be empty")
)

type HistoricService struct {
	ID         int64
	DataInicio *time.Time
	DataFim    *time.Time
	Valor      float64
	ServiceID  int64
	ClientID   int64
	StatusID   int64
}

func NewHistoricService(dataInicio time.Time, dataFim time.Time, valor float64, serviceId int64, clientId int64, statusId int64) (*HistoricService, error) {
	if dataInicio.IsZero() {
		return nil, ErrDataInicioIsInvalid
	}

	if dataFim.Before(dataInicio) {
		return nil, ErrDataFimIsInvalid
	}

	if valor == 0 {
		return nil, ErrValorIsInvalid
	}

	if serviceId == 0 {
		return nil, ErrServiceIDIsInvalid
	}

	if clientId == 0 {
		return nil, ErrClientIDIsInvalid
	}

	return &HistoricService{
		DataInicio: &dataInicio,
		DataFim:    &dataFim,
		Valor:      valor,
		ServiceID:  serviceId,
		ClientID:   clientId,
		StatusID:   statusId,
	}, nil
}
