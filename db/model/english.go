package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const CWord = "word"

type Definition struct {
	PartOfSpeech string `bson:"partOfSpeech"`
	Translation  string `bson:"translation"`
}

type Word struct {
	Letters     string       `bson:"letters"`
	Definitions []Definition `bson:"definitions"`
	Level       int          `bson:"level"`
}

const CWordSight = "word_sight"

type WordSight struct {
	Letters string `bson:"letters"`
	Level   string `bson:"level"`
}

const CWordToday = "word_today"

type WordToday struct {
	WordID      primitive.ObjectID `bson:"wordId"`
	WordSightID primitive.ObjectID `bson:"wordSightId"`
	Date        string             `bson:"date"`
}
