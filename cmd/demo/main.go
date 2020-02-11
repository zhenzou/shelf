package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/zhenzou/shelf"
)

var rule = shelf.SourceRule{
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
			Selector: "body > div:nth-child(4) > div.list > div > ul > li",
		},
		Book: shelf.BookRule{
			Name: shelf.TextRule{
				Selector: "a",
				Attr:     "text",
			},
			Author: shelf.TextRule{
				Selector: "div.s",
				Attr:     "text",
				Regexp:   ".*作者：(.+)大小",
			},
			Introduce: shelf.TextRule{
				Selector: "div.u",
				Attr:     "text",
			},
			Chapter: shelf.ChapterRule{
				Name: shelf.TextRule{
					Selector: "o > a",
					Attr:     "text",
				},
				URL: shelf.TextRule{
					Selector: "o > a",
					Attr:     "href",
				},
			},
		},
	}
	rule.Rules.Book = shelf.BookRule{
		Name: shelf.TextRule{
			Selector: "#info > div.hh",
			Attr:     "text",
		},
		Author: shelf.TextRule{
			Selector: "#conml > table > tbody > tr:nth-child(1) > td > div.bcont",
			Attr:     "text",
			Regexp:   `.*作者：(.+)状态`,
		},
		Cover: shelf.TextRule{
			Selector: "#conml > table > tbody > tr:nth-child(1) > td > div.bcont > img",
			Attr:     "src",
		},
		Introduce: shelf.TextRule{
			Selector: "#conml > table > tbody > tr:nth-child(1) > td > div.bcont",
			Attr:     "text",
			Regexp:   ".*简介：(.+\n)",
		},
		Chapter: shelf.ChapterRule{
			List: shelf.TextRule{
				Selector: "#list > dl > dd",
			},
			Name: shelf.TextRule{
				Selector: "a",
				Attr:     "text",
				Remove:   "[www.126shu.co]",
			},
			URL: shelf.TextRule{
				Selector: "a",
				Attr:     "href",
			},
		},
	}

	rule.Rules.Chapter = shelf.ChapterRule{
		Name: shelf.TextRule{
			Selector: "#info > div.hh",
			Attr:     "text",
		},
		Content: shelf.TextRule{
			Selector: "#content",
			Attr:     "text",
			Remove:   "www.126shu.co",
		},
	}
}

func GetBook(source shelf.Source) {
	book, err := source.GetBookDetail(context.Background(), "https://www.126shu.co/90497/")
	if err != nil {
		println("err:", err.Error())
	} else {
		println(book.Name)
		println(book.Author)
		println(book.Introduce)

		chapters := book.Chapters
		for _, chapter := range chapters {
			println(fmt.Sprintf("chapter:%s url:%s", chapter.Name, chapter.URL))

			detail, err := source.GetChapterDetail(context.Background(), chapter.URL)
			if err != nil {
				println("get book datail err:", err.Error())
				return
			}
			println(detail.Content)
			return
		}
	}
}

func Search(source shelf.Source) {
	books, err := source.Search(context.Background(), "斗罗大陆")
	if err != nil {
		println("err:", err.Error())
	} else {
		for _, book := range books {
			println(fmt.Sprintf("%s %s %s", book.Name, book.Author, book.Introduce))
		}
	}
}

func main() {
	s := shelf.New(shelf.NewExecutor(http.DefaultClient))
	s.AddSource(rule, shelf.DefaultExtractor())

	source, ok := s.Source(shelf.WithName("奇书网"))
	if ok {
		GetBook(source)
		//Search(source)
	}
}
