package service_type

import "errors"

var (
	ErrNameIsInvalid        = errors.New("name don't to be empty")
	ErrDescriptionIsInvalid = errors.New("description don't to be empty")
)

type ServiceType struct {
	ID          int64
	Name        string
	Description string
}

func NewServiceType(name, description string) (*ServiceType, error) {

	if name == "" {
		return nil, ErrNameIsInvalid
	}

	if description == "" {
		return nil, ErrDescriptionIsInvalid
	}

	return &ServiceType{
		Name:        name,
		Description: description,
	}, nil
}
