package mgdb

import (
	"fmt"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/qiniu/qmgo"
	"github.com/qiniu/qmgo/operator"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type File struct {
	Path   []interface{} `bson:"path";json:"path"`
	Length int           `bson:"length"`
}

type BitTorrent struct {
	InfoHash   string    `bson:"infohash"`
	Name       string    `bson:"name"`
	Files      []File    `bson:"files,omitempty"`
	Length     int       `bson:"length,omitempty"`
	Updatetime time.Time `bson:"updatetime" json:"updatetime"`
	Createtime time.Time `bson:"createtime" json:"createtime"`
}

type BitTorrentBid struct {
	MgID       string    `bson:"_id"`
	InfoHash   string    `bson:"infohash"`
	Name       string    `bson:"name"`
	Files      []File    `bson:"files,omitempty"`
	Length     int       `bson:"length,omitempty"`
	Updatetime time.Time `bson:"updatetime" json:"updatetime"`
	Createtime time.Time `bson:"createtime" json:"createtime"`
}

func TorrentAdd(data BitTorrent) (result *qmgo.InsertOneResult, err error) {

	one := BitTorrentBid{}
	err = cliContent.Find(ctx, M{"infohash": data.InfoHash}).One(&one)

	if err != nil {
		return TorrentOriginAdd(data)
	}

	oneData := M{"$set": M{
		"files":      data.Files,
		"updatetime": time.Now(),
	}}

	err = cliContent.UpdateOne(ctx, M{"infohash": data.InfoHash}, oneData)
	if err != nil {
		return nil, errors.Wrap(err, "bt update")
	}
	return nil, nil
}

func TorrentOriginAdd(data BitTorrent) (result *qmgo.InsertOneResult, err error) {

	dlen := 0
	for _, f := range data.Files {
		dlen += f.Length
	}

	data.Length = dlen
	data.Updatetime = time.Now()
	data.Createtime = time.Now()

	result, err = collection.InsertOne(ctx, data)
	if err != nil {
		return nil, errors.Wrap(err, "bt add")
	}
	return result, nil
}

func TorrentOriginFindSoso(id, sort, op, keyword string, limit ...int64) (result []BitTorrentBid, err error) {
	var batch []BitTorrentBid

	var bNum int64
	if len(limit) > 0 {
		bNum = limit[0]
	} else {
		bNum = 15
	}

	sortField := fmt.Sprintf("%s_id", sort)

	opt := M{}
	if !strings.EqualFold(id, "") {
		mgId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return batch, err
		}
		opt["_id"] = M{op: mgId}
	}

	if !strings.EqualFold(keyword, "") {
		opt["name"] = M{operator.Regex: keyword, "$options": "im"}
	}

	err = cliContent.Find(ctx, opt).Sort(sortField).Limit(bNum).All(&batch)
	return batch, err
}
