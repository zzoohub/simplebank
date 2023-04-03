package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/zzoopro/simple_bank/db/sqlc"
)

// Server serves HTTP requests for our banking services.
type Server struct {
	store *db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing.
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.DELETE("/accounts/:id", server.deleteAccount)
	router.GET("/accounts", server.getAccountList)

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