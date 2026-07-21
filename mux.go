package mux

import (
	"errors"
	"net/http"
	"net/url"
	"regexp"
)

var (
	ErrMethodMismatch = errors.New("method is not allowed")

	ErrNotFound = errors.New("no matching route was found")

	RegexpCompileFunc = regexp.Compile

	ErrMetadataKeyNotFound = errors.New("key not found in metadata")
)

func NewRouter() *Router { _ = "STUB: not implemented"; return nil }

type Router struct {
	NotFoundHandler http.Handler

	MethodNotAllowedHandler http.Handler

	routes []*Route

	namedRoutes map[string]*Route

	KeepContext bool

	middlewares []middleware

	routeConf
}

type routeConf struct {
	useEncodedPath bool

	strictSlash bool

	skipClean bool

	omitRouteFromContext bool

	omitRouterFromContext bool

	regexp routeRegexpGroup

	matchers []matcher

	buildScheme string

	buildVarsFunc BuildVarsFunc
}

func copyRouteConf(r routeConf) routeConf { _ = "STUB: not implemented"; return *new(routeConf) }

func copyRouteRegexp(r *routeRegexp) *routeRegexp { _ = "STUB: not implemented"; return nil }

func (r *Router) Match(req *http.Request, match *RouteMatch) bool {
	_ = "STUB: not implemented"
	return false
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (r *Router) Get(name string) *Route { _ = "STUB: not implemented"; return nil }

func (r *Router) GetRoute(name string) *Route { _ = "STUB: not implemented"; return nil }

func (r *Router) StrictSlash(value bool) *Router { _ = "STUB: not implemented"; return nil }

func (r *Router) SkipClean(value bool) *Router { _ = "STUB: not implemented"; return nil }

func (r *Router) OmitRouteFromContext(value bool) *Router { _ = "STUB: not implemented"; return nil }

func (r *Router) OmitRouterFromContext(value bool) *Router { _ = "STUB: not implemented"; return nil }

func (r *Router) UseEncodedPath() *Router { _ = "STUB: not implemented"; return nil }

func (r *Router) NewRoute() *Route { _ = "STUB: not implemented"; return nil }

func (r *Router) Name(name string) *Route { _ = "STUB: not implemented"; return nil }

func (r *Router) Handle(path string, handler http.Handler) *Route {
	_ = "STUB: not implemented"
	return nil
}

func (r *Router) HandleFunc(path string, f func(http.ResponseWriter,
	*http.Request)) *Route {
	_ = "STUB: not implemented"
	return nil
}

func (r *Router) Headers(pairs ...string) *Route { _ = "STUB: not implemented"; return nil }

func (r *Router) Host(tpl string) *Route { _ = "STUB: not implemented"; return nil }

func (r *Router) MatcherFunc(f MatcherFunc) *Route { _ = "STUB: not implemented"; return nil }

func (r *Router) Methods(methods ...string) *Route { _ = "STUB: not implemented"; return nil }

func (r *Router) Path(tpl string) *Route { _ = "STUB: not implemented"; return nil }

func (r *Router) PathPrefix(tpl string) *Route { _ = "STUB: not implemented"; return nil }

func (r *Router) Queries(pairs ...string) *Route { _ = "STUB: not implemented"; return nil }

func (r *Router) Schemes(schemes ...string) *Route { _ = "STUB: not implemented"; return nil }

func (r *Router) BuildVarsFunc(f BuildVarsFunc) *Route { _ = "STUB: not implemented"; return nil }

func (r *Router) Walk(walkFn WalkFunc) error { _ = "STUB: not implemented"; return nil }

var SkipRouter = errors.New("skip this router")

type WalkFunc func(route *Route, router *Router, ancestors []*Route) error

func (r *Router) walk(walkFn WalkFunc, ancestors []*Route) error {
	_ = "STUB: not implemented"
	return nil
}

type RouteMatch struct {
	Route   *Route
	Handler http.Handler
	Vars    map[string]string

	MatchErr error
}

type contextKey int

const (
	varsKey contextKey = iota
	routeKey
	routerKey
)

func Vars(r *http.Request) map[string]string { _ = "STUB: not implemented"; return nil }

func CurrentRoute(r *http.Request) *Route { _ = "STUB: not implemented"; return nil }

func CurrentRouter(r *http.Request) *Router { _ = "STUB: not implemented"; return nil }

func requestWithVars(r *http.Request, vars map[string]string) *http.Request {
	_ = "STUB: not implemented"
	return nil
}

func requestWithRouteAndVars(r *http.Request, route *Route, vars map[string]string) *http.Request {
	_ = "STUB: not implemented"
	return nil
}

func requestWithRouter(r *http.Request, router *Router) *http.Request {
	_ = "STUB: not implemented"
	return nil
}

func cleanPath(p string) string { _ = "STUB: not implemented"; return "" }

func replaceURLPath(u *url.URL, p string) string { _ = "STUB: not implemented"; return "" }

func uniqueVars(s1, s2 []string) error { _ = "STUB: not implemented"; return nil }

func checkPairs(pairs ...string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func mapFromPairsToString(pairs ...string) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func mapFromPairsToRegex(pairs ...string) (map[string]*regexp.Regexp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func matchInArray(arr []string, value string) bool { _ = "STUB: not implemented"; return false }

func matchMapWithString(toCheck map[string]string, toMatch map[string][]string, canonicalKey bool) bool {
	_ = "STUB: not implemented"
	return false
}

func matchMapWithRegex(toCheck map[string]*regexp.Regexp, toMatch map[string][]string, canonicalKey bool) bool {
	_ = "STUB: not implemented"
	return false
}

func methodNotAllowed(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

func methodNotAllowedHandler() http.Handler { _ = "STUB: not implemented"; return *new(http.Handler) }
