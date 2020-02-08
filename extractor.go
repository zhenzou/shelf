package shelf

import (
	"bytes"

	"github.com/PuerkitoBio/goquery"
)

func newExtractor(url string, content []byte, rule BookRule) extractor {
	return extractor{
		content:  content,
		bookRule: rule,
	}
}

type extractor struct {
	url         string
	content     []byte
	bookRule    BookRule
	chapterRule ChapterRule
}

func (e extractor) ExtractBook() (bookDetail, error) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(e.content))
	if err != nil {
		return bookDetail{}, NewHTMLParseError(err, e.content)
	}
	bookRule := e.bookRule
	name := doc.Find(bookRule.Name).Text()
	author := doc.Find(bookRule.Author).Text()
	introduce := doc.Find(bookRule.Introduce).Text()

	chapters := []chapter{}
	doc.Find(e.chapterRule.URL).Each(func(i int, elm *goquery.Selection) {
		val, _ := elm.Attr("href")
		chapters = append(chapters, chapter{url: val})
	})
	doc.Find(e.chapterRule.Name).Each(func(i int, elm *goquery.Selection) {
		val := elm.Text()
		c := chapters[i]
		c.name = val
		chapters[i] = c
	})

	chapter := chapters[len(chapters)-1]

	return NewBookDetail(NewBook(name, e.url, author, introduce, &chapter), chapters), nil
}

func (e extractor) ExtractChapter() (chapterDetail, error) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(e.content))
	if err != nil {
		return chapterDetail{}, NewHTMLParseError(err, e.content)
	}
	chapterRule := e.chapterRule
	name := doc.Find(chapterRule.Name).Text()

	chapter := NewChapter(name, e.url)
	content := doc.Find(chapterRule.Content).Text()
	return NewChapterDetail(chapter, content), nil
}
