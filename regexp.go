package mux

import (
	"net/http"
	"regexp"
)

type routeRegexpOptions struct {
	strictSlash    bool
	useEncodedPath bool
}

type regexpType int

const (
	regexpTypePath regexpType = iota
	regexpTypeHost
	regexpTypePrefix
	regexpTypeQuery
)

func newRouteRegexp(tpl string, typ regexpType, options routeRegexpOptions) (*routeRegexp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type routeRegexp struct {
	template string

	regexpType regexpType

	options routeRegexpOptions

	regexp *regexp.Regexp

	reverse string

	varsN []string

	varsR []*regexp.Regexp

	wildcardHostPort bool
}

func (r *routeRegexp) Match(req *http.Request, match *RouteMatch) bool {
	_ = "STUB: not implemented"
	return false
}

func (r *routeRegexp) url(values map[string]string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (r *routeRegexp) getURLQuery(req *http.Request) string { _ = "STUB: not implemented"; return "" }

func findFirstQueryKey(rawQuery, key string) (value string, ok bool) {
	_ = "STUB: not implemented"
	return "", false
}

func (r *routeRegexp) matchQueryString(req *http.Request) bool {
	_ = "STUB: not implemented"
	return false
}

func braceIndices(s string) ([]int, error) { _ = "STUB: not implemented"; return nil, nil }

func varGroupName(idx int) string { _ = "STUB: not implemented"; return "" }

type routeRegexpGroup struct {
	host    *routeRegexp
	path    *routeRegexp
	queries []*routeRegexp
}

func (v routeRegexpGroup) setMatch(req *http.Request, m *RouteMatch, r *Route) {
	_ = "STUB: not implemented"
	return
}

func getHost(r *http.Request) string { _ = "STUB: not implemented"; return "" }

func extractVars(input string, matches []int, names []string, output map[string]string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}
