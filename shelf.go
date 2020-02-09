package shelf

import (
	"context"
)

func New(executor Executor) Shelf {
	return &shelf{
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

func (s *shelf) Source(f func(args *SourceArgs)) (Source, bool) {
	args := SourceArgs{}
	f(&args)
	if args.Name != "" {
		source, ok := s.sources[args.Name]
		return source, ok
	}
	if args.URL != "" {
		for _, source := range s.sources {
			if source.Rule().BaseURL == args.URL {
				return source, true
			}
		}
	}
	return nil, false
}

func (s *shelf) Sources() map[string]Source {
	return s.sources
}

func (s *shelf) Search(ctx context.Context, name string) (map[string][]book, error) {
	ret := make(map[string][]book, len(s.sources))
	for name, source := range s.sources {
		books, err := source.Search(ctx, name)
		if err != nil {
			return nil, err
		}
		ret[name] = books
	}
	return ret, nil
}
