package theTree

import (
	"reflect"
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
