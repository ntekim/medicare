package logic

import (
	"context"
	"fmt"
	db "medicare/internal/dao/sqlc" // sqlc package
	"time"

	"medicare/api/v1/patient"
	"medicare/utility/helpers"

	"github.com/google/uuid"
)

var Patient = patientLogic{}

type patientLogic struct{}

func (p patientLogic) Create(ctx context.Context, req *patient.CreatePatientReq) (res *patient.CreatePatientRes, err error) {
	dob, parseErr := time.Parse("2006-01-02", req.DateOfBirth)
	if parseErr != nil {
		err = parseErr
		return nil, fmt.Errorf("invalid date format: %w", err)
	}

	resp, respErr := Queries.CreatePatient(ctx, db.CreatePatientParams{
		ID: uuid.New(),
		FirstName: req.FirstName,
		LastName: req.LastName,
		Gender: req.Gender,
		DateOfBirth: helpers.ToPgDate(dob),
		ContactNumber: req.ContactNumber,
		Address: req.Address,
	})
	if respErr != nil {
		err = respErr
		return nil, err
	}

	res = &patient.CreatePatientRes{
		ID: resp.String(),
	}
	return res, nil
}

func (p patientLogic) Get(ctx context.Context, id string) (*patient.GetPatientRes, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	res, err := Queries.GetPatientByID(ctx, uid)
	if err != nil {
		return nil, err
	}
	resp := &patient.GetPatientRes{
		ID: res.ID.String(),
		FirstName: res.FirstName,
		LastName: res.LastName,
		DateOfBirth: helpers.PgDateToString(res.DateOfBirth),
		Gender: res.Gender,
		ContactNumber: res.ContactNumber,
		Address: res.Address,
		CreatedAt:     res.CreatedAt.Time.Format(time.RFC3339),
	    UpdatedAt:     res.UpdatedAt.Time.Format(time.RFC3339),
	}
	return resp, nil
}

func ListPatients(ctx context.Context, limit, offset int32) ([]patient.GetPatientRes, error) {
	rows, err := Queries.ListPatients(ctx, db.ListPatientsParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return []patient.GetPatientRes{}, err
	}

	patients := make([]patient.GetPatientRes, 0, len(rows))
	for _, row := range rows {
		patients = append(patients, patient.GetPatientRes{
			ID:            row.ID.String(),
			FirstName:     row.FirstName,
			LastName:      row.LastName,
			DateOfBirth:   helpers.PgDateToString(row.DateOfBirth),
			Gender:        row.Gender,
			ContactNumber: row.ContactNumber,
			Address:       row.Address,
			CreatedAt:     helpers.PgTimeToString(row.CreatedAt),
			UpdatedAt:     helpers.PgTimeToString(row.UpdatedAt),
		})
	}
	return patients, nil
}

func (p patientLogic) Update(ctx context.Context, req *patient.UpdatePatientReq) (*patient.UpdatePatientRes, error) {
	id, err := uuid.Parse(req.ID)
	if err != nil {
		return nil, err
	}

	dob, err := time.Parse("2006-01-02", req.DateOfBirth)
	if err != nil {
		return nil, fmt.Errorf("invalid date format: %w", err)
	}

	err = Queries.UpdatePatient(ctx, db.UpdatePatientParams{
		ID:            id,
		FirstName:     req.FirstName,
		LastName:      req.LastName,
		DateOfBirth:   helpers.ToPgDate(dob),
		Gender:        req.Gender,
		ContactNumber: req.ContactNumber,
		Address:       req.Address,
	})

	return &patient.UpdatePatientRes{
		Message: "patient updated",
	}, err
}

func (p patientLogic) Delete(ctx context.Context, id string) (*patient.DeletePatientRes, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return &patient.DeletePatientRes{}, err
	}
	err = Queries.DeletePatient(ctx, uid)
	return &patient.DeletePatientRes{
		Message: "patient deleted",
	}, err
}
