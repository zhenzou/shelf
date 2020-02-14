package shelf

type SourceRule struct {
	Name     string
	Enable   bool
	Encoding string
	BaseURL  string
	Tags     []string
	Order    int
	Rules    struct {
		Find    ListRule    // 发现
		Search  ListRule    // 搜索
		Book    BookRule    // 书籍详情
		Chapter ChapterRule // 章节详情
	}
}

type ListRule struct {
	URL     string      // URL模版
	List    ElementRule //
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
	NextURL TextRule
}

type ElementRule struct {
	Rule string
}

type TextRule struct {
	Rule   string
	Attr   string
	Regexp string
	Clean  CleanRule
}

type TimeRule struct {
	Rule   string
	Attr   string
	Format string
}

type CleanRule struct {
	Regexps string
	Rules   string
}

type Args struct {
	Name string
	Page uint64
}
