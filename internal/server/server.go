package server

import (
	"go/github.com/uh-zz/gocon2024/gen"
	"log"

	"github.com/labstack/echo/v4"
)

var _ gen.ServerInterface = (*Server)(nil)

type Server struct{}

func NewServer() Server {
	return Server{}
}

func (s Server) GetAdminThings(ctx echo.Context) error {
	log.Printf("GetAdminThings")
	return nil
}

func (s Server) GetThings(ctx echo.Context) error {
	log.Printf("GetThings")
	return nil
}
