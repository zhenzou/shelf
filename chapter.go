package shelf

func NewChapter(name, url string) chapter {
	return chapter{
		name: name,
		url:  url,
	}
}

type chapter struct {
	name string
	url  string
}

func (c *chapter) URL() string {
	return c.url
}

func (c *chapter) Name() string {
	return c.name
}

func NewChapterDetail(chapter chapter, content string) chapterDetail {
	return chapterDetail{
		chapter: chapter,
		content: content,
	}
}

type chapterDetail struct {
	chapter
	content string
}

func (c *chapterDetail) Content() string {
	return c.content
}
