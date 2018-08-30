package task

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/leopoldxx/go-utils/middleware"
	"github.com/leopoldxx/go-utils/trace"
	"io/ioutil"
	"net/http"
	"reboot/server/controller"
	"reboot/server/service"
	"reboot/server/utils"
)

type task struct {
	opt *controller.Options
}

// 客户端返回数据定义格式
/*
{
    "code":200,
    "msg":"suceess",
}
{
    "code":400,
    "msg":"error message",
}

*/

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
	err := t.opt.Service.GetTask(r.Context())
	if err != nil {
		tracer.Error(err)
		utils.CommReply(w, r, http.StatusBadRequest, err.Error())
		return
	}
	//fmt.Fprintln(w, "call getTask")
	utils.CommReply(w, r, http.StatusOK, "success")
}

func (t *task) listTask(w http.ResponseWriter, r *http.Request) {
	tracer := trace.GetTraceFromRequest(r)
	tracer.Info("call listTask")
	fmt.Fprintln(w, "call listTask")
}

func (t *task) createTask(w http.ResponseWriter, r *http.Request) {
	tracer := trace.GetTraceFromRequest(r)
	tracer.Info("call createTask")
	vars := mux.Vars(r)
	ns := vars["namespace"]
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		tracer.Error(err)
		utils.CommReply(w, r, http.StatusBadRequest, err.Error())
		return
	}

	info := &service.Task{}
	err = json.Unmarshal(data, info)
	if err != nil {
		tracer.Error(err)
		utils.CommReply(w, r, http.StatusBadRequest, err.Error())
		return
	}
	fmt.Fprintln(w, "call createTask")
	//Todo: validate  验证 ns  resource  是否合法
	t.opt.Service.CreateTask(r.Context(), ns, info.Resource)
	utils.CommReply(w, r, http.StatusAccepted, "Accept")
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
