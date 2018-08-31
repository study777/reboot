package server

import (
	"github.com/gorilla/mux"
	"net/http"
	"reboot/server/controller"
	"reboot/server/controller/task"
	"reboot/server/service"
)

type Options struct {
	ListenAddr string
	CtrlOpts   *controller.Options
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
	//add router
	v1Router := r.PathPrefix("/reboot/api/v1").Subrouter()
	//v2Router := r.PathPrefix("/reboot/api/v2").Subrouter()
	task.New(opt.CtrlOpts).Register(v1Router)
	opt.CtrlOpts.Service = service.New(&service.Options{
		DB: opt.CtrlOpts.DB,
	})
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
