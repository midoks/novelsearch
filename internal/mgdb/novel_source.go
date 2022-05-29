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

	cliContent, err := qmgo.Open(ctx, &qmgo.Config{Uri: link, Database: conf.Mongodb.Db, Coll: "novel_source"})
	if err != nil {
		return nil, fmt.Errorf("mgdb open table err: %v", err)
	}

	one := NovelSource{}
	err = cliContent.Find(ctx, M{"name": data.Name}).One(&one)

	if err != nil {
		return NovelSourceOriginAdd(data)
	}

	oneData := M{"$set": M{
		"name":       data.Name,
		"rule_json":  data.RuleJson,
		"updatetime": time.Now(),
	}}

	err = cliContent.UpdateOne(ctx, M{"name": data.Name}, oneData)
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

	cliContent, err := qmgo.Open(ctx, &qmgo.Config{Uri: link, Database: conf.Mongodb.Db, Coll: "novel_source"})
	if err != nil {
		return nil, fmt.Errorf("mgdb open table err: %v", err)
	}

	var bNum int64
	if len(limit) > 0 {
		bNum = limit[0]
	} else {
		bNum = 15
	}
	err = cliContent.Find(ctx, D{}).Sort("-_id").Limit(bNum).All(&batch)
	return batch, err
}

func NovelSourceSearch(id, sort string, limit ...int64) (result []NovelSourceBid, err error) {
	var batch []NovelSourceBid

	cliContent, err := qmgo.Open(ctx, &qmgo.Config{Uri: link, Database: conf.Mongodb.Db, Coll: "novel_source"})
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
		err = cliContent.Find(ctx, D{}).Sort(sortField).Limit(bNum).All(&batch)
		return batch, err
	}

	mgId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return batch, err
	}

	opt := M{"_id": M{operator.Lt: mgId}}
	err = cliContent.Find(ctx, opt).Sort(sortField).Limit(bNum).All(&batch)
	return batch, err
}
