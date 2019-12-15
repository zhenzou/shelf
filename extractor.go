package shelf

import (
	"bytes"

	"github.com/PuerkitoBio/goquery"
)

func newExtractor(url string, content []byte, rule BookRule) extractor {
	return extractor{
		content: content,
		rule:    rule,
	}
}

type extractor struct {
	url     string
	content []byte
	rule    BookRule
}

func (e extractor) ExtractBook() (book, error) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(e.content))
	if err != nil {
		return book{}, err
	}
	rule := e.rule
	name := doc.Find(rule.Name).Text()
	author := doc.Find(rule.Author).Text()
	introduce := doc.Find(rule.Introduce).Text()

	chapters := []chapter{}

	doc.Find(rule.ChapterURL).Each(func(i int, elm *goquery.Selection) {
		val, _ := elm.Attr("href")
		chapters = append(chapters, chapter{index: i, url: val})
	})

	return book{
		name:      name,
		url:       e.url,
		author:    author,
		chapters:  chapters,
		introduce: introduce,
	}, nil
}

func (e extractor) ExtractChapter() (chapter, error) {
	panic("implement me")
}
