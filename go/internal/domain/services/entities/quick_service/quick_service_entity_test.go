package quick_service

import "testing"

func TestNewQuickService(t *testing.T) {
	t.Run("should return error when service type id is empty", func(t *testing.T) {
		newQuickService, err := NewQuickService(0)

		if err == nil {
			t.Errorf("expected error %v, got %v", ErrorServiceTypeIDIsInvalid, err)
		}

		if newQuickService != nil {
			t.Errorf("expected quick_service %v, got %v", nil, newQuickService)
		}
	})

	t.Run("should return a new quick_service correctly", func(t *testing.T) {
		newQuickService, err := NewQuickService(1)

		if err != nil {
			t.Errorf("expected error %v, got %v", nil, err)
		}

		if newQuickService == nil {
			t.Errorf("expected quick_service %v, got %v", &QuickService{}, newQuickService)
		}
	})
}
