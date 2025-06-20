package logic_test

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/google/uuid"
	"medicare/utility/helpers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"medicare/api/v1/consultation"
	postgres "medicare/internal/dao/sqlc"
	"medicare/internal/logic"
	mocksqlc "medicare/internal/mocks"
)

func TestCreateConsultation_Success(t *testing.T) {
	mockQ := new(mocksqlc.Querier)
	logic.Queries = mockQ

	ctx := context.Background()
	doctorID := uuid.New()
	patientID := uuid.New()
	consultID := uuid.New()

	vitals := consultation.Vitals{
		Temperature: "36.5°C",
		BloodPressure: "120/80",
	}

	req := &consultation.CreateConsultationReq{
		PatientID: patientID.String(),
		Vitals: vitals,
		Diagnosis: "Flu",
		Prescription: "Rest + Paracetamol",
		Notes: "No allergies",
	}

	vitalsJSON, _ := json.Marshal(vitals)

	mockQ.On("CreateConsultation", mock.Anything, mock.MatchedBy(func(p postgres.CreateConsultationParams) bool {
		return p.PatientID == patientID && p.DoctorID == doctorID && string(p.Vitals) == string(vitalsJSON)
	})).Return(postgres.Consultation{ID: consultID}, nil)

	res, err := logic.CreateConsultation(ctx, req, doctorID.String())
	assert.NoError(t, err)
	assert.Equal(t, consultID.String(), res.ID)
}

func TestListConsultations_Success(t *testing.T) {
	mockQ := new(mocksqlc.Querier)
	logic.Queries = mockQ

	ctx := context.Background()
	patientID := uuid.New()
	doctorID := uuid.New()
	consultID := uuid.New()

	vitals := consultation.Vitals{
		Temperature: "36.5",
		BloodPressure: "120/80",
	}
	vitalsJSON, _ := json.Marshal(vitals)

	mockQ.On("ListConsultationsByPatient", mock.Anything, patientID).
		Return([]postgres.Consultation{
			{
				ID: consultID,
				ConsultationDate: helpers.SqlNullTime(time.Now()),
				DoctorID: doctorID,
				Vitals: vitalsJSON,
				Diagnosis: helpers.ToPgText("Cold"),
				Prescription: helpers.ToPgText("Panadol"),
				Notes: helpers.ToPgText("Test note"),
			},
		}, nil)

	res, err := logic.ListConsultations(ctx, patientID.String())
	assert.NoError(t, err)
	assert.Len(t, res.Consultations, 1)
	assert.Equal(t, consultID.String(), res.Consultations[0].ID)
}


func TestGetConsultation_Success(t *testing.T) {
	mockQ := new(mocksqlc.Querier)
	logic.Queries = mockQ

	ctx := context.Background()
	consultID := uuid.New()
	doctorID := uuid.New()
	vitals := consultation.Vitals{
		Temperature: "37",
	}
	vitalsJSON, _ := json.Marshal(vitals)

	mockQ.On("GetConsultation", mock.Anything, consultID).Return(postgres.Consultation{
		ID: consultID,
		DoctorID: doctorID,
		Vitals: vitalsJSON,
		Diagnosis: helpers.ToPgText("Malaria"),
		Prescription: helpers.ToPgText("ACT"),
		Notes: helpers.ToPgText("Test"),
		ConsultationDate: helpers.SqlNullTime(time.Now()),
	}, nil)

	res, err := logic.GetConsultation(ctx, consultID.String())
	assert.NoError(t, err)
	assert.Equal(t, consultID.String(), res.ID)
}

func TestUpdateConsultation_Success(t *testing.T) {
	mockQ := new(mocksqlc.Querier)
	logic.Queries = mockQ

	ctx := context.Background()
	consultID := uuid.New()
	vitals := consultation.Vitals{Temperature: "38"}
	// vitalsJSON, _ := json.Marshal(vitals)

	mockQ.On("UpdateConsultation", mock.Anything, mock.AnythingOfType("postgres.UpdateConsultationParams")).
		Return(nil)

	res, err := logic.UpdateConsultation(ctx, &consultation.UpdateConsultationReq{
		ID: consultID.String(),
		Vitals: vitals,
		Diagnosis: "Fever",
		Prescription: "Tylenol",
		Notes: "Patient improving",
	})

	assert.NoError(t, err)
	assert.Equal(t, "consultation updated", res.Message)
}

func TestDeleteConsultation_Success(t *testing.T) {
	mockQ := new(mocksqlc.Querier)
	logic.Queries = mockQ

	ctx := context.Background()
	id := uuid.New()

	mockQ.On("DeletePatient", ctx, id).Return(nil)

	res, err := logic.DeleteConsultation(ctx, id.String())
	assert.NoError(t, err)
	assert.Equal(t, "consultation deleted", res.Message)
}
