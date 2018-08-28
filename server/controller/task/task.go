package task

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/leopoldxx/go-utils/middleware"
	"github.com/leopoldxx/go-utils/trace"
	"net/http"
	"reboot/server/controller"
	"reboot/server/utils"
)

type task struct {
	opt *controller.Options
}

func New(opt *controller.Options) controller.Controller {
	return &task{opt: opt}
}

func (t *task) Register(router *mux.Router) {
	//localhost:7878/reboot/api/v1/namespace/{xxx}/tasks/{task}
	subrouter := router.PathPrefix("/namespaces/{namespace}").Subrouter()
	subrouter.Use(utils.LoggingMiddleware)
	//subrouter.Methods("GET").Path("/tasks/{task}").HandlerFunc(t.getTask)

	//为了测试 loggingMiddleware
	//subrouter.Methods("GET").Path("/tasks/{task}").HandlerFunc(middleware.RecoverWithTrace("getTask").HandlerFunc(
	//	utils.AuthenticateMW().HandlerFunc(t.getTask),
	//),
	//)

	subrouter.Methods("GET").Path("/tasks/{task}").HandlerFunc(
		middleware.RecoverWithTrace("getTask").HandlerFunc(t.getTask),
	)

	subrouter.Methods("GET").Path("/tasks").HandlerFunc(middleware.RecoverWithTrace("listTask").HandlerFunc(t.listTask))
	//subrouter.Methods("GET").Path("/tasks").HandlerFunc(t.listTask)
	subrouter.Methods("POST").Path("/tasks").HandlerFunc(middleware.RecoverWithTrace("createTask").HandlerFunc(t.createTask))
	//subrouter.Methods("POST").Path("/tasks").HandlerFunc(t.createTask)
	subrouter.Methods("DELETE").Path("/tasks/{task}").HandlerFunc(middleware.RecoverWithTrace("deleteTask").HandlerFunc(t.deleteTask))
	//subrouter.Methods("DELETE").Path("/tasks/{task}").HandlerFunc(t.deleteTask)
	subrouter.Methods("PUT").Path("/tasks/{task}").HandlerFunc(middleware.RecoverWithTrace("updateTask").HandlerFunc(t.updateTask))
	//subrouter.Methods("PUT").Path("/tasks/{task}").HandlerFunc(t.updateTask)

}

func (t *task) getTask(w http.ResponseWriter, r *http.Request) {
	tracer := trace.GetTraceFromRequest(r)
	tracer.Info("call geTask")
	fmt.Fprintln(w, "call getTask")
}

func (t *task) listTask(w http.ResponseWriter, r *http.Request) {
	tracer := trace.GetTraceFromRequest(r)
	tracer.Info("call listTask")
	fmt.Fprintln(w, "call listTask")
}

func (t *task) createTask(w http.ResponseWriter, r *http.Request) {
	tracer := trace.GetTraceFromRequest(r)
	tracer.Info("call createTask")
	fmt.Fprintln(w, "call createTask")
}

func (t *task) deleteTask(w http.ResponseWriter, r *http.Request) {
	tracer := trace.GetTraceFromRequest(r)
	tracer.Info("call deleteTask")
	fmt.Fprintln(w, "call deleteTask")
}

func (t *task) updateTask(w http.ResponseWriter, r *http.Request) {
	tracer := trace.GetTraceFromRequest(r)
	tracer.Info("call updateTask")
	fmt.Fprintln(w, "call updateTask")
}
