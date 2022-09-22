package server

import (
	"context"
	"net"

	"net/http"
)

type (
	Server struct {
		Server *http.Server
	}
)

func NewServer(handler http.HandlerFunc) (s *Server) {
	s = &Server{
		Server: &http.Server{
			Handler:         handler,
		},
	}

	return
}

func (s *Server) Serve(l net.Listener) (err error) {
	err = s.Server.Serve(l)
	return
}

func (s *Server) Close() (err error) {
	err = s.Server.Shutdown(context.Background())
	return
}
