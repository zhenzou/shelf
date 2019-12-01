package shelf

import "context"

type book struct {
	extractor Extractor
	name      string
	url       string
	author    string
	introduce string
	chapters  []chapter
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

// TODO cache
func (b *book) Chapters() []Chapter {
	chapters := make([]Chapter, 0, len(b.chapters))
	for _, chapter := range b.chapters {
		chapters = append(chapters, newChapter(chapter.name, chapter.url))
	}
	return chapters
}

func (b *book) ChapterAt(index int) Chapter {
	ch := b.chapters[index]
	return newChapter(ch.name, ch.url)
}

func (b *book) SearchChapter(name string) (Chapter, bool) {
	for _, chapter := range b.chapters {
		if chapter.Name() == name {
			return newChapter(chapter.name, chapter.url), true
		}
	}
	return nil, false
}

type bookImpl struct {
}

func (b bookImpl) Get(ctx context.Context) (book, error) {
	panic("implement me")
}
