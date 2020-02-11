package shelf

import "strings"

func IsBlank(str string) bool {
	return strings.TrimSpace(str) == ""
}

func IsNotBlank(str string) bool {
	return !IsBlank(str)
}
