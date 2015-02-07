package theTree

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/sumitasok/theTree/byteOps"
)

func SetNodeValue(node *Node, byteArr []byte) error {
	byteArr = bytePluckByteRecursively(byteArr, R_SPACE, R_SPACE)

	// fmt.Println(string(byteArr))
	node.Set(byteArr)

	if byteIs(byteArr, R_OPEN_CURL) {
		if byteArr, err := byteStripOuterCurls(byteArr); err != nil {
			return err
		} else {
			SetNodeValue(node, byteArr)
		}
	} else if byteIs(byteArr, R_OPEN_SQUARE) {
		return errors.New(("arrays undefined"))
	} else if bytes.Contains(byteArr, []byte(`{`)) == false {
		node.Set(byteArr)
		return nil
	}

	// value := byteArr

	// parseNodes(node, value)

	kvList := parseKeyValue(byteArr)

	for i, kv := range kvList {
		nC, e := node.Append(string(kv.Key))
		if e == nil {
			SetNodeValue(nC, kv.Value)
		} else {
			return errors.New(fmt.Sprintf("key couldnt be identified %i %s", i, e))
		}
	}

	return nil
}

type KV struct {
	Key   []byte
	Value []byte
}

func parseKeyValue(byteArr []byte) []KV {

	var kvList = make([]KV, 0)

	cntIdx := 0

	if len(byteArr) == 0 {
		return kvList
	}

	key := []byte{}
	value := []byte{}

	for i, b := range byteArr {
		if b == 58 {
			key = byteArr[cntIdx:i]
			key = byteOps.PrepareKeyByte(key)
			fmt.Println(printStrOfBytes(key))
			cntIdx = i + 1
		}

		if b == 44 {
			value = byteArr[cntIdx:i]
			value = byteOps.PrepareValueByte(value)
			cntIdx = i + 1
			kv := KV{key, value}
			kvList = append(kvList, kv)
		}

		if i == len(byteArr)-1 {
			value = byteArr[cntIdx:]
			value = byteOps.PrepareValueByte(value)
			cntIdx = i + 1
			kv := KV{key, value}
			kvList = append(kvList, kv)
		}
	}

	return kvList

}

func parseNodes(node *Node, value []byte) {
	byteArr := value
	// go through each , seperated key-value pair
	// and create a child

	byteArrOfArr, err := byteSplitKeyValue(byteArr, []byte(`:`), 2)
	keyPart := byteArrOfArr[0]

	if len(byteArrOfArr) > 1 {
		valuePart := byteArrOfArr[1]
		if err == nil {
			keyStr := byteOps.PrepareKey(keyPart)
			thisNodeC, errAppend := node.Append(keyStr)
			if errAppend == nil {

				// split on next, and then set value,
				// and send next parse with remaining value
				baa, errGetValue := GetValue(valuePart)
				// strip the first value to append
				if errGetValue == nil {
					thisChildValue := baa[0]
					nextChildKeyValue := baa[1]
					SetNodeValue(thisNodeC, thisChildValue)
					SetNodeValue(node, nextChildKeyValue)
					// fmt.Println("nextChildKeyValue", string(nextChildKeyValue))
				} else {
					thisNodeC, err := node.Append(keyStr)
					if err != nil {
						// fmt.Println("In err not nil", string(valuePart))
						thisNodeC.Set(valuePart)
					} else {
						// fmt.Println("In err nil", string(valuePart))
						thisNodeC.Set(valuePart)
					}
				}
			} else {
				// c, e :=
				node.UpdateChild(keyStr, valuePart)
				// fmt.Println("If Update", c, "If Error", e)
				// fmt.Println("If Update Value", c.Value)
			}
		}
	}
}
