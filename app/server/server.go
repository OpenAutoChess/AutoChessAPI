package server

import (
    "github.com/gorilla/mux"
    "net/http"
    // "github.com/googollee/go-socket.io"
    "encoding/json"
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
    type Move struct {
        Name string
    }
    return func(w http.ResponseWriter, r *http.Request) {

        resp := Move{"Nurik"}

        js, err := json.Marshal(resp)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.Header().Set("Content-Type", "application/json")
        w.Write(js)

    }
}
