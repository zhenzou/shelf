package shelf

import (
	"context"
	"strings"
)

func New(executor Executor) shelf {
	return shelf{
		sources:  map[string]Source{},
		executor: executor,
	}
}

type shelf struct {
	sources  map[string]Source
	executor Executor
}

func (s *shelf) AddSource(rule SourceRule, extractor Extractor) {
	s.sources[rule.Name] = NewSource(rule, s.executor, extractor)
}

func (s *shelf) Source(urlOrName string) (Source, bool) {
	source, ok := s.sources[urlOrName]
	if ok {
		return source, true
	}
	for _, source := range s.sources {
		if strings.Contains(source.Rule().BaseURL, urlOrName) {
			return source, true
		}
	}
	return nil, false
}

func (s *shelf) Sources() map[string]Source {
	return s.sources
}

func (s *shelf) Search(ctx context.Context, name string) (map[string][]Book, error) {
	ret := make(map[string][]Book, len(s.sources))
	for name, source := range s.sources {
		books, err := source.Search(ctx, name)
		if err != nil {
			return nil, err
		}
		ret[name] = books
	}
	return ret, nil
}
