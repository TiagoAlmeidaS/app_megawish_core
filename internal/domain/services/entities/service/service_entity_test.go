package service

import (
	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
	"testing"
)

func FuzzNewService(f *testing.F) {

	f.Add(faker.RandomUnixTime(), faker.Name(), faker.Name())
	f.Add(faker.RandomUnixTime(), "", faker.Name())
	f.Add(faker.RandomUnixTime(), faker.Name(), "")

	f.Fuzz(func(t *testing.T, serviceTypeID int64, name, description string) {
		newService, err := NewService(serviceTypeID, name, description)

		if description == "" && err != nil {
			assert.Nil(t, newService)
			assert.Equal(t, err, ErrDescriptionServiceIsInvalid)
			return
		}

		if name == "" && err != nil {
			assert.Nil(t, newService)
			assert.Equal(t, err, ErrNameServiceIsInvalid)
			return
		}

		assert.Nil(t, err)
		assert.Equal(t, name, newService.Name)
		assert.Equal(t, description, newService.Description)
		assert.Equal(t, serviceTypeID, newService.ServiceTypeID)
	})
}
