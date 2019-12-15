package shelf

import (
	"context"
	"net/url"
)

type Shelf interface {
	Sources() []Source
	SourceByName(name string) (Source, bool)
	SourceByURL(url url.URL) (Source, bool)
	Search(ctx context.Context, name string) ([]Book, error)
	Extractor(name string, url url.URL) (Extractor, error)
}

type Source interface {
	Name() string
	Search(ctx context.Context, name string) ([]Book, error)
	Classes(ctx context.Context) ([]Class, error)
}

type Class interface {
	Name() string
	Search(ctx context.Context, name string) []Book
}

type Book interface {
	Get(ctx context.Context) (book, error)
}

type Chapters interface {
	ChapterAt(index int) Chapter
	SearchChapter(name string) (Chapter, bool)
}

type Chapter interface {
	Get(ctx context.Context) (chapter, error)
}

type Extractor interface {
	ExtractBook() (book, error)
	ExtractChapter() (chapter, error)
}
