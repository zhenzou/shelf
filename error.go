package shelf

import (
	"fmt"
)

type Error struct {
	Type  string
	Cause error
	Scene interface{}
}

func (e *Error) Error() string {
	return fmt.Sprintf("error type:%s\ncause:%s\nscene:%v\n", e.Type, e.Cause.Error(), e.Scene)
}

func (e *Error) Is(err error) bool {
	if other, ok := err.(*Error); ok {
		return other.Type == other.Type
	}
	return false
}

type HTMLParseError struct {
	Error
}

func NewHTMLParseError(err error, html []byte) error {
	return &HTMLParseError{Error{
		Type:  "parse_html_error",
		Cause: err,
		Scene: string(html),
	}}
}

type ExecutorError struct {
	Error
}


func NewExecutorParseError(err error, req Request) error {
	return &ExecutorError{Error{
		Type:  "executor_error",
		Cause: err,
		Scene: req,
	}}
}