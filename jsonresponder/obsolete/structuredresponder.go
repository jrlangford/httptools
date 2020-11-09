package jsonresponder

import (
	"encoding/json"
	"net/http"
)

//ResponseStatus, includes attributes that every response should include
type Detail struct {
	DebugMessage string `json:"error,omitempty"`
	Code string `json:"code,omitempty"`
}

type ResponseStatus struct {
	Success      int    `json:"success"`
    Details      []Detail `json:"details,omitempty"`
}

type JSONStructuredResponder struct {
    *JSONResponder;
}

func NewJSONStructuredResponder(logger ErrorLogger) (r *JSONStructuredResponder) {
    r = &JSONStructuredResponder {
        JSONResponder: NewJSONResponder(logger),
    }
    return
}

func (r *JSONStructuredResponder) SendSuccess(w http.ResponseWriter, data map[string]interface{}) {
	response.setSuccess(1)
	Respond(w, response)
}

func (r *JSONStructuredResponder) SendError(w http.ResponseWriter, errCode string, responseCode int, response ServerResponse) {
	response.setSuccess(0)
	response.setCode(errCode)
	w.WriteHeader(responseCode)
	Respond(w, response)
}


