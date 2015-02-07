package byteOps

import (
// "fmt"
)

func PrepareKey(byteArr []byte) string {

	byteLen := len(byteArr)
	revByteArr := reverse(byteArr)

	startIndex, endIndex := 0, 0

	for i, b := range byteArr {
		if (b > 96 && b < 123) || (b > 64 && b < 91) {
			startIndex = i
			break
		}
	}

	for i, b := range revByteArr {
		if (b > 96 && b < 123) || (b > 64 && b < 91) {
			endIndex = byteLen - i
			break
		}
	}

	return string(byteArr[startIndex:endIndex])
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
