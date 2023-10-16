package service

import (
	"errors"
)

var (
	ErrNameServiceIsInvalid        = errors.New("name don't to be empty")
	ErrDescriptionServiceIsInvalid = errors.New("description don't to be empty")
	ErrTypeServiceIsInvalid        = errors.New("service type don't to be empty")
)

type Service struct {
	ID            int64
	Name          string
	Description   string
	ServiceTypeID int64
}

func NewService(serviceTypeID int64, name, description string) (*Service, error) {
	if name == "" {
		return nil, ErrNameServiceIsInvalid
	}

	if description == "" {
		return nil, ErrDescriptionServiceIsInvalid
	}

	if serviceTypeID == 0 {
		return nil, ErrTypeServiceIsInvalid
	}

	return &Service{
		Name:          name,
		Description:   description,
		ServiceTypeID: serviceTypeID,
	}, nil
}
