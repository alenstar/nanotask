package server

import (
	"net/http"
)

type Server struct {
	addr string
	mux  *http.ServeMux
}

func New(address string) *Server {
	return &Server{addr: address}
}
func (s *Server) Handle(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	if s.mux == nil {
		s.mux = http.NewServeMux()
	}
	s.mux.HandleFunc(pattern, handler)
}

func (s *Server) Run() {
	http.ListenAndServe(s.addr, s.mux)
}

func (s *Server) Shutdown() {
	// TODO
}
