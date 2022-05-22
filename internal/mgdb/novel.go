package mgdb

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
	MgID       string    `bson:"_id"`
	Name       string    `bson:"name"`
	Author     string    `bson:"author"`
	Newsest    string    `bson:"newsest"`
	Chapter    string    `bson:"chapter"`
	Source     string    `bson:"source"`
	Updatetime time.Time `bson:"updatetime,omitempty" json:"updatetime"`
	Createtime time.Time `bson:"createtime,omitempty" json:"createtime"`
}
