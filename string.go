package theTree

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/sumitasok/theTree/byteOps"
)

const (
	R_OPEN_CURL      = 123
	R_CLOSED_CURLED  = 125
	R_COLON          = 58
	R_SPACE          = 32
	R_DOUBLE_QOUTE   = 34
	R_SLASH_N        = 10
	R_OPEN_SQUARE    = 91
	R_CLOSED_SQUARE  = 93
	R_BACKWARD_SLASH = 92
	R_TAB            = 9
	R_COMMA          = 44
)

func parseNode(node *Node, value []byte) {
	byteArr := value
	// go through each , seperated key-value pair
	// and create a child

	str := byteArr
	fmt.Println("parseNode 1", node.Key, string(str))

	byteArrOfArr, err := byteSplitKeyValue(byteArr, []byte(`:`), 2)
	keyPart := byteArrOfArr[0]
	fmt.Println("keyPart", string(keyPart))
	if len(byteArrOfArr) > 1 {
		valuePart := byteArrOfArr[1]
		if err == nil {
			fmt.Println("err nil in byteSplitKeyValue", string(keyPart))
			keyStr := byteOps.PrepareKey(keyPart)
			newNode, errAppend := node.Append(keyStr)
			if errAppend == nil {
				fmt.Println("err nil in Append", string(valuePart))
				// split on next, and then set value,
				// and send next parse with remaining value
				baa, errGetValue := GetValue(valuePart)
				fmt.Println(errGetValue)
				if errGetValue == nil {
					thisChildValue := baa[0]
					nextChildKeyValue := baa[1]
					fmt.Println("err nil in GetValue", string(thisChildValue))
					fmt.Println("parseNode 3", newNode.Key, string(thisChildValue))
					SetNodeValue(newNode, thisChildValue)

					fmt.Println("parseNode 4", node.Key, string(baa[1]))
					SetNodeValue(node, nextChildKeyValue)
				} else {
					newNode, err := node.Append(string(keyPart))
					if err != nil {
						fmt.Println(err, "cannot append")
						newNode.Set(valuePart)
					} else {
						newNode.Set(valuePart)
					}
				}
			}
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
