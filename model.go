package shelf

type SourceConfig struct {
	BookSourceName        string `json:"bookSourceName"`
	BookSourceGroup       string `json:"bookSourceGroup"`
	BookSourceType        string `json:"bookSourceType"`
	BookSourceURL         string `json:"bookSourceUrl"`
	LoginURL              string `json:"loginUrl"`
	RuleFindURL           string `json:"ruleFindUrl"`
	RuleFindList          string `json:"ruleFindList"`
	RuleFindName          string `json:"ruleFindName"`
	RuleFindAuthor        string `json:"ruleFindAuthor"`
	RuleFindKind          string `json:"ruleFindKind"`
	RuleFindLastChapter   string `json:"ruleFindLastChapter"`
	RuleFindIntroduce     string `json:"ruleFindIntroduce"`
	RuleFindCoverURL      string `json:"ruleFindCoverUrl"`
	RuleFindNoteURL       string `json:"ruleFindNoteUrl"`
	RuleSearchURL         string `json:"ruleSearchUrl"`
	RuleBookURLPattern    string `json:"ruleBookUrlPattern"`
	RuleSearchList        string `json:"ruleSearchList"`
	RuleSearchName        string `json:"ruleSearchName"`
	RuleSearchAuthor      string `json:"ruleSearchAuthor"`
	RuleSearchKind        string `json:"ruleSearchKind"`
	RuleSearchLastChapter string `json:"ruleSearchLastChapter"`
	RuleSearchIntroduce   string `json:"ruleSearchIntroduce"`
	RuleSearchCoverURL    string `json:"ruleSearchCoverUrl"`
	RuleSearchNoteURL     string `json:"ruleSearchNoteUrl"`
	RuleBookInfoInit      string `json:"ruleBookInfoInit"`
	RuleBookName          string `json:"ruleBookName"`
	RuleBookAuthor        string `json:"ruleBookAuthor"`
	RuleBookKind          string `json:"ruleBookKind"`
	RuleBookLastChapter   string `json:"ruleBookLastChapter"`
	RuleIntroduce         string `json:"ruleIntroduce"`
	RuleCoverURL          string `json:"ruleCoverUrl"`
	RuleChapterURL        string `json:"ruleChapterUrl"`
	RuleChapterURLNext    string `json:"ruleChapterUrlNext"`
	RuleChapterList       string `json:"ruleChapterList"`
	RuleChapterName       string `json:"ruleChapterName"`
	RuleContentURL        string `json:"ruleContentUrl"`
	RuleBookContent       string `json:"ruleBookContent"`
	RuleContentURLNext    string `json:"ruleContentUrlNext"`
	HTTPUserAgent         string `json:"httpUserAgent"`
	SerialNumber          int    `json:"serialNumber"`
	Weight                int    `json:"weight"`
	Enable                bool   `json:"enable"`
}
