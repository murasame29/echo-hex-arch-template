package server

import "context"

type Server interface {
	Run() error
	Close(context.Context) error
}
