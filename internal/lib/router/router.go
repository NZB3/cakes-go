package router

import (
	"github.com/gorilla/mux"
	"net/http"
)

type router struct {
	mux *mux.Router
}

func NewRouter() *router {
	r := mux.NewRouter()
	apiRouter := r.PathPrefix("/api/v1/").Subrouter() // Create a subrouter with the prefix
	return &router{
		mux: apiRouter,
	}
}

func (r *router) Group(path string) *router {
	return &router{
		mux: r.mux.PathPrefix(path).Subrouter(),
	}
}

func (r *router) Mount(path string, handler http.HandlerFunc) {
	r.mux.HandleFunc(path, handler)
}

func (r *router) MountWithMethod(path string, method string, handler http.HandlerFunc) {
	r.mux.HandleFunc(path, handler).Methods(method)
}

func (r *router) GetHandler() http.Handler {
	return r.mux
}
