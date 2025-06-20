package logic

import (
	"context"
	"encoding/json"
	"time"

	"medicare/api/v1/consultation"
	db "medicare/internal/dao/sqlc"

	"medicare/utility/helpers"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/google/uuid"
)
func CreateConsultation(ctx context.Context, req *consultation.CreateConsultationReq, doctorID string) (*consultation.CreateConsultationRes, error) {
	id := uuid.New()
	vitalsJSON, err := json.Marshal(req.Vitals)
	if err != nil {
		return nil, gerror.NewCode(gcode.CodeInternalError, "Invalid vitals JSON")
	}
	resp, err := Queries.CreateConsultation(ctx, db.CreateConsultationParams{
		ID:           id,
		PatientID:    uuid.MustParse(req.PatientID),
		DoctorID:     uuid.MustParse(doctorID),
		Vitals:       vitalsJSON,
		Diagnosis:    helpers.ToPgText(req.Diagnosis),
		Prescription: helpers.ToPgText(req.Prescription),
		Notes:        helpers.ToPgText(req.Notes),
	})
	if err != nil {
		return nil, gerror.Wrap(err, "failed to create consultation")
	}
	res := &consultation.CreateConsultationRes{ID: resp.ID.String()}
	return res, nil
}

func ListConsultations(ctx context.Context, patientID string) (*consultation.ListConsultationsRes, error) {
	rows, err := Queries.ListConsultationsByPatient(ctx, uuid.MustParse(patientID))
	if err != nil {
		return nil, gerror.Wrap(err, "failed to list consultations")
	}
	var res []consultation.GetConsultationRes
	for _, row := range rows {
		var vitals consultation.Vitals
		if err := json.Unmarshal(row.Vitals, &vitals); err != nil {
			return nil, gerror.NewCode(gcode.CodeInternalError, "invalid vitals format")
		}
		res = append(res, consultation.GetConsultationRes{
			ID:               row.ID.String(),
			ConsultationDate: row.ConsultationDate.Time.Format(time.RFC3339),
			Vitals:       vitals,
			Diagnosis:    helpers.PgTextToString(row.Diagnosis),
			Prescription: helpers.PgTextToString(row.Prescription),
			Notes:        helpers.PgTextToString(row.Notes),
			DoctorID:     row.DoctorID.String(),
		})
	}
	resp := &consultation.ListConsultationsRes{
		Consultations: res,
	}
	return resp, nil
}

func GetConsultation(ctx context.Context, id string) (*consultation.GetConsultationRes, error) {
	row, err := Queries.GetConsultation(ctx, uuid.MustParse(id))
	if err != nil {
		return nil, gerror.Wrap(err, "not found")
	}
	var vitals consultation.Vitals
	if err := json.Unmarshal(row.Vitals, &vitals); err != nil {
		return nil, gerror.NewCode(gcode.CodeInternalError, "invalid vitals format")
	}
	return &consultation.GetConsultationRes{
		ID:               row.ID.String(),
		ConsultationDate: row.ConsultationDate.Time.Format(time.RFC3339),
		Vitals:       vitals,
		Diagnosis:    helpers.PgTextToString(row.Diagnosis),
		Prescription: helpers.PgTextToString(row.Prescription),
		Notes:        helpers.PgTextToString(row.Notes),
		DoctorID:         row.DoctorID.String(),
	}, nil
}

func UpdateConsultation(ctx context.Context, req *consultation.UpdateConsultationReq) (*consultation.UpdateConsultationRes, error) {
	vitalsJSON, err := json.Marshal(req.Vitals)
	if err != nil {
		return &consultation.UpdateConsultationRes{}, gerror.NewCodef(gcode.CodeInternalError, "Invalid vitals JSON: %v", err.Error())
	}
	err = Queries.UpdateConsultation(ctx, db.UpdateConsultationParams{
		ID:           uuid.MustParse(req.ID),
		Vitals:       vitalsJSON,
		Diagnosis:    helpers.ToPgText(req.Diagnosis),
		Prescription: helpers.ToPgText(req.Prescription),
		Notes:        helpers.ToPgText(req.Notes),
	})

	return &consultation.UpdateConsultationRes{
		Message: "consultation updated",
	}, err
}

func DeleteConsultation(ctx context.Context, id string) (*consultation.DeleteConsultationRes, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return &consultation.DeleteConsultationRes{}, err
	}
	err = Queries.DeletePatient(ctx, uid)
	return &consultation.DeleteConsultationRes{
		Message: "consultation deleted",
	}, err
}
