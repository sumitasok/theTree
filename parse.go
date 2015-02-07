package theTree

import (
	"errors"
	"fmt"
	"github.com/sumitasok/theTree/byteOps"
)

func SetNodeValue(node *Node, byteArr []byte) error {
	byteArr = bytePluckByteRecursively(byteArr, R_SPACE, R_SPACE)

	fmt.Println(string(byteArr))
	node.Set(byteArr)

	if byteIs(byteArr, R_OPEN_CURL) {
		if byteArr, err := byteStripOuterCurls(byteArr); err != nil {
			return err
		} else {
			// node.Set(byteArr)
			// parseNode(node, byteArr)
			SetNodeValue(node, byteArr)
		}
	}

	if byteIs(byteArr, R_OPEN_SQUARE) {
		return errors.New(("arrays undefined"))
	}

	value := byteArr

	parseNodes(node, value)

	return nil
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
				if errGetValue == nil {
					thisChildValue := baa[0]
					nextChildKeyValue := baa[1]
					SetNodeValue(thisNodeC, thisChildValue)
					SetNodeValue(node, nextChildKeyValue)
					fmt.Println("nextChildKeyValue", string(nextChildKeyValue))
				} else {
					thisNodeC, err := node.Append(keyStr)
					if err != nil {
						thisNodeC.Set(valuePart)
					} else {
						thisNodeC.Set(valuePart)
					}
				}
			} else {
				node.UpdateChild(keyStr, valuePart)
			}
		}
	}
}
