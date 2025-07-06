package server

import (
	"go.uber.org/fx"
	"harris.com/api"
)

var _ api.ServerInterface = (*Server)(nil)

type Server struct {
	fx.In
}
