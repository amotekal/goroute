package goroute

import (
	"fmt"
	"strings"
)

//Router is
type Router struct {
	trees map[string]*node
}

//New is
func New() *Router {
	return &Router{}
}

//Handle is
func (r *Router) Handle(method, path string) {
	if path[0] != '/' {
		panic("path must begit with '/' in path'" + path + "'")
	}

	path = strings.TrimPrefix(path, "/")

	if r.trees == nil {
		r.trees = make(map[string]*node)
	}

	root := r.trees[method]
	if root == nil {
		root = new(node)
		r.trees[method] = root
	}

	root.insert(path)
}

//Get is
func (r *Router) Get(path string) {
	r.Handle("GET", path)
}

//PrintTrees is
func (r *Router) PrintTrees() {
	for method, tree := range r.trees {
		fmt.Println(method)
		tree.printTrie("")
	}
}
