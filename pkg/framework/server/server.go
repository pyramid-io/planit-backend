package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/pyramid.io/planit-backend/pkg/framework/http/router"
)

type ServerInterface interface {
	Start(port string)
}

type Server struct {
	Router *router.RouterInterface
}

func New(router router.RouterInterface) (ServerInterface, error) {
	return &Server{
		Router: &router,
	}, nil
}

func (s *Server) Start(port string) {
	fmt.Fprintf(os.Stdout, "Starting server on port %s\n", []any{port}...)

	if err := http.ListenAndServe(":"+port, *s.Router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}