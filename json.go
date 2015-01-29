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

func stripOuter(str string, sep string) string {
	initialIndex := strings.Index(str, sep)
	finalIndex := strings.LastIndex(str, sep)
	return str[initialIndex+1 : finalIndex]
}

func fetchInnerCurls(str string) Element {
	e := Element{Content: str}
	str = stripOuterCurls(str)
	e.Content = str

	for _, split := range strings.Split(e.Content, ",") {
		t := Element{Content: split}
		e.Next = append(e.Next, &t)
	}

	// for _, element := range e.Next {

	// }
	return e
}

type Element struct {
	Content string
	Next    []*Element
	Key     string
	Value   string
}

func (e *Element) OnlyColon() {
	if strings.Contains(e.Content, ",") == false {
		if strings.Contains(e.Content, "{") == false {
			if strings.Contains(e.Content, "}") == false {
				keyValue := strings.Split(e.Content, ":")
				if len(keyValue) == 2 {
					e.Key = stripOuter(keyValue[0], "\"")
					e.Value = stripOuter(keyValue[1], "\"")
				}
			}
		}
	}
}

// node is a data struct thure which allows
// nested key value pair data to be traversed in many ways
