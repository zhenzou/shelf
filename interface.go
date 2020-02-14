package shelf

import (
	"context"
)

type Source interface {
	Name() string
	Rule() SourceRule
	GetBookDetail(ctx context.Context, url string) (BookDetail, error)
	GetChapterDetail(ctx context.Context, url string) (ChapterDetail, error)
	Search(ctx context.Context, name string) ([]Book, error)
}

type Extractor interface {
	ExtractBook(rule BookRule, url string, html []byte) (BookDetail, error)
	ExtractChapter(rule ChapterRule, url string, html []byte) (ChapterDetail, error)
	ExtractBooks(rule ListRule, url string, html []byte) ([]Book, error)
}

type Executor interface {
	Exec(ctx context.Context, req Request) (Response, error)
}
