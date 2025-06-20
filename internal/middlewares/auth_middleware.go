package middlewares

import (
	"strings"

	"medicare/utility/helpers"

	"medicare/internal/consts"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func Auth(r *ghttp.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		r.Response.WriteStatusExit(401, g.Map{"message": "Missing or invalid Authorization header"})
		return
	}

	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

	// Validate token
	claims, err := helpers.Verify(tokenStr)
	if err != nil {
		r.Response.WriteStatusExit(401, g.Map{"message": "Invalid or expired token"})
		return
	}

	// Inject user context
	r.SetCtxVar(consts.CtxUserID, claims.UserID)
	r.SetCtxVar(consts.CtxUserRole, claims.Role)

	r.Middleware.Next()
}
