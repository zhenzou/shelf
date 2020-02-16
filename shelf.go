package shelf

import (
	"context"
	"strings"
	"sync"
)

func New(executor Executor, extractor Extractor) *Shelf {
	return &Shelf{
		sources:   map[string]Source{},
		executor:  executor,
		extractor: extractor,
	}
}

type Shelf struct {
	sync.RWMutex
	sources   map[string]Source
	executor  Executor
	extractor Extractor
}

func (s *Shelf) AddSource(config SourceConfig) {
	s.Lock()
	defer s.Unlock()
	s.sources[config.Name] = NewSource(config, s.executor, s.extractor)
}

func (s *Shelf) RemoveSource(name string) {
	s.Lock()
	defer s.Unlock()
	delete(s.sources, name)
}

func (s *Shelf) Source(urlOrName string) (Source, bool) {
	s.RLock()
	defer s.RUnlock()
	source, ok := s.sources[urlOrName]
	if ok {
		return source, true
	}
	for _, source := range s.sources {
		if strings.Contains(source.Config().BaseURL, urlOrName) {
			return source, true
		}
	}
	return nil, false
}

func (s *Shelf) Sources() map[string]Source {
	return s.sources
}

func (s *Shelf) Search(ctx context.Context, kw string) (map[string][]Book, error) {
	snap := map[string]Source{}
	s.RLock()
	for name, source := range s.sources {
		snap[name] = source
	}
	s.RUnlock()

	ret := make(map[string][]Book, len(snap))
	for name, source := range snap {
		books, err := source.Search(ctx, kw)
		if err != nil {
			return nil, err
		}
		ret[name] = books
	}
	return ret, nil
}
