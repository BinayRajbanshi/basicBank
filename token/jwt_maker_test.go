package token

import (
	"testing"
	"time"

	"github.com/BinayRajbanshi/GoBasicBank/util"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/require"
)

func TestJwtMaker(t *testing.T) {
	maker, err := NewJWTMaker(util.RandomString(32))
	require.NoError(t, err)

	username := util.RandOwner()
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, err := maker.CreateToken(username, duration)

	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.NotZero(t, payload.ID)
	require.Equal(t, username, payload.Username)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}

func TestExpiredJWTToken(t *testing.T) {
	maker, err := NewJWTMaker(util.RandomString(32))
	require.NoError(t, err)

	token, err := maker.CreateToken(util.RandOwner(), -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrExpiredToken.Error())
	require.Nil(t, payload)
}

// TestInvalidJWTTokenAlgNone ensures that the JWT verification logic
// correctly rejects tokens signed with the "none" algorithm (i.e., unsigned tokens).
// This test helps prevent a critical security vulnerability where accepting such tokens
// could allow attackers to bypass authentication.
func TestInvalidJWTTokenAlgNone(t *testing.T) {
	payload, err := NewPayload(util.RandOwner(), time.Minute)
	require.NoError(t, err)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, payload)
	// UnsafeAllowNoneSignatureType should be only used in test and not in production
	token, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
	require.NoError(t, err)

	maker, err := NewJWTMaker(util.RandomString(32))
	require.NoError(t, err)

	payload, err = maker.VerifyToken(token)
	require.Error(t, err) // need error here
	require.EqualError(t, err, ErrInvalidToken.Error())
	require.Nil(t, payload)
}
