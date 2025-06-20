package middlewares

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"medicare/internal/consts"
)

func RequireRole(role string) ghttp.HandlerFunc {
	return func(r *ghttp.Request) {
		userRole := gconv.String(r.GetCtxVar(consts.CtxUserRole))
		if userRole != role {
			r.Response.WriteStatusExit(403, g.Map{"message": "Forbidden: insufficient permission"})
			return
		}
		r.Middleware.Next()
	}
}

func RequireRoleAny(roles ...string) ghttp.HandlerFunc {
	roleSet := make(map[string]struct{}, len(roles))
	for _, role := range roles {
		roleSet[role] = struct{}{}
	}
	return func(r *ghttp.Request) {
		userRole := gconv.String(r.GetCtxVar(consts.CtxUserRole))
		if _, ok := roleSet[userRole]; !ok {
			r.Response.WriteStatusExit(403, g.Map{"message": "Forbidden: insufficient permission"})
			return
		}
		r.Middleware.Next()
	}
}
