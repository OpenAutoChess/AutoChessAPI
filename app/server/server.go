package server

import (
    "io"
    "github.com/gorilla/mux"
    "net/http"
)

type Server struct{
    config *Config
    router *mux.Router
}

func New(config *Config) *Server {
    return &Server{
        config: config,
        router: mux.NewRouter(),
    }
}

func (s *Server) Start() error {
    s.configureRouter()

    return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *Server) configureRouter() {
    s.router.HandleFunc("/", s.handleHello())
}

func (s *Server) handleHello() http.HandlerFunc {

    return func(w http.ResponseWriter, r *http.Request) {
        io.WriteString(w, "Hello")
    }
}
