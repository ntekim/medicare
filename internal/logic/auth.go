package logic

import (
	"context"
	"errors"

	db "medicare/internal/dao/sqlc"
	"medicare/utility/helpers"
	"golang.org/x/crypto/bcrypt"
)

func Authenticate(ctx context.Context, email, password string) (*db.User, string, error) {
	user, err := Queries.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, "", errors.New("user not found")
	}

	// Compare hash
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return nil, "", errors.New("invalid password")
	}

	// Create JWT token
	token, err := helpers.GenerateToken(&user)
	if err != nil {
		return nil, "", err
	}

	return &user, token, nil
}


