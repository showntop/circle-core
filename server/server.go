package server

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ServerConf struct {
}

type Server struct {
	////其它信息
	////addr
	Router      *httprouter.Router
	middlewares []Middleware
}

func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	////执行middleware
	// logger.Info(req.RequestURI)
	// logger.Info(req.URL)
	fmt.Println(req.URL)
	fmt.Println(req.RequestURI)
	s.Router.ServeHTTP(w, req)
}

func (c *Server) use(middleware ...Middleware) {
	c.middlewares = append(c.middlewares, middleware...)
}

func (c *Server) AddRoute(method string, path string, handler http.HandlerFunc) {
	c.Router.HandlerFunc(method, path, handler)
}

// 给指定的http方法和路径注册handler
func (c *Server) Handle(method, path string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) {
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}

	c.Handle(method, path, handler)
}

func Handle404(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	fmt.Println(".........................not found.............................")

}

func New() *Server {
	r := httprouter.New()
	return &Server{Router: r}
}

func Fire(config map[string]interface{}) {
	fmt.Println("server is starting.......")

	c := New()

	c.Router.NotFound = http.HandlerFunc(Handle404)
	RegistRoutes(c)
	http.ListenAndServe(":8080", c)
}

func RegistRoutes(c *Server) {
	///根据routes的path,server config里面应该有的
	for _, route := range routes {
		var handler http.HandlerFunc
		handler = route.HandlerFunc
		c.AddRoute(route.Method, route.Pattern, handler)
	}
}
