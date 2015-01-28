package theTree

type Node struct {
	Engine
	Core
}

type Core struct {
	Children []*Node
	Parent   *Node
	Key      string
	Value    interface{}
	dataType string
}

func Init(engine Engine, key string) *Node {
	core := Core{Key: key}
	return &Node{engine, core}
}
