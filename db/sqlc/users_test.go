package db

import (
	"context"
	"testing"
	"time"

	"github.com/BinayRajbanshi/GoBasicBank/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	hashedPassword, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)
	args := CreateUserParams{
		Username:       util.RandOwner(),
		HashedPassword: hashedPassword,
		FullName:       util.RandOwner(),
		Email:          util.RandomEmail(),
	}

	user, err := testStore.CreateUser(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, args.Username, user.Username)
	require.Equal(t, args.FullName, user.FullName)
	require.Equal(t, args.HashedPassword, user.HashedPassword)
	require.Equal(t, user.Email, args.Email)
	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	randomUser := createRandomUser(t)

	user, err := testStore.GetUser(context.Background(), randomUser.Username)

	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, randomUser.Email, user.Email)
	require.Equal(t, randomUser.FullName, user.FullName)
	require.Equal(t, randomUser.HashedPassword, user.HashedPassword)
	require.WithinDuration(t, randomUser.PasswordChangedAt, user.PasswordChangedAt, time.Second)
	require.WithinDuration(t, randomUser.CreatedAt, user.CreatedAt, time.Second)
}
