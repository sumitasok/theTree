package theTree

import (
	"encoding/json"
	"errors"
	"fmt"
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
	parent   *Node
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
	child := Node{n.Engine, Core{Key: key, parent: n}}

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

func (n *Node) Ancestry() string {
	if n.parent == nil {
		return n.Key
	} else {
		return fmt.Sprintf("%s:%s", n.parent.Ancestry(), n.Key)
	}
}

func (n Node) Root() *Node {
	if n.parent == nil {
		return &n
	} else {
		return n.parent.Root()
	}
}

func (n Node) Count() int {
	return len(n.Children)
}

func (n Node) CountPre() int {
	if n.parent == nil {
		return 0
	} else {
		return 1 + n.parent.CountPre()
	}
}

func (n Node) CountDeep() int {
	if len(n.Children) == 0 {
		return 0
	} else {
		i := 0
		for _, child := range n.Children {
			i = i + child.CountDeep()
		}
		return len(n.Children) + i
	}
}

func (n Node) Json() ([]byte, error) {
	return json.Marshal(n)
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
