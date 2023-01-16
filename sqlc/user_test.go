package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"mina.com/blog/util"
)

type RegisterUserParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Nickname string `json:"nickname"`
}

func TestCreateUser(t *testing.T) {
	randomcreateUser(t)
}

func randomcreateUser(t *testing.T) User {
	registeredUserParams := RegisterUserParams{
		Username: util.RandomUsername(),
		Password: util.RandomPassword(),
		Email:    util.RandomEmail(),
		Nickname: util.RandomNickname(),
	}
	hashedPassword, err := util.HashPassword(registeredUserParams.Password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)
	createUserParams := CreateUserParams{
		Username:       registeredUserParams.Username,
		HashedPassword: hashedPassword,
		Email:          registeredUserParams.Email,
		Nickname:       registeredUserParams.Nickname,
	}
	user, err := testQueries.CreateUser(context.Background(), createUserParams)
	require.NoError(t, err)
	require.Equal(t, user.Username, registeredUserParams.Username)
	require.Equal(t, user.HashedPassword, hashedPassword)
	require.Equal(t, user.Email, registeredUserParams.Email)
	require.Equal(t, user.Nickname, registeredUserParams.Nickname)
	require.NotZero(t, user.CreateAt)
	require.Zero(t, user.PasswordChangedat)
	return user
}
