package auth

import "github.com/gogf/gf/v2/frame/g"

type LoginReq struct {
	g.Meta   `path:"/login" method:"post" tags:"Auth" summary:"Login for doctors or receptionists"`
	Email    string `json:"email"     v:"required|email" example:"john.doe@hospital.com"`
	Password string `json:"password"  v:"required|min-length:6" example:"password123"`
}

type LoginRes struct {
	UserID   string `json:"user_id"`
	FullName string `json:"full_name"`
	Role     string `json:"role"`  // doctor or receptionist
	Token    string `json:"token"` // e.g. JWT
}
