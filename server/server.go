package server

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Options struct {
	ListenAddr string
}

type Server interface {
	http.Handler
	ListenAndServer() error
}

type server struct {
	opt    Options
	router *mux.Router
}

func New(opt Options) Server {
	r := mux.NewRouter()
	return &server{
		opt:    opt,
		router: r,
	}

}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

//ListenAndServer xxx
func (s *server) ListenAndServer() error {
	httpServer := &http.Server{
		Addr:    s.opt.ListenAddr,
		Handler: s.router,
	}
	return httpServer.ListenAndServe()
}
