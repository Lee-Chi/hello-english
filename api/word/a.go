package word

import "hello-english/api/word/bookmark"

type Group struct {
	Bookmark bookmark.Group
}

var g Group

type Definition struct {
	PartOfSpeech string `json:"partOfSpeech"`
	Translation  string `json:"translation"`
}

type Word struct {
	ID          string       `json:"id"`
	Letters     string       `json:"letters"`
	Definitions []Definition `json:"definitions"`
}
