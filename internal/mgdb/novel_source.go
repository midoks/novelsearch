package mgdb

import (
	"fmt"
	"strings"
	"time"

	// "github.com/pkg/errors"
	"github.com/qiniu/qmgo"
	"github.com/qiniu/qmgo/operator"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/midoks/novelsearch/internal/conf"
)

type NovelSource struct {
	Name       string    `bson:"name"`
	RuleJson   string    `bson:"rule_json"`
	Status     bool      `bson:"status"`
	Updatetime time.Time `bson:"updatetime,omitempty" json:"updatetime"`
	Createtime time.Time `bson:"createtime,omitempty" json:"createtime"`
}

type NovelSourceBid struct {
	MgID       string    `bson:"_id"`
	Name       string    `bson:"name"`
	RuleJson   string    `bson:"rule_json"`
	Status     bool      `bson:"status"`
	Updatetime time.Time `bson:"updatetime,omitempty" json:"updatetime"`
	Createtime time.Time `bson:"createtime,omitempty" json:"createtime"`
}

func NovelSourceAdd(data NovelSource) (result *qmgo.InsertOneResult, err error) {
	mutex.Lock()
	defer mutex.Unlock()

	cli, err := qmgo.Open(ctx, &qmgo.Config{Uri: link, Database: conf.Mongodb.Db, Coll: "novel_source"})
	if err != nil {
		return nil, fmt.Errorf("mgdb open table error: %v", err)
	}

	one := NovelSource{}
	err = cli.Find(ctx, M{"name": data.Name}).One(&one)

	if err != nil {
		return NovelSourceOriginAdd(data)
	}

	oneData := M{"$set": M{
		"name":       data.Name,
		"rule_json":  data.RuleJson,
		"updatetime": time.Now(),
	}}

	err = cli.UpdateOne(ctx, M{"name": data.Name}, oneData)
	if err != nil {
		return nil, fmt.Errorf("content update error: %v", err)
	}
	return nil, nil
}

func NovelSourceOriginAdd(data NovelSource) (result *qmgo.InsertOneResult, err error) {

	data.Updatetime = time.Now()
	data.Createtime = time.Now()

	result, err = db.Collection("novel_source").InsertOne(ctx, data)
	if err != nil {
		return nil, fmt.Errorf("add error: %T", err)
	}
	return result, nil
}

func NovelSourceOriginFind(limit ...int64) (result []NovelSourceBid, err error) {
	var batch []NovelSourceBid

	cli, err := qmgo.Open(ctx, &qmgo.Config{Uri: link, Database: conf.Mongodb.Db, Coll: "novel_source"})
	if err != nil {
		return nil, fmt.Errorf("mgdb open table error: %v", err)
	}

	var bNum int64
	if len(limit) > 0 {
		bNum = limit[0]
	} else {
		bNum = 15
	}
	err = cli.Find(ctx, D{}).Sort("-_id").Limit(bNum).All(&batch)
	return batch, err
}

func NovelSourceId(id string) (result NovelSourceBid, err error) {
	var ns NovelSourceBid

	cli, err := qmgo.Open(ctx, &qmgo.Config{Uri: link, Database: conf.Mongodb.Db, Coll: "novel_source"})
	if err != nil {
		return ns, fmt.Errorf("mgdb open table error: %v", err)
	}

	mgId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ns, err
	}

	opt := M{"_id": M{operator.Eq: mgId}}
	err = cli.Find(ctx, opt).One(&ns)
	return ns, err
}

func NovelSourceSearch(id, sort string, limit ...int64) (result []NovelSourceBid, err error) {
	var batch []NovelSourceBid

	cli, err := qmgo.Open(ctx, &qmgo.Config{Uri: link, Database: conf.Mongodb.Db, Coll: "novel_source"})
	if err != nil {
		return nil, fmt.Errorf("mgdb open table error: %v", err)
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
