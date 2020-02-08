package shelf

func NewBook(name, url, author, introduce string, chapter *chapter) book {
	return book{
		name:      name,
		url:       url,
		author:    author,
		introduce: introduce,
		chapter:   chapter,
	}
}

type book struct {
	name      string
	url       string
	author    string
	introduce string
	chapter   *chapter
}

func (b *book) Name() string {
	return b.name
}

func (b *book) URL() string {
	return b.url
}

func (b *book) Author() string {
	return b.author
}

func (b *book) Introduce() string {
	return b.introduce
}

func (b *book) Chapter() *chapter {
	return b.chapter
}

func NewBookDetail(book book, chapters []chapter) bookDetail {
	return bookDetail{
		book:     book,
		chapters: chapters,
	}
}

type bookDetail struct {
	book
	chapters []chapter
}

func (b *bookDetail) Chapters() []chapter {
	return b.chapters
}

func (b *bookDetail) ChapterAt(index int) chapter {
	return b.chapters[index]
}

func (b *bookDetail) SearchChapter(name string) (*chapter, bool) {
	for _, chapter := range b.chapters {
		if chapter.Name() == name {
			return &chapter, true
		}
	}
	return nil, false
}
