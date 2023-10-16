package service_type

import "testing"

func TestNewServiceType(t *testing.T) {
	t.Run("should return error when name is empty", func(t *testing.T) {
		newServiceType, err := NewServiceType("", "description")

		if err == nil {
			t.Errorf("expected error %v, got %v", ErrNameIsInvalid, err)
		}

		if newServiceType != nil {
			t.Errorf("expected service_type %v, got %v", nil, newServiceType)
		}
	})

	t.Run("should return error when description is empty", func(t *testing.T) {
		newServiceType, err := NewServiceType("name", "")

		if err == nil {
			t.Errorf("expected error %v, got %v", ErrDescriptionIsInvalid, err)
		}

		if newServiceType != nil {
			t.Errorf("expected service_type %v, got %v", nil, newServiceType)
		}
	})

	t.Run("should return a new service_type correctly", func(t *testing.T) {
		newServiceType, err := NewServiceType("name", "description")

		if err != nil {
			t.Errorf("expected error %v, got %v", nil, err)
		}

		if newServiceType == nil {
			t.Errorf("expected service_type %v, got %v", &ServiceType{}, newServiceType)
		}
	})
}
