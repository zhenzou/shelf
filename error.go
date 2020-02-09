package shelf

import (
	"fmt"
)

type BaseError struct {
	Type  string
	Cause error
	Scene interface{}
}

func (e *BaseError) Error() string {
	return fmt.Sprintf("error type:%s\ncause:%s\nscene:%v\n", e.Type, e.Cause.Error(), e.Scene)
}

func (e *BaseError) Is(err error) bool {
	if other, ok := err.(*BaseError); ok {
		return other.Type == other.Type
	}
	return false
}

type HTMLParseError struct {
	BaseError
}

func NewHTMLParseError(err error, html []byte) error {
	return &HTMLParseError{BaseError{
		Type:  "parse_html_error",
		Cause: err,
		Scene: string(html),
	}}
}

type ExecutorError struct {
	BaseError
}


func NewExecutorParseError(err error, req Request) error {
	return &ExecutorError{BaseError{
		Type:  "executor_error",
		Cause: err,
		Scene: req,
	}}
}