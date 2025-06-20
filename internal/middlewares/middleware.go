package middlewares

import (
	"medicare/utility/helpers"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
)

func HandlerMiddleware(r *ghttp.Request) {

	r.Middleware.Next()

	var (
		res = r.GetHandlerResponse()
		err = r.GetError()
	)

	if err != nil {
		code := gerror.Code(err)

		// Set actual HTTP status code
		httpCode := helpers.MapGcodeToHTTPStatus(code)

		r.Response.WriteStatus(httpCode)
		r.Response.WriteJson(
			helpers.Response{
			Code: code.Code(),
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	r.Response.WriteJson(
		helpers.Response{
			Code: 0,
			Message: "OK",
			Data:    res,
		},
	)
}
