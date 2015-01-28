package theTree

import (
	"errors"
	"reflect"
)

const (
	err_key_already_exists = "Key already exists, use Update"
)

type Node struct {
	Engine
	Core
}

type Core struct {
	Children []*Node
	Parent   *Node
	Key      string
	Value    interface{}
	DataType string
}

func Init(engine Engine, key string) *Node {
	core := Core{Key: key}
	return &Node{engine, core}
}

func (n *Node) Set(value interface{}) *Node {
	n.Value = value
	n.DataType = reflect.TypeOf(value).String()
	return n
}

func (n *Node) Append(key string) (*Node, error) {
	child := Node{n.Engine, Core{Key: key, Parent: n.Parent}}
	exists := false
	for _, c := range n.Children {
		if c.Key == key {
			exists = true
		}
	}

	if exists {
		child = Node{}
		return &child, errors.New(err_key_already_exists)
	}

	child.Key = key
	child.Parent = n
	n.Children = append(n.Children, &child)

	return &child, nil
}
