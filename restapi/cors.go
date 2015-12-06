package restapi

import (
	"net/http"
)

type CORS struct {
	allowedOrigins map[string]bool
}

func NewCORS() *CORS {
	return &CORS{
		allowedOrigins: map[string]bool{
			"http://localhost:8000": true,
		},
	}
}

func (cors *CORS) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	clientOrigin := r.Header["Origin"]

	if len(clientOrigin) == 1 {
		if cors.allowedOrigins[clientOrigin[0]] == true {
			rw.Header().Set("Access-Control-Allow-Origin", clientOrigin[0])
			rw.Header().Set("Access-Control-Allow-Methods", "GET, PUT, POST, DELETE")
			rw.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-Auth-Token")
		}
	}

	next(rw, r)
}
