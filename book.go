package shelf

func NewBook(name, url, author, introduce string, chapter *chapter) book {
	return book{
		Name:      name,
		URL:       url,
		Author:    author,
		Introduce: introduce,
		Chapter:   chapter,
	}
}

type book struct {
	Name      string
	URL       string
	Author    string
	Introduce string
	Chapter   *chapter
}

func NewBookDetail(book book, chapters []chapter) bookDetail {
	return bookDetail{
		book:     book,
		Chapters: chapters,
	}
}

type bookDetail struct {
	book
	Chapters []chapter
}

func (b *bookDetail) ChapterAt(index int) chapter {
	return b.Chapters[index]
}

func (b *bookDetail) SearchChapter(name string) (*chapter, bool) {
	for _, chapter := range b.Chapters {
		if chapter.Name == name {
			return &chapter, true
		}
	}
	return nil, false
}
