package server

import (
	"net/http"
	"time"

	"github.com/braintree/manners"
	"github.com/gorilla/mux"
	"github.com/showntop/circle-core/logger"
)

type ServerConf struct {
}

var Router *mux.Router

func Fire(config map[string]interface{}) {
	New()
	RegistHandlers()
	Start()
}

func New() {
	Router = mux.NewRouter()
	Router.NotFoundHandler = http.HandlerFunc(Handle404)
}

func Start() {
	// go func() {
	err := manners.ListenAndServe(":8000", Router)
	if err != nil {
		logger.Fatal("api.server.start_server.starting.critical")
		time.Sleep(time.Second)
		panic("api.server.start_server.starting.panic")
	}
	// }()
}

func RegistHandlers() {
	for _, route := range routes {
		var handler http.HandlerFunc
		handler = route.HandlerFunc
		Router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
}

func Handle404(w http.ResponseWriter, r *http.Request) {
	logger.Info("404")

}
