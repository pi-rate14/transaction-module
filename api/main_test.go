package api

import (
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	db "github.com/pi-rate14/transaction-module/db/sqlc"
	"github.com/pi-rate14/transaction-module/util"
	"github.com/stretchr/testify/require"
)

func NewTestServer(t *testing.T, store db.Store) *Server{
	config := util.Config{
		TokenSymmetricKey: util.RandomString(31),
		AcessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store)
	require.NoError(t, err)

	return server
}
 
func TestMain(m *testing.M) {

	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}