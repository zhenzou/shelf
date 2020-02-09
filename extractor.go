package shelf

import (
	"bytes"
	"regexp"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

func DefaultExtractor() Extractor {
	return &extractor{}
}

type extractor struct {
	regCache sync.Map
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

	doc.Find(rule.Chapter.List.Selector).Each(func(i int, elm *goquery.Selection) {
		chapter := e.extractChapter(elm, rule.Chapter)
		chapters = append(chapters, chapter)
	})
	chapter := chapters[len(chapters)-1]
	return NewBookDetail(NewBook(name, url, author, introduce, &chapter), chapters), nil
}

func (e *extractor) ExtractChapter(rule ChapterRule, url string, html []byte) (chapterDetail, error) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(html))
	if err != nil {
		return chapterDetail{}, NewHTMLParseError(err, html)
	}
	name := e.text(doc.Selection, rule.Name)
	content := e.text(doc.Selection, rule.Content)

	chapter := NewChapter(name, url)
	return NewChapterDetail(chapter, content), nil
}

func (e *extractor) ExtractBooks(rule ListRule, url string, html []byte) ([]book, error) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(html))
	if err != nil {
		return nil, NewHTMLParseError(err, html)
	}
	books := []book{}

	doc.Find(rule.List.Selector).Each(func(i int, elm *goquery.Selection) {
		book := e.extractBook(elm, rule.Book)
		if rule.Chapter.URL.Selector != "" {
			chapter := e.extractChapter(elm, rule.Chapter)
			book.Chapter = &chapter
		}
		books = append(books, book)
	})
	return books, nil
}

func (e *extractor) extractBook(elm *goquery.Selection, rule BookRule) book {
	name := e.text(elm, rule.Name)
	author := e.text(elm, rule.Author)
	introduce := e.text(elm, rule.Introduce)
	url := e.text(elm, rule.URL)
	return NewBook(name, url, author, introduce, nil)
}

func (e *extractor) extractChapter(selection *goquery.Selection, rule ChapterRule) chapter {
	name := e.text(selection, rule.Name)
	url := e.text(selection, rule.URL)
	return NewChapter(name, url)
}

func (e *extractor) text(selection *goquery.Selection, rule TextRule) (value string) {
	if rule.Selector != "" {
		elm := selection.Find(rule.Selector).First()
		if rule.Attr == "text" {
			value = elm.Text()
		} else {
			attr, _ := elm.Attr(rule.Attr)
			value = attr
		}
	}
	reg, ok := e.getOrCreateReg(rule.Regexp)
	if ok {
		matched, ok := e.findMatchedString(reg, value)
		if ok {
			value = matched
		}
	}

	if rule.Remove != "" {
		value = strings.ReplaceAll(value, rule.Remove, "")
	}
	return strings.TrimSpace(value)
}

func (e *extractor) findMatchedString(reg *regexp.Regexp, str string) (string, bool) {
	match := reg.FindSubmatch([]byte(str))
	if len(match) > 1 {
		return string(match[1]), true
	}
	return "", false
}

func (e *extractor) getOrCreateReg(pattern string) (reg *regexp.Regexp, ok bool) {
	if IsBlank(pattern) {
		return nil, false
	}

	obj, ok := e.regCache.Load(pattern)
	if !ok {
		reg, err := regexp.Compile(pattern)
		if err != nil {
			// TODO LOG
			return nil, false
		}
		e.regCache.Store(pattern, reg)
		return reg, true
	} else {
		reg := obj.(*regexp.Regexp)
		return reg, true
	}
}
