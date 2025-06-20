package auth

import "context"

type AuthV1 interface {
	Login(ctx context.Context, req *LoginReq) (res *LoginRes, err error)
}
