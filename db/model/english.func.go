package model

import (
	"github.com/Lee-Chi/go-sdk/db/mongo"
)

const (
	Type_Base  string = ""
	Type_Sight string = "sight"
)

type KeyWordBase struct {
	ID      mongo.K
	Letters mongo.K
}

type KeyWordSight struct {
	ID      mongo.K
	Letters mongo.K
}

type KeyWordBookmark struct {
	ID      mongo.K
	Letters mongo.K
	UserID  mongo.K
	Index   mongo.K
}
