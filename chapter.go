package shelf

func NewChapter(name, url string) Chapter {
	return Chapter{
		Name: name,
		URL:  url,
	}
}

type Chapter struct {
	Name string
	URL  string
}

func NewChapterDetail(chapter Chapter, content, nextURL string) ChapterDetail {
	return ChapterDetail{
		Chapter: chapter,
		Content: content,
		NextURL: nextURL,
	}
}

type ChapterDetail struct {
	Chapter
	Content string
	NextURL string
}
