package shelf

type Shelf interface {
	Search(name string)
	Source(name string)
}

type Source interface {
	Name() string
	Search(name string)
	Class()
}

type Book interface {
	Name() string
	Author() string
	Platform() string
	Chapters()
}

type Chapter interface {
	Name() string
	Index() int
	Content() string
}
