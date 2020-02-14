package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/zhenzou/shelf"
)

var rule = shelf.SourceRule{
	Name:    "笔趣阁",
	BaseURL: "https://m.biqugetv.com",
	Tags:    []string{"网络小说"},
	Order:   0,
	Enable:  true,
}

func init() {

	rule.Rules.Chapter = shelf.ChapterRule{
		Name: shelf.TextRule{
			Rule: "#chaptercontent > div > span",
			Attr: "text",
			Clean: shelf.CleanRule{
				Regexps: ".点击下一页继续阅读",
				Rules:   "",
			},
		},
		Content: shelf.TextRule{
			Rule: "p.Readpage:nth-child(13) > a:nth-child(3)",
			Attr: "text",
		},
		NextURL: shelf.TextRule{
			Rule: "#pt_next",
			Attr: "href",
		},
	}
}

func Iterate(source shelf.Source) {

	url := "https://m.biqugetv.com/35_35756/22696509.html"
	for shelf.IsNotBlank(url) {
		chapter, err := source.GetChapterDetail(context.Background(), url)
		if err != nil {
			println("err:", err.Error())
		} else {
			println(fmt.Sprintf("%s:%s", chapter.Name, url))
			url = chapter.NextURL
		}
	}
}

func main() {
	s := shelf.New(shelf.NewExecutor(http.DefaultClient))
	s.AddSource(rule, shelf.NewHTMLExtractor())

	source, ok := s.Source("笔趣阁")
	if ok {
		Iterate(source)
	}
}
