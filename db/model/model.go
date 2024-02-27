package model

import "github.com/Lee-Chi/go-sdk/db/mongo"

type K struct {
	WordBase     KeyWordBase
	WordSight    KeyWordSight
	WordBookmark KeyWordBookmark
}

var Key K

func init() {
	Key = K{
		WordBase: KeyWordBase{
			ID:      mongo.Key("_id"),
			Letters: mongo.Key("letters"),
		},
		WordSight: KeyWordSight{
			ID:      mongo.Key("_id"),
			Letters: mongo.Key("letters"),
		},
		WordBookmark: KeyWordBookmark{
			ID:      mongo.Key("_id"),
			Letters: mongo.Key("letters"),
			UserID:  mongo.Key("userId"),
			Index:   mongo.Key("index"),
		},
	}
}
