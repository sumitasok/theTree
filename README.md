# theTree
Tree like datastructure

Node is a data struct which allows nested key value pair data. Node provides easy methods to traverse the data in it.



To initialize the node
pass Normal{} as the Engine, this is kept for a later extension, if some one wants to override the methods.

```
node = Init(Normal{}, "name")
```
Set the Value for the Key.
Value can be anything as it is of datatype interface{}

```
node.Set("Bob")
```
If you wish to know the data type of the Value in the Node

```
node.DataType()
```

In order to append a Child key-value pair to the current Node

```
child, err := node.Append("age")
child.Set(27)
```
This method returns an error when the Child with key as 'age' exists

In order to Update an existing child's value

```
child, err = node.UpdateChild("age", 22)
```
To find a child node with the key

```
child, errFound := node.Child("age")
```
throughs error if the child with the key is not present

To find the child node from an ancestral path extending from root node(the node on which we are calling) to child node.

```
child, actErr = node.Find("name:age")
```
Throughs error if the node is not found.
The nesting can go any long as the developer wants it to. eg `node.Find("name:age:category:employed")`

In order to find the heirarchical pat of a child node.

```
node.Ancestry()
```

To get the root node from any child node, call

```
child.Root()
```

To find the number of children the node has

```
node.Count()
```

To find the number of nodes till root node from the child node

```
child.CountPre()
```
To count all the children till tail node

```
node.CountDeep()
```

---
