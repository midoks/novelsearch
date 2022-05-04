package mgdb

import (
	"context"

	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/midoks/novelsearch/internal/conf"
)

var (
	err        error
	ctx        context.Context
	client     *qmgo.Client
	db         *qmgo.Database
	collection *qmgo.Collection

	cliContent *qmgo.QmgoClient
)

type (
	// M is an alias of bson.M
	M = bson.M
	// A is an alias of bson.A
	A = bson.A
	// D is an alias of bson.D
	D = bson.D
	// E is an alias of bson.E
	E = bson.E
)

func Init() error {
	link := "mongodb://" + conf.Mongodb.Addr

	ctx = context.Background()
	client, err = qmgo.NewClient(ctx, &qmgo.Config{Uri: link})

	if err != nil {
		return err
	}

	db = client.Database(conf.Mongodb.Db)
	collection = db.Collection("list")

	cliContent, err = qmgo.Open(ctx, &qmgo.Config{Uri: link, Database: conf.Mongodb.Db, Coll: "list"})
	if err != nil {
		return err
	}

	return nil
}
