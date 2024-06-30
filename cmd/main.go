package main

import (
	"log"

	"go/github.com/uh-zz/gocon2024/gen"
	"go/github.com/uh-zz/gocon2024/internal/server"

	"github.com/labstack/echo/v4"
)

func main() {
	// create a type that satisfies the `api.ServerInterface`, which contains an implementation of every operation from the generated code
	s := server.NewServer()

	e := echo.New()
	e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return h(server.NewOriginalContext(c))
		}
	})

	gen.RegisterHandlers(e, s)

	// And we serve HTTP until the world ends.
	log.Fatal(e.Start("0.0.0.0:8080"))
}
