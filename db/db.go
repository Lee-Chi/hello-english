package db

import (
	"context"

	"github.com/Lee-Chi/go-sdk/db/mongo"
)

var db *mongo.Database = nil

func Build(ctx context.Context, mongoURI string, dbName string) (err error) {
	db, err = mongo.NewDatabase(ctx, mongoURI, dbName)
	if err != nil {
		return
	}

	return
}

func Destroy(ctx context.Context) error {
	if err := db.Close(ctx); err != nil {
		return err
	}

	return nil
}

func Collection(name string) *mongo.Collection {
	return db.Collection(name)
}
