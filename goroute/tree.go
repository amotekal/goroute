package goroute

import (
	"fmt"
	"net/http"
	"strings"
)

type nodeType int

const (
	static nodeType = iota
	param
	wild
	catchAll
)

//node represents a level in the trie
type node struct {
	path string

	children map[string]node

	wildChild *node

	nodeType

	handle *http.HandlerFunc
}

//Insert a new path into the trie
func (node *node) Insert(path string) error {
	parts := strings.Split(path, "/")
	curr := node

	for _, part := range parts {
		if n, contains := curr.children[part]; contains {
			curr = &n
		} else {
			newNode := new(node)
			curr.children[part] = newNode
			curr = &newNode
		}
	}
	return nil
}

//PrintTrie prints the trie to standard out
func (node *node) PrintTrie(depth string) {
	fmt.Println(depth + node.path)
	for _, n := range node.children {
		n.PrintTrie(depth + " ")
	}
}
