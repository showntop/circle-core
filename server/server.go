package server

type Server struct {
}

var Router *mux.Router

func Load() {

}

func Start() {
	Router = mux.NewRouter()
	Router.NotFoundHandler = http.HandlerFunc(Handle404)
}
