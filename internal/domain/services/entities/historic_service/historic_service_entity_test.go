package historic_service

import (
	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewHistoricService(t *testing.T) {
	t.Run("should return a new historic_service correctly", func(t *testing.T) {
		timeNow := time.Now()
		valor, serviceId, clientId, statusId := faker.Latitude(), faker.RandomUnixTime(), faker.RandomUnixTime(), faker.RandomUnixTime()
		newHistoricService, err := NewHistoricService(timeNow, timeNow, valor, serviceId, clientId, statusId)

		assert.Nil(t, err)
		assert.Equal(t, &timeNow, newHistoricService.DataInicio)
		assert.Equal(t, &timeNow, newHistoricService.DataFim)
		assert.Equal(t, valor, newHistoricService.Valor)
		assert.Equal(t, serviceId, newHistoricService.ServiceID)
		assert.Equal(t, clientId, newHistoricService.ClientID)
		assert.Equal(t, statusId, newHistoricService.StatusID)
	})

	t.Run("should return a error when dataInicio is zero", func(t *testing.T) {
		timeZero := time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)
		newHistoricService, err := NewHistoricService(timeZero, timeZero, faker.Latitude(), faker.RandomUnixTime(), faker.RandomUnixTime(), faker.RandomUnixTime())

		assert.Nil(t, newHistoricService)
		assert.Equal(t, err, ErrDataInicioIsInvalid)
	})

	t.Run("should return a error when dataFim is before dataInicio", func(t *testing.T) {
		timeNow := time.Now()
		timeSub := timeNow.AddDate(-1, 1, 1)
		newHistoricService, err := NewHistoricService(timeNow, timeSub, faker.Latitude(), faker.RandomUnixTime(), faker.RandomUnixTime(), faker.RandomUnixTime())

		assert.Nil(t, newHistoricService)
		assert.Equal(t, err, ErrDataFimIsInvalid)
	})

	t.Run("should return a error when valor is zero", func(t *testing.T) {
		timeNow := time.Now()
		newHistoricService, err := NewHistoricService(timeNow, timeNow.AddDate(1, 1, 1), 0, faker.RandomUnixTime(), faker.RandomUnixTime(), faker.RandomUnixTime())

		assert.Nil(t, newHistoricService)
		assert.Equal(t, err, ErrValorIsInvalid)
	})

	t.Run("should return a error when serviceID is zero", func(t *testing.T) {
		timeNow := time.Now()
		newHistoricService, err := NewHistoricService(timeNow, timeNow.AddDate(1, 1, 1), faker.Latitude(), 0, faker.RandomUnixTime(), faker.RandomUnixTime())

		assert.Nil(t, newHistoricService)
		assert.Equal(t, err, ErrServiceIDIsInvalid)
	})

	t.Run("should return a error when clientID is zero", func(t *testing.T) {
		timeNow := time.Now()
		newHistoricService, err := NewHistoricService(timeNow, timeNow.AddDate(1, 1, 1), faker.Latitude(), faker.RandomUnixTime(), 0, faker.RandomUnixTime())

		assert.Nil(t, newHistoricService)
		assert.Equal(t, err, ErrClientIDIsInvalid)
	})
}
