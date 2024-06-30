package server

import (
	"go/github.com/uh-zz/gocon2024/gen"
	"log"

	"github.com/labstack/echo/v4"
)

var _ gen.ServerInterface = (*Server)(nil)

type Server struct {
	m Middlewares
}

func NewServer() Server {
	m := Middlewares{
		"admin": []Middleware{
			func() {
				log.Printf("admin middleware")
			},
			func() {
				log.Printf("admin middleware 2")
			},
		},
		"normal": []Middleware{
			func() {
				log.Printf("normal middleware")
			},
		},
	}

	return Server{m: m}
}

func (s Server) GetAdminThings(ctx echo.Context) error {
	log.Printf("GetAdminThings")

	c := ctx.(*OriginalContext)
	if err := c.BindValidate(s.m, nil); err != nil {
		return err
	}
	return nil
}

func (s Server) GetThings(ctx echo.Context) error {
	log.Printf("GetThings")
	
	c := ctx.(*OriginalContext)
	if err := c.BindValidate(s.m, nil); err != nil {
		return err
	}
	return nil
}
