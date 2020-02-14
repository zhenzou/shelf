package shelf

type ExtractorBuilder = func(content []byte) (Extractor, error)

type ExtractorFactory struct {
	builders map[string]ExtractorBuilder
}

func (f *ExtractorFactory) AddBuilder(name string, extractor ExtractorBuilder) {
	f.builders[name] = extractor
}

func (f *ExtractorFactory) NewExtractor(name string, content []byte) (Extractor, error) {
	builder, ok := f.builders[name]
	if ok {
		return builder(content)
	} else {
		return nil, NewUnsupportedEncodingError(nil, name)
	}
}
