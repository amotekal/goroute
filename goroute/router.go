package goroute

//Router is
type Router struct {
	tree map[string]*node
}

//New is
func New() *Router {
	return &Router{}
}

//Handle is
func (r *Router) Handle(method, path string) {
	if r.tree == nil {
		r.tree = make(map[string]*node)
	}

	root := r.tree[method]
	if root == nil {
		root = new(node)
		r.tree[method] = root
	}

	root.insert(path)
}

//Get is
func (r *Router) Get(path string) {
	r.Handle("GET", path)
}

//PrintTrees is
func (r *Router) PrintTrees() {
	for _, t := range r.tree {
		t.printTrie("")
	}
}
