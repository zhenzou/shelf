package shelf

import (
	"strings"
)

func IsBlank(str string) bool {
	return strings.TrimSpace(str) == ""
}

func IsNotBlank(str string) bool {
	return !IsBlank(str)
}

// will trim empty split
func Split(s string, sep rune) []string {
	strs := []string{}
	sb := strings.Builder{}
	escape := false
	for _, r := range s {
		if r == '\\' {
			if escape {
				sb.WriteRune(r)
				escape = false
			} else {
				escape = true
			}
			continue
		}
		if r == ';' {
			if escape {
				sb.WriteRune(r)
				escape = false
			} else if sb.Len() > 0 {
				strs = append(strs, sb.String())
				sb.Reset()
			}
		} else {
			sb.WriteRune(r)
		}
	}

	if sb.Len() > 0 {
		strs = append(strs, sb.String())
	}

	return strs
}