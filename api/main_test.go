package api

import (
	"os"
	"testing"
	"time"

	db "github.com/BinayRajbanshi/GoBasicBank/db/sqlc"
	"github.com/BinayRajbanshi/GoBasicBank/util"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	config := util.Config{
		TokenSymmetrcKey:    util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store)
	require.NoError(t, err)

	return server
}

// special function that's run before any tests in the package.
func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode) //Puts the Gin web framework in test mode, disabling unnecessary logging.

	os.Exit(m.Run()) //Runs all the tests and exits with the status code returned by m.Run().
}
