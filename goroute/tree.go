package goroute

import (
	"fmt"
	"net/http"
	"strings"
)

type nodeType uint8

const wildPrefix string = ":"

const (
	static nodeType = iota
	param
	wild
	catchAll
)

//node represents a level in the trie
type node struct {
	path string

	children map[string]*node

	wildChild *node

	nodeType

	handle http.HandlerFunc
}

//Insert a new path into the trie
func (n *node) insert(path string, handle http.HandlerFunc) error {
	path = strings.TrimPrefix(path, "/")
	parts := strings.Split(path, "/")
	curr := n

	for _, part := range parts {
		if next, contains := curr.children[part]; contains {
			curr = next
		} else if curr.wildChild != nil && isWild(part) {
			if strings.TrimPrefix(part, wildPrefix) != curr.wildChild.path {
				panic("param conflict inserting path " + path + " for param " + part + " != " + curr.wildChild.path)
			}
			curr = curr.wildChild

		} else {
			newNode := new(node)
			if isWild(part) {
				newNode.nodeType = wild
				newNode.path = strings.TrimPrefix(part, wildPrefix)
				curr.wildChild = newNode
			} else {
				if curr.children == nil {
					curr.children = make(map[string]*node)
				}
				newNode.nodeType = static
				newNode.path = part
				curr.children[part] = newNode
			}
			curr = newNode
		}
	}
	curr.handle = handle
	return nil
}

func isWild(path string) bool {
	return strings.HasPrefix(path, wildPrefix)
}

//PrintTrie prints the trie to standard out
func (n *node) printTrie(depth string) {
	for path, next := range n.children {
		fmt.Println(depth + path)
		fmt.Println(n.children[path].handle)
		next.printTrie(depth + " ")
	}
	if n.wildChild != nil {
		fmt.Println(depth + n.wildChild.path)
		n.wildChild.printTrie(depth + " ")
	}
}

func (n *node) searchIgnoreCase(path string) (handle http.HandlerFunc, params map[string]string, success bool) {
	path = strings.TrimPrefix(path, "/")
	parts := strings.Split(path, "/")
	curr := n
	if len(path) > 0 {
	walk:
		for _, part := range parts {
			for path, next := range curr.children {
				if strings.ToLower(path) == strings.ToLower(part) {
					curr = next
					continue walk
				}
			}

			if curr.wildChild != nil {
				curr = curr.wildChild
				if params == nil {
					params = make(map[string]string)
				}
				params[curr.path] = part
				continue walk
			}

			success = false
			return
		}
	}
	handle = curr.handle
	success = true
	return
}

func (n *node) search(path string) (handle http.HandlerFunc, params map[string]string, success bool) {
	path = strings.TrimPrefix(path, "/")
	parts := strings.Split(path, "/")
	curr := n
	for _, part := range parts {
		if next, contains := curr.children[part]; contains {
			curr = next
		} else if curr.wildChild != nil {
			curr = curr.wildChild
			if params == nil {
				params = make(map[string]string)
			}
			params[curr.path] = part
		} else {
			success = false
			return
		}
	}
	handle = curr.handle
	success = true
	return
}
