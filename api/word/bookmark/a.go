package bookmark

type Group struct {
}

var g Group

type Word struct {
	Letters      string   `json:"letters"`
	PartOfSpeech string   `json:"partOfSpeech"`
	Translation  string   `json:"translation"`
	Sentences    []string `json:"sentences"`
}
