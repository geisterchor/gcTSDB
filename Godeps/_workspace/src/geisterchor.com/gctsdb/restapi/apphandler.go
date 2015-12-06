package restapi

import (
	"github.com/gorilla/context"
	"github.com/gorilla/mux"

	"net/http"
)

type Authentication struct {
	Scopes []string
	Token  string
	UserId string
}

func (y *Authentication) HasScope(scope string) bool {
	if y == nil || y.Scopes == nil {
		return false
	}
	for _, b := range y.Scopes {
		if b == scope {
			return true
		}
	}
	return false
}

func GetAuthentication(ctx *AppContext, r *http.Request) *Authentication {
	return nil
}

type AppHandler struct {
	router  *mux.Router
	Context *AppContext
}

func NewAppHandler(appContext *AppContext, r *mux.Router) AppHandler {
	return AppHandler{router: r, Context: appContext}
}

func optionsHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func (h *AppHandler) Handle(method string, uri string, handler func(*AppContext, http.ResponseWriter, *http.Request), scopes ...string) {
	authHandler := UberAuthorizator{
		appHandler: h,
		method:     method,
		uri:        uri,
		scopes:     scopes,
		endHandler: handler,
	}

	h.router.HandleFunc(uri, authHandler.Handler).Methods(method)

	// CORS
	h.router.HandleFunc(uri, optionsHandler).Methods("OPTIONS")
}

type UberAuthorizator struct {
	appHandler *AppHandler
	method     string
	uri        string
	scopes     []string
	endHandler func(*AppContext, http.ResponseWriter, *http.Request)
}

func (ua *UberAuthorizator) Handler(rw http.ResponseWriter, r *http.Request) {
	y := GetAuthentication(ua.appHandler.Context, r)

	if len(ua.scopes) > 0 && y == nil {
		NewHttpError(401, "NOT_AUTHENTICATED", "Invalid or no authorization token.").WriteResponse(rw)
		return
	}

	for _, requiredScope := range ua.scopes {
		if !y.HasScope(requiredScope) {
			NewHttpError(403, "NOT_ALLOWED", "You do not have permission to access this resource.").WriteResponse(rw)
			return
		}
	}

	context.Set(r, "AUTHENTICATION", y)

	ua.endHandler(ua.appHandler.Context, rw, r)
}
