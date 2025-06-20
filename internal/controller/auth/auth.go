package auth

import "medicare/api/v1/auth"

var Auth = LAuth{}

type LAuth struct{}

func NewV1() auth.AuthV1 {
	return &LAuth{}
}
