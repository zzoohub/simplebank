package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/zzoopro/simple_bank/db/sqlc"
)

// Server serves HTTP requests for our banking services.
type Server struct {
	store db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing.
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	if v, ok :=binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrancy)
	}


	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.DELETE("/accounts/:id", server.deleteAccount)
	router.GET("/accounts", server.getAccountList)
	router.POST("/transfers", server.createTransfer)

	server.router = router
	return server 
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