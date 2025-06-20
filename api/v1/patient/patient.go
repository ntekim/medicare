package patient

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

type CreatePatientReq struct {
	g.Meta         `path:"/" method:"post" tags:"Patients" summary:"Create new patient"`
	FirstName      string `json:"first_name" v:"required"`
	LastName       string `json:"last_name" v:"required"`
	DateOfBirth    string `json:"date_of_birth"` // "YYYY-MM-DD"
	Gender         string `json:"gender"`
	ContactNumber  string `json:"contact_number"`
	Address        string `json:"address"`
}

type CreatePatientRes struct {
	ID string `json:"id"`
}

type GetPatientReq struct {
	g.Meta `path:"/{id}" method:"get" tags:"Patients" summary:"Get patient"`
	ID     string `json:"id"`
}

type GetPatientRes struct {
	ID            string `json:"id"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	DateOfBirth   string `json:"date_of_birth"`
	Gender        string `json:"gender"`
	ContactNumber string `json:"contact_number"`
	Address       string `json:"address"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}

type ListPatientsReq struct {
	g.Meta `path:"/" method:"get" tags:"Patients" summary:"List patients"`
	Limit  int `json:"limit"  d:"10"`
	Offset int `json:"offset" d:"0"`
}

type ListPatientsRes struct {
	Patients []GetPatientRes `json:"patients"`
}

type UpdatePatientReq struct {
	g.Meta        `path:"/{id}" method:"put" tags:"Patients" summary:"Update patient"`
	ID            string `json:"id"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	DateOfBirth   string `json:"date_of_birth"`
	Gender        string `json:"gender"`
	ContactNumber string `json:"contact_number"`
	Address       string `json:"address"`
}
type UpdatePatientRes struct {
	Message string `json:"message"`
}

type DeletePatientReq struct {
	g.Meta `path:"/{id}" method:"delete" tags:"Patients" summary:"Delete patient"`
	ID     string `json:"id"`
}

type DeletePatientRes struct {
	Message string `json:"message"`
}

type PatientV1 interface {
	Create(ctx context.Context, req *CreatePatientReq) (res *CreatePatientRes, err error)
	Get(ctx context.Context, req *GetPatientReq) (res *GetPatientRes, err error)
	UpdatePatient(ctx context.Context, req *UpdatePatientReq) (res *UpdatePatientRes, err error)
	DeletePatient(ctx context.Context, req *DeletePatientReq) (res *DeletePatientRes, err error)
	ListPatients(ctx context.Context, req *ListPatientsReq) (res *ListPatientsRes, err error)
}
