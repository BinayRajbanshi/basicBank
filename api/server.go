package api

import (
	db "github.com/BinayRajbanshi/GoBasicBank/db/sqlc"
	"github.com/gin-gonic/gin"
)

// server serves HTTP request for our application
type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	router := gin.Default()
	server := &Server{store: store, router: router}

	// routes
	server.router.POST("/api/v1/accounts", server.createAccount)
	server.router.GET("/api/v1/accounts/:id", server.getAccount)
	server.router.GET("/api/v1/accounts", server.getAccounts)

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
