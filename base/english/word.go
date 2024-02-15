package english

import (
	"hello-english/db/model"

	"github.com/Lee-Chi/go-sdk/db/mongo"
)

const (
	Type_Normal string = ""
	Type_Sight  string = "sight"
)

var (
	Word_Key_ID      mongo.K = mongo.Key("_id")
	Word_Key_Letters mongo.K = mongo.Key("letters")
)

type Word struct {
	model.Word
}

type WordSight struct {
	model.WordSight
}
