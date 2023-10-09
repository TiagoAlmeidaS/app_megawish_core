package quick_service

import "errors"

var (
	ErrorServiceTypeIDIsInvalid = errors.New("service type don't to be empty")
)

type QuickService struct {
	ID            int64
	ServiceTypeID int64
}

func NewQuickService(serviceTypeID int64) (*QuickService, error) {

	if serviceTypeID == 0 {
		return nil, ErrorServiceTypeIDIsInvalid
	}

	return &QuickService{
		ServiceTypeID: serviceTypeID,
	}, nil
}
