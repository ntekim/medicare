package consultation

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

type Vitals struct {
	BloodPressure string `json:"blood_pressure" dc:"Blood pressure reading, e.g., 120/80"`
	Temperature   string `json:"temperature" dc:"Body temperature, e.g., 36.6°C"`
	Pulse         string `json:"pulse" dc:"Heart rate, e.g., 72 bpm"`
	Respiratory   string `json:"respiratory_rate" dc:"Respiratory rate, e.g., 18 bpm"`
	Weight        string `json:"weight" dc:"Weight, e.g., 70kg"`
}


// --- CREATE ---
type CreateConsultationReq struct {
	g.Meta        `path:"/" method:"post" tags:"Consultation" summary:"Create a new consultation"`
	PatientID     string `json:"patient_id" v:"required#Patient ID is required" dc:"Patient's UUID"`
	Vitals       Vitals  `json:"vitals" dc:"Patient vital signs"`
	Diagnosis     string `json:"diagnosis" dc:"Doctor's diagnosis"`
	Prescription  string `json:"prescription" dc:"Medications prescribed"`
	Notes         string `json:"notes" dc:"Additional consultation notes"`
}

type CreateConsultationRes struct {
	ID string `json:"id" dc:"UUID of the created consultation"`
}

// --- LIST ---
type ListConsultationsReq struct {
	g.Meta    `path:"/" method:"get" tags:"Consultation" summary:"List consultations by patient"`
	PatientID string `json:"patient_id" in:"query" v:"required#Patient ID required" dc:"Patient's UUID"`
}

type ListConsultationsRes struct {
	Consultations []GetConsultationRes `json:"consultations"`
}

// --- GET ---
type GetConsultationReq struct {
	g.Meta `path:"/{id}" method:"get" tags:"Consultation" summary:"Get consultation details"`
	ID     string `json:"id" v:"required#ID is required" dc:"Consultation UUID"`
}

type GetConsultationRes struct {
	ID               string `json:"id" dc:"Consultation UUID"`
	ConsultationDate string `json:"consultation_date" dc:"Date of consultation"`
	Vitals       Vitals  `json:"vitals" dc:"Patient vital signs"`
	Diagnosis        string `json:"diagnosis" dc:"Diagnosis made"`
	Prescription     string `json:"prescription" dc:"Prescription issued"`
	Notes            string `json:"notes" dc:"Additional notes"`
	DoctorID         string `json:"doctor_id" dc:"UUID of the doctor"`
}

// --- UPDATE ---
type UpdateConsultationReq struct {
	g.Meta        `path:"/{id}" method:"put" tags:"Consultation" summary:"Update a consultation"`
	ID            string `json:"id" in:"path" v:"required#ID is required" dc:"Consultation UUID"`
	Vitals       Vitals  `json:"vitals" dc:"Patient updated vital"`
	Diagnosis     string `json:"diagnosis" dc:"Updated diagnosis"`
	Prescription  string `json:"prescription" dc:"Updated prescription"`
	Notes         string `json:"notes" dc:"Updated notes"`
}

type UpdateConsultationRes struct {
	Message string `json:"message"`
}

// --- DELETE ---
type DeleteConsultationReq struct {
	g.Meta `path:"/{id}" method:"delete" tags:"Consultation" summary:"Delete a consultation"`
	ID     string `json:"id" v:"required#ID is required" dc:"Consultation UUID"`
}

type DeleteConsultationRes struct {
	Message string `json:"message"`
}

type ConsultationV1 interface {
	Create(ctx context.Context, req *CreateConsultationReq) (res *CreateConsultationRes, err error)
	Get(ctx context.Context, req *GetConsultationReq) (res *GetConsultationRes, err error)
	Update(ctx context.Context, req *UpdateConsultationReq) (res *UpdateConsultationRes, err error)
	Delete(ctx context.Context, req *DeleteConsultationReq) (res *DeleteConsultationRes, err error)
	List(ctx context.Context, req *ListConsultationsReq) (res *ListConsultationsRes, err error)
}
