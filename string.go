package theTree

import (
// "bytes"
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
		byteArr = bytePluckByte(byteArr, R_SPACE)

		if byteIs(byteArr, R_OPEN_CURL) {

		}
	}
}

func bytePluckByte(byteArr []byte, byteChar byte) []byte {
	byteArr = byteRemoveByte(byteArr, byteChar)
	byteArr = byteRemoveByteFromBack(byteArr, byteChar)
	return byteArr
}

func byteStripOuterCurls(byteArr []byte) []byte {
	return byteArr
}

func byteRemoveByte(byteArr []byte, byteChar byte) []byte {
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

func byteRemoveByteFromBack(byteArr []byte, byteChar byte) []byte {
	revByteArr := Reverse(byteArr)
	revByteArr = byteRemoveByte(revByteArr, byteChar)
	return Reverse(revByteArr)
}

func byteIs(byteArr []byte, byteChar byte) bool {
	byteArr = byteRemoveByte(byteArr, R_SPACE)
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
	byteArr = byteRemoveByte(byteArr, R_SPACE)
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