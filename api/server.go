package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/pi-rate14/transaction-module/db/sqlc"
)

// Server serves all HTTP request for the service

type Server struct {
	store db.Store
	router *gin.Engine
}

// NewServer returns an instance of a new HTTP server
func NewServer(store db.Store) *Server {
	server := &Server{
		store: store,
	}

	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	// add routes to router
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)

	router.POST("/transfers", server.createTransfer)


	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
 
func errorResponse(err error) gin.H {
	return gin.H{"error":err.Error()}
}