package mgdb

import (
	"fmt"
	"strings"
	"time"

	"github.com/qiniu/qmgo"
	"github.com/qiniu/qmgo/operator"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/midoks/novelsearch/internal/conf"
)

type Novel struct {
	Name       string    `bson:"name"`
	Author     string    `bson:"author"`
	Newsest    string    `bson:"newsest"`
	Chapter    string    `bson:"chapter"`
	Source     string    `bson:"source"`
	Updatetime time.Time `bson:"updatetime,omitempty" json:"updatetime"`
	Createtime time.Time `bson:"createtime,omitempty" json:"createtime"`
}

type NovelBid struct {
	MgID string `bson:"_id"`

	Name       string    `bson:"name"`
	Author     string    `bson:"author"`
	Newsest    string    `bson:"newsest"`
	Chapter    string    `bson:"chapter"`
	Source     string    `bson:"source"`
	Updatetime time.Time `bson:"updatetime,omitempty" json:"updatetime"`
	Createtime time.Time `bson:"createtime,omitempty" json:"createtime"`
}

func NovelAdd(data Novel) (result *qmgo.InsertOneResult, err error) {
	mutex.Lock()
	defer mutex.Unlock()

	cli, err := qmgo.Open(ctx, &qmgo.Config{Uri: link, Database: conf.Mongodb.Db, Coll: "novel"})
	if err != nil {
		return nil, fmt.Errorf("mgdb open table err: %v", err)
	}

	one := Novel{}
	err = cli.Find(ctx, M{"name": data.Name}).One(&one)

	if err != nil {
		return NovelOriginAdd(data)
	}

	oneData := M{"$set": M{
		"name":       data.Name,
		"newsest":    data.Newsest,
		"chapter":    data.Chapter,
		"source":     data.Source,
		"updatetime": time.Now(),
	}}

	err = cli.UpdateOne(ctx, M{"name": data.Name}, oneData)
	if err != nil {
		return nil, fmt.Errorf("content update error: %v", err)
	}
	return nil, nil
}

func NovelOriginAdd(data Novel) (result *qmgo.InsertOneResult, err error) {

	data.Updatetime = time.Now()
	data.Createtime = time.Now()

	result, err = db.Collection("novel").InsertOne(ctx, data)
	if err != nil {
		return nil, fmt.Errorf("add error: %T", err)
	}
	return result, nil
}

func NovelOriginAddFindId(id, sort string, limit ...int64) (result []NovelSourceBid, err error) {
	var batch []NovelSourceBid

	cli, err := qmgo.Open(ctx, &qmgo.Config{Uri: link, Database: conf.Mongodb.Db, Coll: "novel"})
	if err != nil {
		return nil, fmt.Errorf("mgdb open table err: %v", err)
	}

	var bNum int64
	if len(limit) > 0 {
		bNum = limit[0]
	} else {
		bNum = 15
	}

	sortField := fmt.Sprintf("%s_id", sort)
	if strings.EqualFold(id, "") {
		err = cli.Find(ctx, D{}).Sort(sortField).Limit(bNum).All(&batch)
		return batch, err
	}

	mgId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return batch, err
	}

	opt := M{"_id": M{operator.Lt: mgId}}
	err = cli.Find(ctx, opt).Sort(sortField).Limit(bNum).All(&batch)
	return batch, err
}
