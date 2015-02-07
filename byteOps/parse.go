package byteOps

import (
	// "fmt"
	"bytes"
)

func PrepareKeyByte(byteArr []byte) []byte {
	ignoreList := []byte(`!"#$%&'()*+,-./:;<=>?@[\]^_{|}~ `)
	return Prepare(byteArr, ignoreList)
}

func PrepareValueByte(byteArr []byte) []byte {
	byteArr = bytes.TrimSpace(byteArr)
	ignoreList := []byte(`"`)

	byteLen := len(byteArr)
	revByteArr := reverse(byteArr)

	startIndex, endIndex := 0, 0

	for i, b := range byteArr {
		if (bytes.Contains(ignoreList, []byte{b}) || b == 96) == false {
			startIndex = i
			break
		}
	}

	for i, b := range revByteArr {
		if (bytes.Contains(ignoreList, []byte{b})) == false {
			endIndex = byteLen - i
			break
		}
	}

	return byteArr[startIndex:endIndex]
}

func Prepare(byteArr []byte, ignoreList []byte) []byte {
	byteLen := len(byteArr)
	revByteArr := reverse(byteArr)

	startIndex, endIndex := 0, 0

	for i, b := range byteArr {
		if (bytes.Contains(ignoreList, []byte{b}) || b == 96) == false {
			startIndex = i
			break
		}
	}

	for i, b := range revByteArr {
		if (bytes.Contains(ignoreList, []byte{b}) || b == 96) == false {
			endIndex = byteLen - i
			break
		}
	}

	return byteArr[startIndex:endIndex]
}

func PrepareKey(byteArr []byte) string {
	kByte := PrepareKeyByte(byteArr)
	return string(kByte)
}

func reverse(b []byte) []byte {
	r := make([]byte, len(b))

	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = b[j], b[i]
	}

	return r
}

//[0a, 1b, 2c , 3d]
// ri = 1
// len = 4
// i = len - r1 - 1
