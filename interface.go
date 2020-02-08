package shelf

import (
	"context"
	"net/url"
)

type Shelf interface {
	Sources() []Source
	SourceByName(name string) (Source, bool)
	SourceByURL(url url.URL) (Source, bool)
	Search(ctx context.Context, name string) (map[string][]book, error)
}

type Source interface {
	Name() string
	GetBookDetail(ctx context.Context, url string) (bookDetail, error)
	GetChapterDetail(ctx context.Context, url string) (chapterDetail, error)
	Search(ctx context.Context, name string) ([]book, error)
}

type Extractor interface {
	ExtractBook() (book, error)
	ExtractChapter() (chapter, error)
}

type Request struct {
	Method string
	URL    string
	Args   Args
}

type Response struct {
	Request
	Data []byte
}

type Executor interface {
	Exec(req Request) (Response, error)
}
