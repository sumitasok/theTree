package theTree

import (
	"bytes"
	"errors"
	"fmt"
)

const (
	R_OPEN_CURL     = 123
	R_CLOSED_CURLED = 125
	R_COLON         = 58
	R_SPACE         = 32
	R_DOUBLE_QOUTE  = 34
	R_SLASH_N       = 10
	R_OPEN_SQUARE   = 91
	R_CLOSED_SQUARE = 93
	R_TAB           = 9
	R_COMMA         = 44
)

func SetNodeValue(node *Node, byteArr []byte) error {
	byteArr = bytePluckByteRecursively(byteArr, R_SPACE, R_SPACE)

	node.Set(byteArr)

	if byteIs(byteArr, R_OPEN_CURL) {
		if byteArr, err := byteStripOuterCurls(byteArr); err != nil {
			return err
		} else {
			node.Set(byteArr)
		}
	}

	if byteIs(byteArr, R_OPEN_SQUARE) {
		return errors.New(("arrays undefined"))
	}

	value := byteArr

	parseNode(node, value)

	return nil
}

func parseNode(node *Node, value []byte) {
	byteArr, _ := node.Value.([]uint8)
	// go through each , seperated key-value pair
	// and create a child

	str, _ := node.Value.([]uint8)
	fmt.Println("parseNode 1", node.Key, string(str))

	byteArrOfArr, err := byteSplitKeyValue(byteArr, []byte(`:`), 2)
	if err == nil {
		fmt.Println("err nil in byteSplitKeyValue", string(byteArrOfArr[0]))
		newNode, errAppend := node.Append(string(byteArrOfArr[0]))
		if errAppend == nil {
			fmt.Println("err nil in Append", string(byteArrOfArr[1]))
			// split on next, and then set value,
			// and send next parse with remaining value
			baa, errGetValue := GetValue(byteArrOfArr[1])
			if errGetValue == nil {
				fmt.Println("err nil in GetValue")
				SetNodeValue(newNode, baa[0])
				fmt.Println("parseNode 3", newNode.Key, string(baa[0]))

				SetNodeValue(node, baa[1])
				fmt.Println("parseNode 4", node.Key, string(baa[1]))

				parseNode(node)
			}
			fmt.Println("parseNode 2", newNode.Key, string(byteArrOfArr[1]))
			SetNodeValue(newNode, byteArrOfArr[1])
		}
	}
}

func GetValue(byteArr []byte) ([][]byte, error) {
	curl_count := 0
	sqaure_count := 0
	quotes := 0
	for i, b := range byteArr {
		if b == R_COMMA && curl_count == 0 && sqaure_count == 0 && quotes == 0 {
			return [][]byte{byteArr[:i], byteArr[i+1:]}, nil
		}
		// if b == R_OPEN_CURL {
		// 	curl_count = curl_count + 1
		// }
		// fmt.Println(i, b)
	}
	return [][]byte{}, errors.New("error")
}

func byteSplitKeyValue(byteArr, sep []byte, n int) ([][]byte, error) {
	list := bytes.SplitN(byteArr, sep, n)
	if len(list) < 2 {
		return list, errors.New("error")
	}
	return list, nil
}

func bytePluckByteRecursively(byteArr []byte, byteChar byte, byteBackChar byte) []byte {
	byteArr = byteRemoveByteRecursively(byteArr, byteChar)
	byteArr = byteRemoveByteRecursivelyFromBack(byteArr, byteBackChar)
	return byteArr
}

func byteRemove(byteArr []byte, byteChar byte) ([]byte, error) {
	if len(byteArr) == 0 {
		return byteArr, errors.New("error")
	} else if byteArr[0] == byteChar {
		return byteArr[1:], nil
	} else {
		return byteArr, errors.New("error")
	}
}

func byteStripOuterCurls(byteArr []byte) ([]byte, error) {
	if byteArr, err := byteRemove(byteArr, R_OPEN_CURL); err == nil {
		if byteArr, err = byteRemove(Reverse(byteArr), R_CLOSED_CURLED); err == nil {
			byteArr = Reverse(byteArr)
			return byteArr, nil
		} else {
			byteArr = Reverse(byteArr)
			return byteArr, err
		}
	} else {
		return byteArr, errors.New("error")
	}
}

func byteRemoveByteRecursively(byteArr []byte, byteChar byte) []byte {
	if len(byteArr) == 0 {
		return byteArr
	} else {
		if byteArr[0] != byteChar {
			return byteArr
		} else {
			return byteRemoveSpace(byteArr[1:])
		}
	}
}

func byteRemoveByteRecursivelyFromBack(byteArr []byte, byteChar byte) []byte {
	revByteArr := Reverse(byteArr)
	revByteArr = byteRemoveByteRecursively(revByteArr, byteChar)
	return Reverse(revByteArr)
}

func byteIs(byteArr []byte, byteChar byte) bool {
	byteArr = byteRemoveByteRecursively(byteArr, R_SPACE)
	if len(byteArr) == 0 {
		return false
	} else if byteArr[0] == byteChar {
		return true
	} else {
		return false
	}
}

// -------------------------------------------------

func byteIsHash(byteArr []byte) bool {
	byteArr = byteRemoveByteRecursively(byteArr, R_SPACE)
	if len(byteArr) == 0 {
		return false
	} else if byteArr[0] == R_OPEN_CURL {
		return true
	} else {
		return false
	}
}

func byteRemoveSpace(byteArr []byte) []byte {
	if len(byteArr) == 0 {
		return byteArr
	} else {
		if byteArr[0] != R_SPACE {
			return byteArr
		} else {
			return byteRemoveSpace(byteArr[1:])
		}
	}
}
