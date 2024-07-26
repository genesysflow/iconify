package iconify

type Collection struct {
	Key    string `json:"key"`
	Name   string `json:"name"`
	Total  int    `json:"total"`
	Author struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"author"`
	Category string `json:"category"`
}

type IconCollection struct {
	Prefix        string              `json:"prefix"`
	Total         int                 `json:"total"`
	Title         string              `json:"title"`
	Uncategorized []string            `json:"uncategorized"`
	Categories    map[string][]string `json:"categories"`
}
