package goroute

import (
	"fmt"
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

//Call is
func (r *Router) Call(method string, path string) {
	if root := r.trees[method]; root != nil {
		params, success := root.search(path)

		if success {
			fmt.Println("success finding " + path)
			if params != nil {
				for param, value := range params {
					fmt.Println(param + " : " + value)
				}
			}
		} else {
			fmt.Println("failed to find " + path)
		}
	}
}
