package theTree

import (
	"errors"
	"reflect"
	"strings"
)

const (
	KEY_SEPERATOR          = ":"
	err_key_already_exists = "Key already exists, use Update"
	err_key_doesnt_exist   = "Key doesn't exist"
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

	if child, err := n.Child(key); err == nil {
		return child, errors.New(err_key_already_exists)
	}

	n.Children = append(n.Children, &child)

	return &child, nil
}

func (n *Node) UpdateChild(key string, value interface{}) (*Node, error) {
	if child, err := n.Child(key); err == nil {
		child.Value = value
		return child, nil
	}
	child := Node{}
	return &child, errors.New(err_key_doesnt_exist)

}

func (n *Node) Child(key string) (*Node, error) {
	return find(n.Children, key)
}

func (n *Node) Find(key string) (*Node, error) {
	keySlice := strings.Split(key, KEY_SEPERATOR)

	if n.Key == keySlice[0] {
		if len(keySlice) == 1 {
			return n, nil
		}
		next, err := find(n.Children, keySlice[1])
		if err == nil {
			return next.Find(strings.Join(keySlice[1:], KEY_SEPERATOR))
		} else {
			return &Node{}, errors.New(err_key_doesnt_exist)
		}
	} else {
		return &Node{}, errors.New(err_key_doesnt_exist)
	}
}

func find(children []*Node, key string) (*Node, error) {
	if len(children) == 0 {
		return &Node{}, errors.New(err_key_doesnt_exist)
	} else {
		if children[0].Key == key {
			return children[0], nil
		} else {
			return find(children[1:], key)
		}
	}
}
