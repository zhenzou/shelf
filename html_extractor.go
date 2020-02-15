package shelf

import (
	"bytes"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func NewHTMLExtractor() Extractor {
	return &HTMLExtractor{}
}

type HTMLExtractor struct {
}

func (e HTMLExtractor) ExtractBook(rule BookRule, html []byte) (BookDetail, error) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(html))
	if err != nil {
		return BookDetail{}, NewHTMLParseError(err, html)
	}

	book := e.extractBook(doc.Selection, rule)
	detail := BookDetail{
		Book: book,
	}

	chapters := []Chapter{}

	doc.Find(rule.ChapterList.Rule).Each(func(i int, elm *goquery.Selection) {
		chapter := e.extractChapter(elm, rule.Chapter)
		chapters = append(chapters, chapter)
	})
	if len(chapters) > 0 {
		chapter := chapters[len(chapters)-1]
		detail.LatestChapter = &chapter
	}
	detail.Chapters = chapters
	return detail, nil
}

func (e *HTMLExtractor) ExtractChapter(rule ChapterRule, html []byte) (ChapterDetail, error) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(html))
	if err != nil {
		return ChapterDetail{}, NewHTMLParseError(err, html)
	}
	name := e.extractText(doc.Selection, rule.Name)
	content := e.extractText(doc.Selection, rule.Content)
	next := e.extractText(doc.Selection, rule.NextURL)

	chapter := Chapter{Name: name}
	return NewChapterDetail(chapter, content, next), nil
}

func (e *HTMLExtractor) ExtractBooks(rule ListRule, html []byte) ([]Book, error) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(html))
	if err != nil {
		return nil, NewHTMLParseError(err, html)
	}
	books := []Book{}

	doc.Find(rule.List.Rule).Each(func(i int, elm *goquery.Selection) {
		book := e.extractBook(elm, rule.Book)
		chapterRule := rule.Book.Chapter
		if chapterRule.URL.Rule != "" {
			chapter := e.extractChapter(elm, chapterRule)
			book.LatestChapter = &chapter
		}
		books = append(books, book)
	})
	return books, nil
}

func (e *HTMLExtractor) extractBook(elm *goquery.Selection, rule BookRule) Book {
	name := e.extractText(elm, rule.Name)
	author := e.extractText(elm, rule.Author)
	introduce := e.extractText(elm, rule.Introduce)
	url := e.extractText(elm, rule.URL)
	return NewBook(name, url, author, introduce, nil)
}

func (e *HTMLExtractor) extractChapter(selection *goquery.Selection, rule ChapterRule) Chapter {
	name := e.extractText(selection, rule.Name)
	url := e.extractText(selection, rule.URL)
	return NewChapter(name, url)
}

func (e *HTMLExtractor) extractText(selection *goquery.Selection, rule TextRule) (value string) {
	if rule.Rule == "" {
		return ""
	}
	elm := selection.Find(rule.Rule).First()
	if rule.Attr == "text" {
		value = elm.Text()
	} else {
		attr, _ := elm.Attr(rule.Attr)
		value = attr
	}
	value = FindRegMatched(rule.Regexp, value)
	return e.cleanText(rule.Clean, selection, value)
}

func (e *HTMLExtractor) cleanText(rule CleanRule, selection *goquery.Selection, text string) string {
	text = RemoveRegsMatched(rule.Regexps, text)
	if IsNotBlank(rule.Rules) {
		selectors := strings.Split(rule.Rules, ";")
		for _, selector := range selectors {
			if IsBlank(selector) {
				continue
			}
			text = strings.ReplaceAll(text, selection.Find(selector).Text(), "")
		}
	}
	return strings.TrimSpace(text)
}
