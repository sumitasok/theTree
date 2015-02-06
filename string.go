package theTree

import (
	// "bytes"
	"errors"
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
)

func SetNodeValue(node *Node, byteArr []byte) {
	if node.parent == nil {
		// stripOuterCurls(string(byteArr))
		byteArr = bytePluckByteRecursively(byteArr, R_SPACE, R_SPACE)

		if byteIs(byteArr, R_OPEN_CURL) {

		}
	}
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

func byteStripOuterCurls(byteArr []byte) []byte {
	return byteArr
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
