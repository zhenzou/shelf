package shelf

import (
	"encoding/json"
	"strings"
	"unicode"
)

func IsBlank(str string) bool {
	return strings.TrimSpace(str) == ""
}

func IsNotBlank(str string) bool {
	return !IsBlank(str)
}

// will trim empty split and trim space
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
		} else if r == ';' {
			if escape {
				sb.WriteRune(r)
				escape = false
			} else if sb.Len() > 0 {
				strs = append(strs, strings.TrimSpace(sb.String()))
				sb.Reset()
			}
		} else {
			sb.WriteRune(r)
		}
	}

	if sb.Len() > 0 {
		strs = append(strs, strings.TrimSpace(sb.String()))
	}

	return strs
}

// Clean: remove useless new line
func Clean(s string) string {
	sb := strings.Builder{}
	sb.Grow(len(s))
	var prev rune
	for _, r := range s {
		if r == '\n' {
			if prev == '\n' {
				continue
			}
		}
		prev = r
		sb.WriteRune(r)
	}
	return sb.String()
}

// Compress: remove useless new line
func Compress(s string) string {
	sb := strings.Builder{}
	sb.Grow(len(s))
	var prev bool
	for _, r := range s {
		cur := unicode.IsSpace(r)
		if cur {
			if prev {
				continue
			}
		}
		prev = cur
		sb.WriteRune(r)
	}
	return sb.String()
}

func WithDefault(str, def string) string {
	if IsBlank(str) {
		return def
	} else {
		return str
	}
}

func ToJSON(obj interface{}) string {
	bs, _ := json.Marshal(obj)
	return string(bs)
}
