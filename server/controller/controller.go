package controller

import (
	"github.com/gorilla/mux"
	"reboot/server/service"
)

type Controller interface {
	Register(router *mux.Router)
}
type Options struct {
	Service service.Operation
}
