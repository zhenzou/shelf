package shelf

type Shelf interface {
	Search(name string) []Book
	Source(name string) (Source, bool)
}

type Source interface {
	Name() string
	Search(name string) []Book
	Classes() []Class
}

type Class interface {
	Name() string
	Search(name string) []Book
}

type Book interface {
	Source() Source
	Name() string
	Author() string
	Chapters() []Chapter
	ChapterAt(index int) Chapter
	SearchChapter(name string) (Chapter, bool)
}

type Chapter interface {
	Book() Book
	Index() int
	Name() int
	Content() string
}
