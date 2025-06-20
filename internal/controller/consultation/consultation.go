package consultation

import (
	"context"
	"log"
	"medicare/api/v1/consultation"
	"medicare/internal/logic"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)
var Consultation = ConsultationController{}

type ConsultationController struct{}

func NewV1() consultation.ConsultationV1 {
	return &ConsultationController{}
}

func (c *ConsultationController) Create(ctx context.Context, req *consultation.CreateConsultationReq) (*consultation.CreateConsultationRes, error) {
	userID := gconv.String(gctx.Ctx(ctx).Value("UserID"))
	log.Println(userID)
	return logic.CreateConsultation(ctx, req, userID)
}

func (c *ConsultationController) List(ctx context.Context, req *consultation.ListConsultationsReq) (*consultation.ListConsultationsRes, error) {
	return logic.ListConsultations(ctx, req.PatientID)
}

func (c *ConsultationController) Get(ctx context.Context, req *consultation.GetConsultationReq) (*consultation.GetConsultationRes, error) {
	return logic.GetConsultation(ctx, req.ID)
}

func (c *ConsultationController) Update(ctx context.Context, req *consultation.UpdateConsultationReq) (*consultation.UpdateConsultationRes, error) {
	role := gconv.String(gctx.Ctx(ctx).Value("UserRole"))
	if role != "doctor" && role!= "receptionist" {
		return nil, gerror.NewCode(gcode.CodeNotAuthorized, "only doctors or receptionists can update consultations")
	}
	return logic.UpdateConsultation(ctx, req)
}

func (c *ConsultationController) Delete(ctx context.Context, req *consultation.DeleteConsultationReq) (*consultation.DeleteConsultationRes, error) {
	role := gconv.String(gctx.Ctx(ctx).Value("UserRole"))
	if role != "doctor" && role != "receptionist"{
		return nil, gerror.NewCode(gcode.CodeNotAuthorized, "only doctors or receptionists can delete consultations")
	}
	return logic.DeleteConsultation(ctx, req.ID)
}
