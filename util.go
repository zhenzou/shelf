package shelf

import "strings"

func IsBlank(str string) bool {
	return strings.TrimSpace(str) == ""
}
