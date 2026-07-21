package mux

import (
	"net/http"
)

type MiddlewareFunc func(http.Handler) http.Handler

type middleware interface {
	Middleware(handler http.Handler) http.Handler
}

func (mw MiddlewareFunc) Middleware(handler http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

func (r *Router) Use(mwf ...MiddlewareFunc) { _ = "STUB: not implemented"; return }

func (r *Router) useInterface(mw middleware) { _ = "STUB: not implemented"; return }

func (r *Route) Use(mwf ...MiddlewareFunc) *Route { _ = "STUB: not implemented"; return nil }

func (r *Route) useInterface(mw middleware) { _ = "STUB: not implemented"; return }

func CORSMethodMiddleware(r *Router) MiddlewareFunc {
	_ = "STUB: not implemented"
	return *new(MiddlewareFunc)
}

func getAllMethodsForRoute(r *Router, req *http.Request) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
