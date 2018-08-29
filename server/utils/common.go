package utils

import (
	"encoding/json"
	"net/http"
)

type commResp struct {
	Code    string `json:"conde"`
	Message string `json:"message"`
}

func Reply(w http.ResponseWriter, r *http.Request, status int, v interface{}) {
	data, _ := json.Marshal(v)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(data)
}

func CommReply(w http.ResponseWriter, r *http.Request, status int, msg string) {
	resp := commResp{
		Code:    http.StatusText(status),
		Message: msg,
	}
	Reply(w, r, status, resp)
}
