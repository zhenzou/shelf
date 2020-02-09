package shelf

import (
	"context"
	"strings"
)

func NewSource(rule SourceRule, executor Executor, builder Extractor) Source {
	return &source{
		rule:      rule,
		executor:  executor,
		extractor: builder,
	}
}

type source struct {
	rule      SourceRule
	executor  Executor
	extractor Extractor
}

func (s *source) Name() string {
	return s.rule.Name
}

func (s *source) Rule() SourceRule {
	return s.rule
}

func (s *source) GetBookDetail(ctx context.Context, url string) (bookDetail, error) {
	response, err := s.executor.Exec(ctx, Request{Method: "GET", URL: url})
	if err != nil {
		return bookDetail{}, err
	}
	response.Data, err = decode(s.rule.Encoding, response.Data)
	if err != nil {
		return bookDetail{}, err
	}
	book, err := s.extractor.ExtractBook(s.rule.Rules.Book, url, response.Data)
	if err != nil {
		return bookDetail{}, err
	}

	for i, chapter := range book.Chapters {
		if !s.isFullURL(chapter.URL) {
			chapter.URL = s.buildFullURL(s.rule.BaseURL, chapter.URL)
			book.Chapters[i] = chapter
		}
	}
	return book, err
}

func (s *source) buildFullURL(baseURL, url string) string {
	if strings.HasPrefix(url, "/") {
		return baseURL + url
	}
	return baseURL + "/" + url
}

func (s *source) isFullURL(url string) bool {
	return strings.HasPrefix(url, "https://") || strings.HasPrefix(url, "http://")
}

func (s *source) GetChapterDetail(ctx context.Context, url string) (chapterDetail, error) {
	response, err := s.executor.Exec(ctx, Request{Method: "GET", URL: url})
	if err != nil {
		return chapterDetail{}, err
	}
	response.Data, err = decode(s.rule.Encoding, response.Data)
	if err != nil {
		return chapterDetail{}, err
	}
	return s.extractor.ExtractChapter(s.rule.Rules.Chapter, url, response.Data)
}

func (s *source) Search(ctx context.Context, name string) ([]book, error) {
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
	return s.extractor.ExtractBooks(s.rule.Rules.Search, url, response.Data)
}
