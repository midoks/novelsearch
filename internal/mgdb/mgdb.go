package mgdb

import (
	"context"
	"fmt"
	"sync"

	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/midoks/novelsearch/internal/conf"
)

var (
	link string
	err  error

	ctx    context.Context
	client *qmgo.Client
	db     *qmgo.Database

	mutex sync.RWMutex
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
	link = "mongodb://" + conf.Mongodb.Addr

	ctx = context.Background()
	client, err = qmgo.NewClient(ctx, &qmgo.Config{Uri: link})

	if err != nil {
		return fmt.Errorf("mgdb client connect error: %v", err)
	}

	db = client.Database(conf.Mongodb.Db)

	return nil
}
