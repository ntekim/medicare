package patient

import (
	"context"
	"medicare/api/v1/patient"
	"medicare/internal/consts"
	"medicare/internal/logic"
	"strings"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

var Patient = PatientController{}

type PatientController struct{}

func NewV1() patient.PatientV1 {
	return &PatientController{}
}

func (c *PatientController) Create(ctx context.Context, req *patient.CreatePatientReq) (res *patient.CreatePatientRes, err error) {

	userRole := gconv.String(gctx.Ctx(ctx).Value(consts.CtxUserRole))
	if userRole != "doctor" {
		return nil, gerror.NewCode(gcode.CodeNotAuthorized, "forbidden: insufficient permission")
	}
	
	res, err = logic.Patient.Create(ctx, req)
	if err != nil {
		return nil, gerror.NewCode(gcode.CodeInternalError, err.Error())
	}
	return res, nil
}

func (c *PatientController) Get(ctx context.Context, req *patient.GetPatientReq) (res *patient.GetPatientRes, err error) {
	userRole := gconv.String(gctx.Ctx(ctx).Value(consts.CtxUserRole))
	role := strings.TrimSpace(userRole)
	if role != "doctor" && userRole != "receptionist" {
		return nil, gerror.NewCode(gcode.CodeNotAuthorized, "forbidden: insufficient permission")
	}
	res, err = logic.Patient.Get(ctx, req.ID)
	if err != nil {
		return nil, gerror.NewCode(gcode.CodeInternalError, err.Error())
	}
	return res, nil
}

func (c *PatientController) ListPatients(ctx context.Context, req *patient.ListPatientsReq) (*patient.ListPatientsRes, error) {
	userRole := gconv.String(gctx.Ctx(ctx).Value(consts.CtxUserRole))
	if userRole != "receptionist" && userRole != "doctor" {
		return nil, gerror.NewCode(gcode.CodeNotAuthorized, "forbidden: insufficient permission")
	}
	patients, err := logic.ListPatients(ctx, int32(req.Limit), int32(req.Offset))
	if err != nil {
		return nil, gerror.NewCode(gcode.CodeInternalError, err.Error())
	}
	return &patient.ListPatientsRes{
		Patients: patients,
	}, nil
}


func (c *PatientController) UpdatePatient(ctx context.Context, req *patient.UpdatePatientReq) (res *patient.UpdatePatientRes, err error) {
	userRole := gconv.String(gctx.Ctx(ctx).Value(consts.CtxUserRole))
	if userRole != "receptionist" {
		return nil, gerror.NewCode(gcode.CodeNotAuthorized, "forbidden: insufficient permission")
	}
	return logic.Patient.Update(ctx, req)
}

func (c *PatientController) DeletePatient(ctx context.Context, req *patient.DeletePatientReq) (res *patient.DeletePatientRes, err error) {
	userRole := gconv.String(gctx.Ctx(ctx).Value(consts.CtxUserRole))
	if userRole != "receptionist" {
		return nil, gerror.NewCode(gcode.CodeNotAuthorized, "forbidden: insufficient permission")
	}
	return logic.Patient.Delete(ctx, req.ID)
}
