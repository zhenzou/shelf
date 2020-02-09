package shelf

func NewChapter(name, url string) chapter {
	return chapter{
		Name: name,
		URL:  url,
	}
}

type chapter struct {
	Name string
	URL  string
}

func NewChapterDetail(chapter chapter, content string) chapterDetail {
	return chapterDetail{
		chapter: chapter,
		Content: content,
	}
}

type chapterDetail struct {
	chapter
	Content string
}
