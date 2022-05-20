package mgdb

import (
	"fmt"
	// "strings"
	"time"

	// "github.com/pkg/errors"
	"github.com/qiniu/qmgo"
	// "github.com/qiniu/qmgo/operator"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

type NovelSource struct {
	Name       string    `bson:"name"`
	RuleJson   string    `bson:"rule_json"`
	Updatetime time.Time `bson:"updatetime,omitempty" json:"updatetime"`
	Createtime time.Time `bson:"createtime,omitempty" json:"createtime"`
}

type NovelSourceBid struct {
	MgID       string    `bson:"_id"`
	Name       string    `bson:"name"`
	RuleJson   string    `bson:"rule_json"`
	Updatetime time.Time `bson:"updatetime,omitempty" json:"updatetime"`
	Createtime time.Time `bson:"createtime,omitempty" json:"createtime"`
}

func NovelSourceAdd(data NovelSource) (result *qmgo.InsertOneResult, err error) {
	mutex.Lock()
	defer mutex.Unlock()

	one := NovelSource{}
	err = cliContent.Find(ctx, M{"name": data.Name}).One(&one)

	fmt.Println("NovelSourceAdd:", err)
	if err != nil {
		return NovelSourceOriginAdd(data)
	}

	oneData := M{"$set": M{
		"name":       one.Name,
		"rule_json":  one.RuleJson,
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

	result, err = collection.InsertOne(ctx, data)
	if err != nil {
		return nil, fmt.Errorf("add error: %T", err)
	}
	return result, nil
}
