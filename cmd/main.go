package main

import (
	"log"

	"go/github.com/uh-zz/gocon2024/gen"
	"go/github.com/uh-zz/gocon2024/internal/server"

	"github.com/labstack/echo/v4"
)

func main() {
	// create a type that satisfies the `api.ServerInterface`, which contains an implementation of every operation from the generated code
	server := server.NewServer()

	e := echo.New()

	gen.RegisterHandlers(e, server)

	// And we serve HTTP until the world ends.
	log.Fatal(e.Start("0.0.0.0:8080"))
}
