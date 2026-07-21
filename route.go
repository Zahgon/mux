package mux

import (
	"net/http"
	"net/url"
	"regexp"
)

type Route struct {
	handler http.Handler

	buildOnly bool

	name string

	err error

	metadata map[any]any

	namedRoutes map[string]*Route

	middlewares []middleware

	routeConf
}

func (r *Route) SkipClean() bool { _ = "STUB: not implemented"; return false }

func (r *Route) Match(req *http.Request, match *RouteMatch) bool {
	_ = "STUB: not implemented"
	return false
}

func (r *Route) GetError() error { _ = "STUB: not implemented"; return nil }

func (r *Route) BuildOnly() *Route { _ = "STUB: not implemented"; return nil }

func (r *Route) Metadata(key any, value any) *Route { _ = "STUB: not implemented"; return nil }

func (r *Route) GetMetadata() map[any]any { _ = "STUB: not implemented"; return nil }

func (r *Route) MetadataContains(key any) bool { _ = "STUB: not implemented"; return false }

func (r *Route) GetMetadataValue(key any) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (r *Route) GetMetadataValueOr(key any, fallbackValue any) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (r *Route) Handler(handler http.Handler) *Route { _ = "STUB: not implemented"; return nil }

func (r *Route) HandlerFunc(f func(http.ResponseWriter, *http.Request)) *Route {
	_ = "STUB: not implemented"
	return nil
}

func (r *Route) GetHandler() http.Handler { _ = "STUB: not implemented"; return *new(http.Handler) }

func (r *Route) GetHandlerWithMiddlewares() http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

func (r *Route) Name(name string) *Route { _ = "STUB: not implemented"; return nil }

func (r *Route) GetName() string { _ = "STUB: not implemented"; return "" }

type matcher interface {
	Match(*http.Request, *RouteMatch) bool
}

func (r *Route) addMatcher(m matcher) *Route { _ = "STUB: not implemented"; return nil }

func (r *Route) addRegexpMatcher(tpl string, typ regexpType) error {
	_ = "STUB: not implemented"
	return nil
}

type headerMatcher map[string]string

func (m headerMatcher) Match(r *http.Request, match *RouteMatch) bool {
	_ = "STUB: not implemented"
	return false
}

func (r *Route) Headers(pairs ...string) *Route { _ = "STUB: not implemented"; return nil }

type headerRegexMatcher map[string]*regexp.Regexp

func (m headerRegexMatcher) Match(r *http.Request, match *RouteMatch) bool {
	_ = "STUB: not implemented"
	return false
}

func (r *Route) HeadersRegexp(pairs ...string) *Route { _ = "STUB: not implemented"; return nil }

func (r *Route) Host(tpl string) *Route { _ = "STUB: not implemented"; return nil }

type MatcherFunc func(*http.Request, *RouteMatch) bool

func (m MatcherFunc) Match(r *http.Request, match *RouteMatch) bool {
	_ = "STUB: not implemented"
	return false
}

func (r *Route) MatcherFunc(f MatcherFunc) *Route { _ = "STUB: not implemented"; return nil }

type methodMatcher []string

func (m methodMatcher) Match(r *http.Request, match *RouteMatch) bool {
	_ = "STUB: not implemented"
	return false
}

func (r *Route) Methods(methods ...string) *Route { _ = "STUB: not implemented"; return nil }

func (r *Route) Path(tpl string) *Route { _ = "STUB: not implemented"; return nil }

func (r *Route) PathPrefix(tpl string) *Route { _ = "STUB: not implemented"; return nil }

func (r *Route) Queries(pairs ...string) *Route { _ = "STUB: not implemented"; return nil }

type schemeMatcher []string

func (m schemeMatcher) Match(r *http.Request, match *RouteMatch) bool {
	_ = "STUB: not implemented"
	return false
}

func (r *Route) Schemes(schemes ...string) *Route { _ = "STUB: not implemented"; return nil }

type BuildVarsFunc func(map[string]string) map[string]string

func (r *Route) BuildVarsFunc(f BuildVarsFunc) *Route { _ = "STUB: not implemented"; return nil }

func (r *Route) Subrouter() *Router { _ = "STUB: not implemented"; return nil }

func (r *Route) URL(pairs ...string) (*url.URL, error) { _ = "STUB: not implemented"; return nil, nil }

func (r *Route) URLHost(pairs ...string) (*url.URL, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Route) URLPath(pairs ...string) (*url.URL, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Route) GetPathTemplate() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (r *Route) GetPathRegexp() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (r *Route) GetQueriesRegexp() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func (r *Route) GetQueriesTemplates() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func (r *Route) GetMethods() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func (r *Route) GetHostTemplate() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (r *Route) GetVarNames() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func (r *Route) prepareVars(pairs ...string) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Route) buildVars(m map[string]string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}
