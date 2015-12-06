package restapi

import (
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/context"

	"net/http"
	"time"
)

type Logger struct {
}

func NewLogger() *Logger {
	return &Logger{}
}

func (lg *Logger) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	start := time.Now()

	remoteHost := r.RemoteAddr
	if len(r.Header["X-Forwarded-For"]) > 0 {
		remoteHost = r.Header["X-Forwarded-For"][0]
	}

	l := log.WithFields(log.Fields{
		"remoteHost": remoteHost,
		"timestamp":  start,
		"method":     r.Method,
		"host":       r.Host,
		"uri":        r.RequestURI})

	var y *Authentication
	if intf, ok := context.GetOk(r, "AUTHENTICATION"); ok {
		y = intf.(*Authentication)
	}

	if y != nil {
		l = l.WithFields(log.Fields{
			"gc-scopes": y.Scopes,
			"gc-userid": y.UserId,
		})
	}

	next(rw, r)

	res := rw.(negroni.ResponseWriter)
	responseTime := float32(time.Since(start)/time.Microsecond) / 1000

	l = l.WithFields(log.Fields{
		"status":         res.Status(),
		"statusMessage":  http.StatusText(res.Status()),
		"responseSize":   res.Size(),
		"responseTimeMs": responseTime,
	})

	l.Info()
}
