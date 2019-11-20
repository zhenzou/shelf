package shelf

import (
	"context"
	"net/url"
)

type Shelf interface {
	Sources() []Source
	SourceByName(name string) (Source, bool)
	SourceByURL(url url.URL) (Source, bool)
	Search(ctx context.Context, name string) []Book
}

type Source interface {
	Name() string
	Search(ctx context.Context, name string) []Book
	Classes(ctx context.Context) []Class
}

type Class interface {
	Name() string
	Search(ctx context.Context, name string) []Book
}

type Book interface {
	Source() Source
	Name() string
	URL() string
	Author() string
	Chapters() []Chapter
	ChapterAt(index int) Chapter
	SearchChapter(name string) (Chapter, bool)
}

type Chapter interface {
	Book() Book
	Index() int
	Name() int
	URL() string
	Content(ctx context.Context) string
}

type Extract interface {
	ExtractBook(content string, rule BookRule) (Book, error)
}
