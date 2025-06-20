package helpers

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"net/http"
)

func MapGcodeToHTTPStatus(code gcode.Code) int {
	switch code {
	case gcode.CodeInvalidParameter:
		return http.StatusBadRequest
	case gcode.CodeNotAuthorized:
		return http.StatusForbidden
	case gcode.CodeInternalError:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}
