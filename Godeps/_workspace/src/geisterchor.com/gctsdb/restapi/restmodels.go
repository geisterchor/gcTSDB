package restapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Created struct {
	Id   string `json:"id"`
	Link string `json:"link"`
}

type HttpError struct {
	HTTPCode  int    `json:"-"`
	ErrorCode string `json:"errorCode"`
	Message   string `json:"errorMessage"`
}

func NewHttpError(httpCode int, errorCode string, errorMessage string, a ...interface{}) *HttpError {
	err := HttpError{}
	err.HTTPCode = httpCode
	err.ErrorCode = errorCode
	err.Message = fmt.Sprintf(errorMessage, a...)
	return &err
}

func (e *HttpError) WriteResponse(w http.ResponseWriter) {
	w.WriteHeader(e.HTTPCode)
	resp, _ := json.Marshal(e)
	w.Write(resp)
}
