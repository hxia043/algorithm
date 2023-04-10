package router

import (
	"fmt"
	"gin-trie/tree"
	"net/http"
)

type Router struct {
	roots    map[string]tree.Tree
	handlers map[string]http.HandlerFunc
}

func validatePattern(pattern string) bool {
	if rune(pattern[0]) != '/' || pattern == "" || pattern == "/" {
		return false
	}

	return true
}

func (r *Router) AddRoute(method, pattern string, handler http.HandlerFunc) {
	if !validatePattern(pattern) {
		return
	}

	if _, ok := r.roots[method]; !ok {
		r.roots[method] = tree.NewTree(method)
	}

	r.roots[method].Insert(pattern)

	key := method + "-" + pattern
	r.handlers[key] = handler
}

func (r *Router) GetRoute(method, pattern string) {
	if !validatePattern(pattern) {
		return
	}

	if r.roots[method].Select(pattern) {
		//key := method + "-" + pattern
		//r.Handlers[key](nil, nil)
		fmt.Println("find pattern")
		return
	}

	fmt.Println("no pattern find")
}

func (r *Router) RouteCount(method string) int {
	if _, ok := r.roots[method]; !ok {
		return 0
	}

	return r.roots[method].PatternCount()
}

func NewRouter() *Router {
	return &Router{
		roots:    make(map[string]tree.Tree),
		handlers: make(map[string]http.HandlerFunc),
	}
}
