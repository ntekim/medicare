package cmd

import (
	"context"

	"medicare/internal/controller/auth"
	"medicare/internal/controller/hello"
	"medicare/internal/controller/patient"
	"medicare/internal/controller/consultation"
	"medicare/internal/middlewares"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					new(hello.ControllerV1),
				)
			})
			s.Group("/auth", func(group *ghttp.RouterGroup) {
				group.Middleware(middlewares.HandlerMiddleware)
				group.Bind(
					new(auth.LAuth),
				)
			})

			s.Group("/patients", func(group *ghttp.RouterGroup) {
				group.Middleware(middlewares.Auth, middlewares.HandlerMiddleware)
				group.Bind(
					new(patient.PatientController),
				)
			})

			s.Group("/consultations", func(group *ghttp.RouterGroup) {
				group.Middleware(middlewares.Auth, middlewares.RequireRole("doctor"), middlewares.HandlerMiddleware)
				group.Bind(
					new(consultation.ConsultationController),
				)
			})
		
			s.SetOpenApiPath("/api.json")
			s.SetSwaggerPath("/swagger")
			s.Run()
			return nil
		},
	}
)
