package server

import (
	"net/http"

	"github.com/praveenmahasena/engine/internal/server/handler"
	"github.com/spf13/viper"
)

type Server struct {
	ListenAddr string
}

const (
	portParams = "PORT"
)

func New(p ...string) *Server {
	s := &Server{}
	if len(p) > 0 {
		s.ListenAddr = p[0]
		return s
	}
	s.ListenAddr = viper.GetString(portParams)
	return s
}

func (s *Server) Run() error {
	m := http.NewServeMux()

	m.HandleFunc("/test", handler.Testing)

	return http.ListenAndServe(s.ListenAddr, m)
}
