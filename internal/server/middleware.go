package server

type Middleware func()

type Middlewares map[string][]Middleware
