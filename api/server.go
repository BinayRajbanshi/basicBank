package api

import (
	db "github.com/BinayRajbanshi/GoBasicBank/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// server serves HTTP request for our application
type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	router := gin.Default()
	server := &Server{store: store, router: router}

	// adding custom validator
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	//user routes
	server.router.POST("/api/v1/users", server.createUser)

	// account routes
	server.router.POST("/api/v1/accounts", server.createAccount)
	server.router.GET("/api/v1/accounts/:id", server.getAccount)
	server.router.GET("/api/v1/accounts", server.getAccounts)

	// transfer routes
	server.router.POST("/api/v1/transfers", server.createTransfer)

	return server
}

// start the server
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

// global error responder. gin.h is nothing but string key  andy value pair data structure
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
