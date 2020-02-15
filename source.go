package shelf

import (
	"context"
	"strings"
)

func NewSource(rule SourceConfig, executor Executor, builder Extractor) Source {
	return &source{
		rule:      rule,
		executor:  executor,
		extractor: builder,
	}
}

type source struct {
	rule      SourceConfig
	executor  Executor
	extractor Extractor
}

func (s *source) Name() string {
	return s.rule.Name
}

func (s *source) Rule() SourceConfig {
	return s.rule
}

func (s *source) GetBookDetail(ctx context.Context, url string) (BookDetail, error) {
	book, err := s.getBookDetail(ctx, url)
	if err != nil {
		return book, err
	}
	if s.rule.Rules.Book.URL.Rule != "" {
		detail2, err := s.getBookDetail(ctx, book.URL)
		if err != nil {
			return detail2, err
		}
		book = s.mergeBookDetail(book, detail2)
	}
	return book, nil
}

func (s *source) mergeBookDetail(detail1 BookDetail, detail2 BookDetail) BookDetail {
	book := BookDetail{
		Book: Book{
			Name:             WithDefault(detail1.Name, detail2.Name),
			URL:              WithDefault(detail1.URL, detail2.URL),
			Author:           WithDefault(detail1.Author, detail2.Author),
			Introduce:        WithDefault(detail1.Introduce, detail2.Introduce),
			LatestChapter:    detail1.LatestChapter,
			LatestUpdateTime: detail1.LatestUpdateTime,
		},
		Chapters: detail1.Chapters,
	}
	if book.LatestChapter == nil {
		book.LatestChapter = detail2.LatestChapter
	}
	if book.LatestUpdateTime == nil {
		book.LatestUpdateTime = detail2.LatestUpdateTime
	}
	if len(book.Chapters) == 0 {
		book.Chapters = detail2.Chapters
	}
	return book
}

func (s *source) getBookDetail(ctx context.Context, url string) (BookDetail, error) {
	url = s.buildFullURL(url)
	response, err := s.executor.Exec(ctx, Request{Method: "GET", URL: url})
	if err != nil {
		return BookDetail{}, err
	}
	response.Data, err = decode(s.rule.Encoding, response.Data)
	if err != nil {
		return BookDetail{}, err
	}
	book, err := s.extractor.ExtractBook(s.rule.Rules.Book, response.Data)
	if err != nil {
		return BookDetail{}, err
	}

	return book, err
}

func (s *source) buildFullURL(url string) string {
	if s.isFullURL(url) {
		return url
	}

	baseURL := s.rule.BaseURL
	if strings.HasPrefix(url, "/") {
		return baseURL + url
	}
	return baseURL + "/" + url
}

func (s *source) isFullURL(url string) bool {
	return strings.HasPrefix(url, "https://") || strings.HasPrefix(url, "http://")
}

func (s *source) GetChapterDetail(ctx context.Context, url string) (ChapterDetail, error) {
	url = s.buildFullURL(url)
	response, err := s.executor.Exec(ctx, Request{Method: "GET", URL: url})
	if err != nil {
		return ChapterDetail{}, err
	}
	response.Data, err = decode(s.rule.Encoding, response.Data)
	if err != nil {
		return ChapterDetail{}, err
	}
	detail, err := s.extractor.ExtractChapter(s.rule.Rules.Chapter, response.Data)
	if err != nil {
		return ChapterDetail{}, err
	}
	detail.URL = url
	detail.NextURL = s.buildFullURL(detail.NextURL)
	return detail, err
}

func (s *source) Search(ctx context.Context, name string) ([]Book, error) {
	url := s.rule.Rules.Search.URL
	if url == "" {
		return nil, nil
	}
	bytes, err := encode(s.rule.Encoding, []byte(name))
	if err != nil {
		return nil, err
	}
	args := Args{Name: string(bytes), Page: 1}
	response, err := s.executor.Exec(ctx, Request{"GET", url, args})
	if err != nil {
		return nil, err
	}
	response.Data, err = decode(s.rule.Encoding, response.Data)
	if err != nil {
		return nil, err
	}
	return s.extractor.ExtractBooks(s.rule.Rules.Search, response.Data)
}
