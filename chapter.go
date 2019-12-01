package shelf

import "context"

type chapter struct {
	book    Book
	name    string
	index   int
	url     string
	content string
}

func (c *chapter) URL() string {
	return c.url
}

func (c *chapter) Book() Book {
	return c.book
}

func (c *chapter) Index() int {
	return c.index
}

func (c *chapter) Name() string {
	return c.name
}

func (c *chapter) Content() string {
	return c.content
}

func newChapter(name, url string) Chapter {
	return &chapterImpl{
		name: name,
		url:  url,
	}
}

type chapterImpl struct {
	name string
	url  string
}

func (c chapterImpl) Get(ctx context.Context) (chapter, error) {
	panic("implement me")
}
