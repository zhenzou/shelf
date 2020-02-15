package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/zhenzou/shelf"
)

var rule = shelf.SourceConfig{
	Name:     "奇书网",
	BaseURL:  "https://www.126shu.co",
	Tags:     []string{"网络小说"},
	Order:    0,
	Encoding: "gbk",
	Enable:   true,
}

func init() {
	rule.Rules.Search = shelf.ListRule{
		URL: "https://www.126shu.co/modules/article/search.php?s=12622474051500695548&orderby=1&show=title,bname,zuozhe,smalltext&myorder=1&searchkey=${name}",
		List: shelf.ElementRule{
			Rule: "body > div:nth-child(4) > div.list > div > ul > li",
		},
		Book: shelf.BookRule{
			Name: shelf.TextRule{
				Rule: "a",
				Attr: "text",
			},
			Author: shelf.TextRule{
				Rule:   "div.s",
				Attr:   "text",
				Regexp: ".*作者：(.+)大小",
			},
			URL: shelf.TextRule{
				Rule: "a",
				Attr: "href",
			},
			Introduce: shelf.TextRule{
				Rule: "div.u",
				Attr: "text",
				Clean: shelf.CleanRule{
					Rules: "o > a",
				},
			},
			Chapter: shelf.ChapterRule{
				Name: shelf.TextRule{
					Rule: "o > a",
					Attr: "text",
				},
				URL: shelf.TextRule{
					Rule: "o > a",
					Attr: "href",
				},
			},
		},
	}
	rule.Rules.Book = shelf.BookRule{
		Name: shelf.TextRule{
			Rule: "#info > div.hh",
			Attr: "text",
		},
		Author: shelf.TextRule{
			Rule:   "#conml > table > tbody > tr:nth-child(1) > td > div.bcont",
			Attr:   "text",
			Regexp: `.*作者：(.+)状态`,
		},
		Cover: shelf.TextRule{
			Rule: "#conml > table > tbody > tr:nth-child(1) > td > div.bcont > img",
			Attr: "src",
		},
		Introduce: shelf.TextRule{
			Rule:   "#conml > table > tbody > tr:nth-child(1) > td > div.bcont",
			Attr:   "text",
			Regexp: ".*简介：(.+\n)",
		},
		URL: shelf.TextRule{
			Rule: "body > div:nth-child(4) > div.show > div:nth-child(4) > div.showDown > ul > li:nth-child(1) > a",
			Attr: "href",
		},
		ChapterList: shelf.ElementRule{
			Rule: "#list > dl > dd",
		},
		Chapter: shelf.ChapterRule{
			Name: shelf.TextRule{
				Rule: "a",
				Attr: "text",
				Clean: shelf.CleanRule{
					Regexps: "[www.126shu.co]",
				},
			},
			URL: shelf.TextRule{
				Rule: "a",
				Attr: "href",
			},
		},
	}

	rule.Rules.Chapter = shelf.ChapterRule{
		Name: shelf.TextRule{
			Rule: "#info > div.hh",
			Attr: "text",
		},
		Content: shelf.TextRule{
			Rule: "#content",
			Attr: "text",
			Clean: shelf.CleanRule{
				Regexps: "www.126shu.co;\\s*-----网友请提示:长时间阅读请注意眼睛的休息。：\\s*; \\S*----这是华丽的分割线---</i>\\s*",
				Rules:   "div.zjtj;div.zjxs",
			},
		},
	}
}

func GetBook(source shelf.Source, url string) {
	println("book url:", url)
	book, err := source.GetBookDetail(context.Background(), url)
	if err != nil {
		println("err:", err.Error())
	} else {
		println(book.Name)
		println(book.Author)
		println(book.Introduce)

		chapters := book.Chapters
		for _, chapter := range chapters {
			println(fmt.Sprintf("chapter:%s url:%s", chapter.Name, chapter.URL))
		}
	}
}

func Search(source shelf.Source) {
	books, err := source.Search(context.Background(), "斗罗大陆")
	if err != nil {
		println("err:", err.Error())
	} else {
		for i, book := range books {
			bs, _ := json.Marshal(book)
			println("book:", string(bs))
			if i == 0 {
				GetBook(source, book.URL)
			}
		}
	}
}

func main() {
	s := shelf.New(shelf.NewExecutor(http.DefaultClient))
	s.AddSource(rule, shelf.NewHTMLExtractor())

	source, ok := s.Source("奇书网")
	if ok {
		//GetBook(source, "https://www.126shu.co/90497/")
		Search(source)
	}
}
