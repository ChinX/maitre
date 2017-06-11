package svrmux

import (
	"net/http"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

var serialNum = map[string]string{
	http.MethodPost:   "01",
	http.MethodPut:    "02",
	http.MethodGet:    "03",
	http.MethodDelete: "04",
	http.MethodPatch:  "05",
}

type Router interface {
	Group(pattern string, fn func())
	Get(pattern, prenum string, classic interface{})
	Patch(pattern, prenum string, classic interface{})
	Post(pattern, prenum string, classic interface{})
	Put(pattern, prenum string, classic interface{})
	Delete(pattern, prenum string, classic interface{})
	Sources() []*Source
}

type sourceList []*Source

type Source struct {
	Pattern string
	Serial  string
	Method  string
	Classic interface{}
}

type router struct {
	isSort  bool
	sources []*Source
	groups  []string
}

func NewRouter() Router {
	return &router{
		sources: make([]*Source, 0, 1024),
	}
}

func (r *router) Sources() []*Source {
	if !r.isSort {
		r.isSort = true
		sort.Sort(sourceList(r.sources))
	}
	return r.sources
}

func (r *router) Group(pattern string, fn func()) {
	r.groups = append(r.groups, "/"+strings.Trim(pattern, "/"))
	fn()
	r.groups = r.groups[:len(r.groups)-1]
}

func (r *router) Get(pattern, prenum string, classic interface{}) {
	r.handle(http.MethodGet, pattern, prenum, classic)
}

func (r *router) Post(pattern, prenum string, classic interface{}) {
	r.handle(http.MethodPost, pattern, prenum, classic)
}

func (r *router) Put(pattern, prenum string, classic interface{}) {
	r.handle(http.MethodPut, pattern, prenum, classic)
}

func (r *router) Delete(pattern, prenum string, classic interface{}) {
	r.handle(http.MethodDelete, pattern, prenum, classic)
}

func (r *router) Patch(pattern, prenum string, classic interface{}) {
	r.handle(http.MethodPatch, pattern, prenum, classic)
}

func (r *router) Any(pattern, prenum string, classic interface{}) {
	r.handle(http.MethodGet, pattern, prenum, classic)
	r.handle(http.MethodPost, pattern, prenum, classic)
	r.handle(http.MethodHead, pattern, prenum, classic)
	r.handle(http.MethodDelete, pattern, prenum, classic)
	r.handle(http.MethodPatch, pattern, prenum, classic)
	r.handle(http.MethodPut, pattern, prenum, classic)
}

func (r *router) handle(method, pattern, prenum string, classic interface{}) {
	if !IsDigit(prenum) {
		panic("The prenum must be a numeric string")
	}
	if reflect.TypeOf(classic).Kind() != reflect.Ptr {
		panic("The classic must be a pointer")
	}
	pattern = "/" + strings.Trim(pattern, "/")
	if len(r.groups) > 0 {
		groupPattern := ""
		for _, prefix := range r.groups {
			groupPattern += prefix
		}
		pattern = groupPattern + pattern
	}

	r.sources = append(r.sources, &Source{
		Pattern: pattern,
		Method:  method,
		Serial:  prenum + serialNum[method],
		Classic: classic,
	})
}

func (sl sourceList) Len() int {
	return len(sl)
}

func (sl sourceList) Swap(i, j int) {
	sl[i], sl[j] = sl[j], sl[i]
}

func (sl sourceList) Less(i, j int) bool {
	ii, _ := strconv.Atoi(sl[i].Serial)
	ij, _ := strconv.Atoi(sl[j].Serial)
	return ii < ij
}

func IsDigit(ch string) bool {
	for i := 0; i < len(ch); i++ {
		if !('0' <= ch[i] && ch[i] <= '9') {
			return false
		}
	}
	return true
}
