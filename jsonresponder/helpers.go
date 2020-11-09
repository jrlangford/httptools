// Visit https://golang.org/src/net/http/status.go for a full list of status codes

package jsonresponder

import (
	"net/http"
)

// Represents a response's details
type Detail struct {
	DebugMessage string `json:"msg,omitempty"`
	Code string `json:"code,omitempty"`
}

// This struct can be embedded in response objects to provide additional response details
type ResponseStatus struct {
	Success      int    `json:"success"`
    Details      []Detail `json:"details,omitempty"`
}

func NewStatusOK() *ResponseStatus {
	return &ResponseStatus{
		Success: 1,
	}
}

func NewStatusNotOK() *ResponseStatus {
	return &ResponseStatus{
		Success: 0,
	}
}

func (s *ResponseStatus) PushDetails(debugMessage, code string){
	s.Details = append(
		s.Details,
		Detail {
			DebugMessage: debugMessage,
			Code: code,
		},
	)
}

func (r *JSONResponder)SendOK(w http.ResponseWriter, data interface{}){
    r.Send(w, http.StatusOK, data)
}

func (r *JSONResponder)SendBadRequest(w http.ResponseWriter, data interface{}){
    r.Send(w, http.StatusBadRequest, data)
}

func (r *JSONResponder)SendInternalServerError(w http.ResponseWriter, data interface{}){
    r.Send(w, http.StatusInternalServerError, data)
}
