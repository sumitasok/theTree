package theTree

import (
	"strings"
)

func Parse(data []byte) (*Node, error) {
	str := string(data)

	return Init(Normal{}, str), nil
}

func stripOuterCurls(str string) string {
	initialIndex := strings.Index(str, "{")
	finalIndex := strings.LastIndex(str, "}")
	return str[initialIndex+1 : finalIndex]
}
