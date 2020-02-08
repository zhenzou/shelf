package shelf

import (
	"fmt"
)

const (
	ParseHTMLError = "parse_html_error"
)

type Error struct {
	Type  string
	Cause error
}

func (e *Error) Error() string {
	return fmt.Sprintf("error type:%s,cause:%s", e.Type, e.Cause.Error())
}

func (e *Error) Is(err error) bool {
	if other, ok := err.(*Error); ok {
		return other.Type == other.Type
	}
	return false
}
