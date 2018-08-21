package controller

import (
	"github.com/gorilla/mux"
)

type Controller interface {
	Register(router *mux.Router)
}
type Options struct {
}
