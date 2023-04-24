package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/zzoopro/simplebank/db/sqlc"
	"github.com/zzoopro/simplebank/token"
	"github.com/zzoopro/simplebank/util"
)

// Server serves HTTP requests for our banking services.
type Server struct {
	config 			util.Config
	store		 	db.Store
	tokenMaker 		token.Maker
	router 			*gin.Engine
}

// NewServer creates a new HTTP server and setup routing.
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config: config,
		store: store, 
		tokenMaker: tokenMaker,
	}	

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrancy)
	}

	server.setUpRouter()	
	return server, nil
}

func (server *Server) setUpRouter() {
	router := gin.Default()

	router.POST("/users", server.createUser)
	router.POST("/login", server.loginUser)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	authRoutes.POST("/accounts", server.createAccount)
	authRoutes.GET("/accounts/:id", server.getAccount)
	authRoutes.DELETE("/accounts/:id", server.deleteAccount)
	authRoutes.GET("/accounts", server.getAccountList)
	authRoutes.POST("/transfers", server.createTransfer)

	server.router = router
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
func okResponse(message string) gin.H {
	return gin.H{"ok": message}
}