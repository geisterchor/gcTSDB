package restapi

import (
	"geisterchor.com/gcTSDB/gctsdb"

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

type Channel struct {
	Name       string `json:"name"`
	DataType   string `json:"datatype"`
	BucketSize int64  `json:"bucketSize"`
}

func NewRestChannel(ch *gctsdb.Channel) (*Channel, error) {
	lookup := map[string]string{
		"i32": "int32",
		"i64": "int64",
		"f32": "float32",
		"f64": "float64",
		"dec": "decimal",
		"str": "string",
	}

	datatype, ok := lookup[ch.DataType]
	if !ok {
		return nil, fmt.Errorf("Channel %s has an unknown data type (%s)", ch.Name, ch.DataType)
	}

	rch := Channel{
		Name:       ch.Name,
		DataType:   datatype,
		BucketSize: int64(*ch.BucketSize),
	}

	return &rch, nil
}
