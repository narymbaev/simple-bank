package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/narymbaev/simple-bank/db/sqlc"
	"github.com/narymbaev/simple-bank/util"
	"os"
	"testing"
	"time"
)

func newServerTest(t *testing.T, store db.Store) *Server {
	config := util.Config{
		TokenSymmetricKey:   util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(store, config)
	if err != nil {
		return nil
	}

	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}