package logic_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"

	db "medicare/internal/dao/sqlc"
	"medicare/internal/logic"
	mocksqlc "medicare/internal/mocks"
)

// mock token generator (optional: you can use real token gen too)
var dummyToken = "mocked.jwt.token"

func TestAuthenticate_Success(t *testing.T) {
	ctx := context.Background()

	password := "validpassword"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	mockQ := new(mocksqlc.Querier)
	logic.Queries = mockQ

	mockUser := db.User{
		Email:        "jane@example.com",
		PasswordHash: string(hashedPassword),
		Role:         "receptionist",
		FirstName:    "Jane",
		LastName:     "Doe",
	}

	// Mock the GetUserByEmail function
	mockQ.On("GetUserByEmail", mock.Anything, "jane@example.com").Return(mockUser, nil)

	user, token, err := logic.Authenticate(ctx, "jane@example.com", password)

	assert.NoError(t, err)
	assert.Equal(t, "jane@example.com", user.Email)
	assert.NotEmpty(t, token)
	mockQ.AssertExpectations(t)
}

func TestAuthenticate_UserNotFound(t *testing.T) {
	ctx := context.Background()

	mockQ := new(mocksqlc.Querier)
	logic.Queries = mockQ

	mockQ.On("GetUserByEmail", mock.Anything, "ghost@example.com").Return(db.User{}, errors.New("not found"))

	user, token, err := logic.Authenticate(ctx, "ghost@example.com", "somepass")

	assert.Error(t, err)
	assert.Nil(t, user)
	assert.Empty(t, token)
	assert.Equal(t, "user not found", err.Error())
	mockQ.AssertExpectations(t)
}

func TestAuthenticate_InvalidPassword(t *testing.T) {
	ctx := context.Background()

	correctPassword := "correctpass"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(correctPassword), bcrypt.DefaultCost)

	mockQ := new(mocksqlc.Querier)
	logic.Queries = mockQ

	mockUser := db.User{
		Email:        "jane@example.com",
		PasswordHash: string(hashedPassword),
		Role:         "doctor",
		FirstName:    "Jane",
		LastName:     "Doe",
	}

	mockQ.On("GetUserByEmail", mock.Anything, "jane@example.com").Return(mockUser, nil)

	user, token, err := logic.Authenticate(ctx, "jane@example.com", "wrongpass")

	assert.Error(t, err)
	assert.Nil(t, user)
	assert.Empty(t, token)
	assert.Equal(t, "invalid password", err.Error())
	mockQ.AssertExpectations(t)
}
