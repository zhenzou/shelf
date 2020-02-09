package shelf

import (
	"bytes"

	"github.com/PuerkitoBio/goquery"
)

type extractor struct {
}

func (e extractor) ExtractBook(rule BookRule, url string, html []byte) (bookDetail, error) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(html))
	if err != nil {
		return bookDetail{}, NewHTMLParseError(err, html)
	}
	name := e.text(doc.Selection, rule.Name)
	author := e.text(doc.Selection, rule.Author)
	introduce := e.text(doc.Selection, rule.Introduce)

	chapters := []chapter{}

	doc.Find(rule.Chapter.List).Each(func(i int, elm *goquery.Selection) {
		chapter := e.extractChapter(elm, rule.Chapter)
		chapters = append(chapters, chapter)
	})
	chapter := chapters[len(chapters)-1]
	return NewBookDetail(NewBook(name, url, author, introduce, &chapter), chapters), nil
}

func (e extractor) ExtractChapter(rule ChapterRule, url string, html []byte) (chapterDetail, error) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(html))
	if err != nil {
		return chapterDetail{}, NewHTMLParseError(err, html)
	}
	name := e.text(doc.Selection, rule.Name)
	content := e.text(doc.Selection, rule.Content)

	chapter := NewChapter(name, url)
	return NewChapterDetail(chapter, content), nil
}

func (e extractor) ExtractBooks(rule ListRule, url string, html []byte) ([]book, error) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(html))
	if err != nil {
		return nil, NewHTMLParseError(err, html)
	}
	books := []book{}

	doc.Find(rule.List).Each(func(i int, elm *goquery.Selection) {
		book := e.extractBook(elm, rule.Book)
		if rule.Chapter.URL != "" {
			chapter := e.extractChapter(elm, rule.Chapter)
			book.chapter = &chapter
		}
		books = append(books, book)
	})
	return books, nil
}

func (e extractor) extractBook(elm *goquery.Selection, rule BookRule) book {
	name := e.text(elm, rule.Name)
	author := e.text(elm, rule.Author)
	introduce := e.text(elm, rule.Introduce)
	url := e.text(elm, rule.URL)

	return NewBook(name, url, author, introduce, nil)
}

func (e extractor) extractChapter(elm *goquery.Selection, rule ChapterRule) chapter {
	name := e.text(elm, rule.Name)
	url := e.text(elm, rule.URL)
	return NewChapter(name, url)
}

func (e extractor) text(doc *goquery.Selection, selector string) (text string) {
	if selector != "" {
		return doc.Find(selector).Text()
	}
	return ""
}

func (e extractor) attr(doc *goquery.Selection, selector, attr string) (value string, existed bool) {
	if selector != "" {
		return doc.Find(selector).Attr(attr)
	}
	return "", false
}
