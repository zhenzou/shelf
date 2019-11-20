package shelf

type book struct {
	source    Source
	extractor Extractor
	name      string
	url       string
	author    string
	introduce string
	chapters  []Chapter
}

func (b *book) Source() Source {
	return b.source
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

func (b *book) Chapters() []Chapter {
	return b.chapters
}

func (b *book) ChapterAt(index int) Chapter {
	return b.chapters[index]
}

func (b *book) SearchChapter(name string) (Chapter, bool) {
	for _, chapter := range b.chapters {
		if chapter.Name() == name {
			return chapter, true
		}
	}
	return nil, false
}
