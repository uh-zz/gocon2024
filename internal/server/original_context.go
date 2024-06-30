package server

import (
	"go/github.com/uh-zz/gocon2024/gen"

	"github.com/labstack/echo/v4"
)

type OriginalContext struct {
	echo.Context
}

func NewOriginalContext(ctx echo.Context) *OriginalContext {
	return &OriginalContext{ctx}
}

func (c OriginalContext) BindValidate(m Middlewares, i interface{}) error {
	scopes, ok := c.Get(gen.RoleScopes).([]string)
	if !ok {
		scopes = []string{}
	}
	for _, scope := range scopes {
		if middleware, ok := m[scope]; ok {
			for _, mw := range middleware {
				mw()
			}
		}
	}

	if i == nil {
		return nil
	}

	if err := c.Bind(i); err != nil {
		return err
	}

	if err := c.Validate(i); err != nil {
		return err
	}
	return nil
}
