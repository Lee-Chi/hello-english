package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Definition struct {
	PartOfSpeech string `bson:"partOfSpeech"`
	Translation  string `bson:"translation"`
}

type Word struct {
	Letters     string       `bson:"letters"`
	Definitions []Definition `bson:"definitions"`
}

type WordExplained struct {
	Letters      string   `bson:"letters"`
	PartOfSpeech string   `bson:"partOfSpeech"`
	Translation  string   `bson:"translation"`
	Sentences    []string `bson:"sentences"`
}

const CWordBase = "word_base"

type WordBase struct {
	Word  `bson:"-,inline"`
	Level int `bson:"level"`
}

const CWordSight = "word_sight"

type WordSight struct {
	Letters string `bson:"letters"`
	Level   string `bson:"level"`
}

const CWordBookmark = "word_bookmark"

type WordBookmark struct {
	WordExplained `bson:"-,inline"`
	UserId        primitive.ObjectID `bson:"userId"`
	Index         int                `bson:"index"`
	CreatedAt     time.Time          `bson:"createdAt"`
}
