package logic_test

import (
	"context"
	"testing"
	"time"

	postgres "medicare/internal/dao/sqlc"
	"medicare/internal/logic"
	"medicare/utility/helpers"
	mockdb "medicare/internal/mocks"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"medicare/api/v1/patient"
)

func TestCreatePatient(t *testing.T) {
	mockQ := new(mockdb.Querier)
	logic.Queries = mockQ

	req := &patient.CreatePatientReq{
		FirstName:     "John",
		LastName:      "Doe",
		Gender:        "Male",
		DateOfBirth:   "2000-01-01",
		ContactNumber: "1234567890",
		Address:       "123 Street",
	}

	mockQ.On("CreatePatient", mock.Anything, mock.AnythingOfType("db.CreatePatientParams")).
		Return(uuid.New(), nil).Once()

	res, err := logic.Patient.Create(context.Background(), req)
	assert.NoError(t, err)
	assert.NotEmpty(t, res.ID)

	mockQ.AssertExpectations(t)
}

func TestGetPatient(t *testing.T) {
	mockQ := new(mockdb.Querier)
	logic.Queries = mockQ

	id := uuid.New()
	now := time.Now()

	mockQ.On("GetPatientByID", mock.Anything, id).
		Return(postgres.Patient{
			ID:            id,
			FirstName:     "John",
			LastName:      "Doe",
			Gender:        "Male",
			DateOfBirth:   helpers.ToPgDate(time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)),
			ContactNumber: "1234567890",
			Address:       "123 Street",
			CreatedAt:     helpers.ToPgTime(now),
			UpdatedAt:     helpers.ToPgTime(now),
		}, nil)

	res, err := logic.Patient.Get(context.Background(), id.String())
	assert.NoError(t, err)
	assert.Equal(t, res.FirstName, "John")

	mockQ.AssertExpectations(t)
}

func TestUpdatePatient(t *testing.T) {
	mockQ := new(mockdb.Querier)
	logic.Queries = mockQ

	id := uuid.New()
	req := &patient.UpdatePatientReq{
		ID:            id.String(),
		FirstName:     "Updated",
		LastName:      "Patient",
		DateOfBirth:   "2000-01-01",
		Gender:        "Female",
		ContactNumber: "9876543210",
		Address:       "456 Street",
	}

	mockQ.On("UpdatePatient", mock.Anything, mock.AnythingOfType("db.UpdatePatientParams")).
		Return(nil).Once()

	res, err := logic.Patient.Update(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, "patient updated", res.Message)

	mockQ.AssertExpectations(t)
}

func TestDeletePatient(t *testing.T) {
	mockQ := new(mockdb.Querier)
	logic.Queries = mockQ

	id := uuid.New()

	mockQ.On("DeletePatient", mock.Anything, id).
		Return(nil).Once()

	res, err := logic.Patient.Delete(context.Background(), id.String())
	assert.NoError(t, err)
	assert.Equal(t, "patient deleted", res.Message)

	mockQ.AssertExpectations(t)
}

func TestListPatients(t *testing.T) {
	mockQ := new(mockdb.Querier)
	logic.Queries = mockQ

	now := time.Now()

	mockQ.On("ListPatients", mock.Anything, postgres.ListPatientsParams{
		Limit:  10,
		Offset: 0,
	}).Return([]postgres.Patient{
		{
			ID:            uuid.New(),
			FirstName:     "Jane",
			LastName:      "Smith",
			DateOfBirth:   helpers.ToPgDate(time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)),
			Gender:        "Female",
			ContactNumber: "1234567890",
			Address:       "789 Road",
			CreatedAt:     helpers.ToPgTime(now),
			UpdatedAt:     helpers.ToPgTime(now),
		},
	}, nil).Once()

	res, err := logic.ListPatients(context.Background(), 10, 0)
	assert.NoError(t, err)
	assert.Len(t, res, 1)
	assert.Equal(t, "Jane", res[0].FirstName)

	mockQ.AssertExpectations(t)
}

func TestCreatePatient_InvalidDate(t *testing.T) {
	// Skip DB call — this fails before DB interaction
	req := &patient.CreatePatientReq{
		DateOfBirth: "invalid-date",
	}

	_, err := logic.Patient.Create(context.Background(), req)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid date format")
}
