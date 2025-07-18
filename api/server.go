package api

import (
	"fmt"

	db "github.com/BinayRajbanshi/GoBasicBank/db/sqlc"
	"github.com/BinayRajbanshi/GoBasicBank/token"
	"github.com/BinayRajbanshi/GoBasicBank/util"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// server serves HTTP request for our application
type Server struct {
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
	config     util.Config
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	router := gin.Default()
	tokenMaker, err := token.NewJWTMaker(config.TokenSymmetrcKey)
	if err != nil {
		return nil, fmt.Errorf("cannot initialize token maker : %w", err)
	}
	server := &Server{store: store, tokenMaker: tokenMaker, config: config, router: router}

	// adding custom validator
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	//user routes
	server.router.POST("/api/v1/users", server.createUser)
	server.router.POST("/api/v1/users/login", server.loginUser)

	// account routes
	// server.router.POST("/api/v1/accounts", server.createAccount)
	server.router.GET("/api/v1/accounts/:id", server.getAccount)
	server.router.GET("/api/v1/accounts", server.getAccounts)

	// transfer routes
	server.router.POST("/api/v1/transfers", server.createTransfer)

	// PROTECTED ROUTES
	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
	authRoutes.POST("/api/v1/accounts", server.createAccount)

	return server, nil
}

// start the server
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

// global error responder. gin.h is nothing but string key  andy value pair data structure
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
