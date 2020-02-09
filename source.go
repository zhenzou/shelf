package shelf

import (
	"context"
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
	return s.extractor.ExtractBook(s.rule.Rules.Book, url, response.Data)
}

func (s *source) GetChapterDetail(ctx context.Context, url string) (chapterDetail, error) {
	response, err := s.executor.Exec(ctx, Request{Method: "GET", URL: url})
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
	args := Args{Name: name, Page: 1}
	response, err := s.executor.Exec(ctx, Request{"GET", url, args})
	if err != nil {
		return nil, err
	}
	return s.extractor.ExtractBooks(s.rule.Rules.Search, url, response.Data)
}
