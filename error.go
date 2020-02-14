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
	return fmt.Sprintf("error type:%s\ncause:%v\nscene:%v\n", e.Type, e.Cause, e.Scene)
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

func NewExecutorError(err error, req Request) error {
	return &ExecutorError{BaseError{
		Type:  "executor_error",
		Cause: err,
		Scene: req,
	}}
}

type UnsupportedEncodingError struct {
	BaseError
}

func NewUnsupportedEncodingError(err error, encoding string) error {
	return &UnsupportedEncodingError{BaseError{
		Type:  "unsupported_encoding",
		Cause: err,
		Scene: encoding,
	}}
}

type UnknownExtractorError struct {
	BaseError
}

func NewUnknownExtractorError(err error, extractor string) error {
	return &UnsupportedEncodingError{BaseError{
		Type:  "unknown_extractor",
		Cause: err,
		Scene: extractor,
	}}
}
