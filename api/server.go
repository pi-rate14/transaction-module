package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/pi-rate14/transaction-module/db/sqlc"
)

// Server serves all HTTP request for the service

type Server struct {
	store *db.Store
	router *gin.Engine
}

// NewServer returns an instance of a new HTTP server
func NewServer(store *db.Store) *Server {
	server := &Server{
		store: store,
	}

	router := gin.Default()

	// add routes to router
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
 
func errorResponse(err error) gin.H {
	return gin.H{"error":err.Error()}
}