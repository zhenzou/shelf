package shelf

type SourceRule struct {
	Name    string `json:"name"`
	BaseURL string `json:"base_url"`
	Type    string `json:"type"`
	Rules   struct {
		Index struct {
			URL  string   `json:"url"`
			Rule BookRule `json:"rule"`
		} `json:"index"`
		Search struct {
			URL  string   `json:"url"`
			Rule BookRule `json:"rule"`
		}
	} `json:"rules"`
}

type BookRule struct {
	Name           string `json:"name"`
	Author         string `json:"author"`
	Cover          string `json:"cover"`
	Class          string `json:"class"`
	Introduce      string `json:"introduce"`
	ChapterURL     string `json:"chapter_url"`
	ChapterName    string `json:"chapter_name"`
	ChapterContent string `json:"chapter_content"`
}
