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

func (e extractor) ExtractBook(rule BookRule, url string, html []byte) (BookDetail, error) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(html))
	if err != nil {
		return BookDetail{}, NewHTMLParseError(err, html)
	}
	name := e.extractText(doc.Selection, rule.Name)
	author := e.extractText(doc.Selection, rule.Author)
	introduce := e.extractText(doc.Selection, rule.Introduce)

	chapters := []Chapter{}

	doc.Find(rule.Chapter.List.Rule).Each(func(i int, elm *goquery.Selection) {
		chapter := e.extractChapter(elm, rule.Chapter)
		chapters = append(chapters, chapter)
	})
	chapter := chapters[len(chapters)-1]
	return NewBookDetail(NewBook(name, url, author, introduce, &chapter), chapters), nil
}

func (e *extractor) ExtractChapter(rule ChapterRule, url string, html []byte) (ChapterDetail, error) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(html))
	if err != nil {
		return ChapterDetail{}, NewHTMLParseError(err, html)
	}
	name := e.extractText(doc.Selection, rule.Name)
	content := e.extractText(doc.Selection, rule.Content)
	next := e.extractText(doc.Selection, rule.NextURL)

	chapter := NewChapter(name, url)
	return NewChapterDetail(chapter, content, next), nil
}

func (e *extractor) ExtractBooks(rule ListRule, url string, html []byte) ([]Book, error) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(html))
	if err != nil {
		return nil, NewHTMLParseError(err, html)
	}
	books := []Book{}

	doc.Find(rule.List.Rule).Each(func(i int, elm *goquery.Selection) {
		book := e.extractBook(elm, rule.Book)
		if rule.Chapter.URL.Rule != "" {
			chapter := e.extractChapter(elm, rule.Chapter)
			book.LatestChapter = &chapter
		}
		books = append(books, book)
	})
	return books, nil
}

func (e *extractor) extractBook(elm *goquery.Selection, rule BookRule) Book {
	name := e.extractText(elm, rule.Name)
	author := e.extractText(elm, rule.Author)
	introduce := e.extractText(elm, rule.Introduce)
	url := e.extractText(elm, rule.URL)
	return NewBook(name, url, author, introduce, nil)
}

func (e *extractor) extractChapter(selection *goquery.Selection, rule ChapterRule) Chapter {
	name := e.extractText(selection, rule.Name)
	url := e.extractText(selection, rule.URL)
	return NewChapter(name, url)
}

func (e *extractor) extractText(selection *goquery.Selection, rule TextRule) (value string) {
	if rule.Rule != "" {
		elm := selection.Find(rule.Rule).First()
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

	return e.cleanText(rule.Clean, selection, value)
}

func (e *extractor) cleanText(rule CleanRule, selection *goquery.Selection, text string) string {

	if IsNotBlank(rule.Regexps) {
		patterns := Split(rule.Regexps, ';')
		for _, pattern := range patterns {
			reg, ok := e.getOrCreateReg(pattern)
			if !ok {
				continue
			}
			text = reg.ReplaceAllString(text, "")
		}
	}
	if IsNotBlank(rule.Rules) {
		selectors := strings.Split(rule.Rules, ";")
		for _, selector := range selectors {
			if IsBlank(selector) {
				continue
			}
			text = strings.ReplaceAll(text, selection.Find(selector).Text(), "")
		}
	}

	return text
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
