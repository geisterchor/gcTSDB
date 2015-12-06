package restapi

import (
	"geisterchor.com/gcTSDB/gctsdb"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"

	"net/http"
	"os"
)

func StartRestAPI(gctsdbServer *gctsdb.GCTSDBServer) {
	ctx := AppContext{
		GCTSDBServer: gctsdbServer,
	}

	r := mux.NewRouter()
	s := r.PathPrefix("/v1/").Subrouter()
	x := NewAppHandler(&ctx, s)

	registerHandlers(&x)

	s.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	n := negroni.New()
	n.Use(negroni.NewRecovery())
	n.Use(NewLogger())
	n.Use(NewCORS())
	n.UseHandler(r)

	listen := "127.0.0.1:3000"
	if len(os.Getenv("LISTEN")) > 0 {
		listen = os.Getenv("LISTEN")
	}

	log.Printf("API is available at http://%s/v1/", listen)
	n.Run(listen)
}

const (
	PUT    = "PUT"
	GET    = "GET"
	POST   = "POST"
	DELETE = "DELETE"
)

func registerHandlers(x *AppHandler) {
	/*
		x.Handle(POST, "/account/login", resources.LoginHandler)
		x.Handle(POST, "/account/logout", resources.LogoutHandler, "AUTHENTICATED")
		x.Handle(POST, "/account/passwordreset", resources.PasswordResetInitiateHandler)
		x.Handle(PUT, "/account/passwordreset", resources.PasswordResetExecuteHandler)
	*/
}
