package shelf

type SourceConfig struct {
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
	List    ElementRule // List元素
	Book    BookRule    // BOOK
	Chapter ChapterRule
}

type BookRule struct {
	Name        TextRule
	Author      TextRule
	Cover       TextRule
	Class       TextRule
	Introduce   TextRule
	URL         TextRule    // 本小说的URL，在列表页使用;有些站点，详情页的需要再次点击才能进入章节列表页
	Update      TimeRule    // 更新时间，在详情页使用
	ChapterList ElementRule // 章节列表
	Chapter     ChapterRule // 章节信息，在列表页使用，包括书籍列表或者章节列表
}

type ChapterRule struct {
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
