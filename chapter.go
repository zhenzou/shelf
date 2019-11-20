package shelf

import "context"

type chapter struct {
	book    Book
	name    string
	index   int
	url     string
	content string
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

func (c *chapter) URL() string {
	return c.url
}

func (c *chapter) Content(ctx context.Context) string {
	return c.content
}
