package shelf

type SourceRule struct {
	Enable   bool
	Encoding string
	Name     string
	BaseURL  string
	Tags     []string
	Order    int
	Rules struct {
		Index   ListRule
		Search  ListRule
		Book    BookRule
		Chapter ChapterRule
	}
}

type ListRule struct {
	URL     string
	List    TextRule
	Book    BookRule
	Chapter ChapterRule
}

type BookRule struct {
	Name      TextRule
	Author    TextRule
	Cover     TextRule
	Class     TextRule
	Introduce TextRule
	URL       TextRule
	Update    TimeRule
	Chapter   ChapterRule
}

type ChapterRule struct {
	List    TextRule
	Name    TextRule
	URL     TextRule
	Content TextRule
}

type TextRule struct {
	Selector string
	Attr     string
	Regexp   string
	Remove   string
}

type TimeRule struct {
	Selector string
	Attr     string
	Format   string
}

type Args struct {
	Name string
	Page uint64
}
