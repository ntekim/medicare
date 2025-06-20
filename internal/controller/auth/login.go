package auth

import (
	"context"
	"fmt"
	auth "medicare/api/v1/auth"
	"medicare/internal/logic"
	"net/http"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

func (c *LAuth) Login(ctx context.Context, req *auth.LoginReq) (res *auth.LoginRes, err error) {
	user, token, authError := logic.Authenticate(ctx, req.Email, req.Password)
	
	if err != nil {
		err = authError
		return nil, gerror.NewCode(gcode.New(http.StatusUnauthorized, "invalid email or password", err.Error()))
	}


	fullname := fmt.Sprintf("%s %s", user.FirstName, user.LastName)
	res = &auth.LoginRes{
		UserID:   user.ID.String(),
		FullName: fullname,
		Role:     string(user.Role),
		Token:    token,
	}
	return res, nil
}
