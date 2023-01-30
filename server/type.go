package server

import (
	"net/http"
	"time"
)

type ServerOptions struct {
	Handler      http.Handler
	WriteTimeout time.Duration
	ReadTimeout  time.Duration
	IdleTimeout  time.Duration

	Host string
	Port string
}

type IServer interface {
	Run()
}
