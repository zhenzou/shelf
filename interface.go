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
	ExtractBook(rule BookRule, html []byte) (BookDetail, error)
	ExtractChapter(rule ChapterRule, html []byte) (ChapterDetail, error)
	ExtractBooks(rule ListRule, html []byte) ([]Book, error)
}

type Executor interface {
	Exec(ctx context.Context, req Request) (Response, error)
}
