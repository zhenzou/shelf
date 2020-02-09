package shelf

import (
	"context"
	"net/http"
	"strconv"
	"strings"
)

type SourceArgs struct {
	Name string
	URL  string
}

func WithName(name string) func(args *SourceArgs) {
	return func(args *SourceArgs) {
		args.Name = name
	}
}

func WithURL(url string) func(args *SourceArgs) {
	return func(args *SourceArgs) {
		args.URL = url
	}
}

type Shelf interface {
	AddSource(rule SourceRule, extractor Extractor)
	Sources() map[string]Source
	Source(func(args *SourceArgs)) (Source, bool)
	Search(ctx context.Context, name string) (map[string][]book, error)
}

type Source interface {
	Name() string
	Rule() SourceRule
	GetBookDetail(ctx context.Context, url string) (bookDetail, error)
	GetChapterDetail(ctx context.Context, url string) (chapterDetail, error)
	Search(ctx context.Context, name string) ([]book, error)
}

type Extractor interface {
	ExtractBook(rule BookRule, url string, html []byte) (bookDetail, error)
	ExtractChapter(rule ChapterRule, url string, html []byte) (chapterDetail, error)
	ExtractBooks(rule ListRule, url string, html []byte) ([]book, error)
}

type Request struct {
	Method   string
	URL      string
	Args     Args
}

func (req Request) BuildRequest() (*http.Request, error) {
	url := req.URL
	url = strings.ReplaceAll(url, "${name}", req.Args.Name)
	url = strings.ReplaceAll(url, "${page}", strconv.FormatUint(req.Args.Page, 10))
	return http.NewRequest(req.Method, url, nil)
}

type Response struct {
	Request     Request
	RawResponse *http.Response
	Data        []byte
}

type Executor interface {
	Exec(ctx context.Context, req Request) (Response, error)
}
