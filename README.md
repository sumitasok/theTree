# theTree
Tree like datastructure


You initialize the data

```
node = Init(Engine{}, "root-key-name")
```
Set the Value for the Key
Value can be anything as it is of datatype interface{}

```
node.Set("Value")
```
If you wish to know the data type of the Value in the Node

```
node.DataType()
```

Node.Child(key)
Node.Value

Node.Set(value)
    Also sets the dataType using reflect
Node.Append(childKey) (Node, error)
    Also set Parent of the child Node as Self

Node.Init(key) Node
Node.Append(key, value) (Node, error) error if key already existing, cannot be appended, instead has to use Update.
Node.Update(key, value) (Node, error)

Check if there is any tree implementation in Go
Learn how error can also have nil as value. Does error embed nil?

Node.Find("colon:separated:keys")
Node.FindFirst("key") length first or breadth first traversal
Node.Ancestry() string colon:separated:keys till root node
Node.Root() Node
Node.Count()
Node.CountPost()
Node.CountPre()
