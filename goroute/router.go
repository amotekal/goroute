package goroute

import (
	"fmt"
	"net/http"
)

//Router is
type Router struct {
	trees map[string]*node

	PanicHandler func(http.ResponseWriter, *http.Request, interface{})
}

//New is
func New() *Router {
	return &Router{}
}

//Handle is
func (r *Router) Handle(method, path string, handle http.HandlerFunc) {
	if path[0] != '/' {
		panic("path must begit with '/' in path'" + path + "'")
	}

	if r.trees == nil {
		r.trees = make(map[string]*node)
	}

	root := r.trees[method]
	if root == nil {
		root = new(node)
		r.trees[method] = root
	}

	root.insert(path, handle)
}

//Get is
func (r *Router) Get(path string, handle http.HandlerFunc) {
	r.Handle("GET", path, handle)
}

//PrintTrees is
func (r *Router) PrintTrees() {
	for method, tree := range r.trees {
		fmt.Println(method)
		tree.printTrie("")
	}
}

func (r *Router) rcvr(w http.ResponseWriter, req *http.Request) {
	if rcv := recover(); rcv != nil {
		r.PanicHandler(w, req, rcv)
	}
}

//ServeHTTP implements http.handler
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if r.PanicHandler != nil {
		defer r.rcvr(w, req)
	}

	path := req.URL.Path

	if root := r.trees[req.Method]; root != nil {
		if handle, params, success := root.search(path); success {
			for key, val := range params {
				addParam(req, key, val)
			}

			handle(w, req)
		}
		return
	}
}
