package shelf

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type extract struct {
}

func (e *extract) ExtractBook(content string, rule BookRule) (Book, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(content))
	if err != nil {
		return nil, err
	}
	name := doc.Find(rule.Name).Text()
	author := doc.Find(rule.Author).Text()
	introduce := doc.Find(rule.Introduce).Text()

	chapters := []Chapter{}

	doc.Find(rule.ChapterURL).Each(func(i int, elm *goquery.Selection) {
		val, _ := elm.Attr("href")
		chapters = append(chapters, &chapter{index: i, url: val})
	})

	return &book{
		source:    nil,
		extractor: e,
		name:      name,
		url:       "",
		author:    author,
		chapters:  chapters,
		introduce: introduce,
	}, nil
}

func (e *extract) ExtractChapter(content string, rule BookRule) (Book, error) {
	panic("implement me")
}
