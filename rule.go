package shelf

type SourceRule struct {
	Name    string
	BaseURL string
	Tags    []string
	Order   int
	Enable  bool
	Rules   struct {
		Index   ListRule
		Search  ListRule
		Book    BookRule
		Chapter ChapterRule
	}
}

type ListRule struct {
	URL     string
	List    string
	Book    BookRule
	Chapter ChapterRule
}

type BookRule struct {
	Name      string
	Author    string
	Cover     string
	Class     string
	Introduce string
	URL       string
	Chapter   ChapterRule
}

type ChapterRule struct {
	List    string
	Name    string
	URL     string
	Content string
}

type Args struct {
	Name string
	Page int64
}
