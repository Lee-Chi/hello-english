package model

import "time"

const CUserAccount string = "user_account"

type UserAccount struct {
	Account   string    `bson:"account"`
	Password  string    `bson:"password"`
	CreatedAt time.Time `bson:"createdAt"`
	ExpiredAt time.Time `bson:"expiredAt"`
}

type UserSession struct {
	UserId    string    `bson:"userId"`
	Token     string    `bson:"token"`
	CreatedAt time.Time `bson:"createdAt"`
	ExpiredAt time.Time `bson:"expiredAt"`
}
