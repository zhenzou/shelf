package shelf

import "time"

func NewBook(name, url, author, introduce string, chapter *Chapter) Book {
	return Book{
		Name:          name,
		URL:           url,
		Author:        author,
		Introduce:     introduce,
		LatestChapter: chapter,
	}
}

type Book struct {
	Name             string
	URL              string
	Author           string
	Introduce        string
	LatestChapter    *Chapter
	LatestUpdateTime *time.Time
}

func NewBookDetail(book Book, chapters []Chapter) BookDetail {
	return BookDetail{
		Book:     book,
		Chapters: chapters,
	}
}

type BookDetail struct {
	Book
	Chapters []Chapter
}

func (b *BookDetail) ChapterAt(index int) Chapter {
	return b.Chapters[index]
}

func (b *BookDetail) SearchChapter(name string) (*Chapter, bool) {
	for _, chapter := range b.Chapters {
		if chapter.Name == name {
			return &chapter, true
		}
	}
	return nil, false
}
