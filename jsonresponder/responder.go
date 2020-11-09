package jsonresponder

import (
	"encoding/json"
	"net/http"
)

type ErrorLogger interface {
    Errorf(format string, args ...interface{})
}

type JSONResponder struct {
    log ErrorLogger;
}

func (r *JSONResponder) Send(w http.ResponseWriter, responseCode int, data interface{}){
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(responseCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		r.log.Errorf("Response send error: %+v", err)
	}
}

func New(logger ErrorLogger) (r *JSONResponder) {
    r = &JSONResponder {
        log: logger,
    }
    return
}
