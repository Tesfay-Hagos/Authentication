package test

// import (
// 	"2fstockshow/internal/constant/state"
// 	"2fstockshow/platform/utils/util"

// 	"testing"
// 	"time"

// 	"github.com/stretchr/testify/require"
// )

// func TestPasetoMaker(t *testing.T) {
// 	maker, err := NewPasetoMaker(util.RandomString(32))
// 	require.NoError(t, err)

// 	username := util.RandomOwner()
// 	duration := time.Minute

// 	issuedAt := time.Now()
// 	expiredAt := issuedAt.Add(duration)

// 	token, err := maker.CreateToken(username, state.CustomerGrant, duration)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, token)

// 	payload, err := maker.VerifyToken(token)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, token)

// 	require.NotZero(t, payload.ID)
// 	require.Equal(t, username, payload.UserID)
// 	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
// 	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
// }
