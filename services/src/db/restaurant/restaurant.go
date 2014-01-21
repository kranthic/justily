package restaurant

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"db"
	dbu "db/user"
	"time"
//	"fmt"
)

type restaurant struct{
	Id bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Name string `bson:"n" json:"name"`
	About string `bson:"d,omitempty" json:"about"`
	Branches []branch `bson:"b" json:"branches"`
	AdminIds []bson.ObjectId `bson:"a" json:"-"`
	Admins []dbu.User	`bson:"-" json:"admins"`
	Key string `bson:"key" json:"-"` //Need a unique key that can be used to retrieve this restaurant. Name will not be unique
	CreateTime time.Time `bson:"ct" json:"created"`
	UpdateTime time.Time `bson:"ut" json:"updated"`
}

type branch struct{
	Name string `bson:"n" json:"name"`
	Address string `bson:"a" json:"address"`
	City string `bson:"c" json:"city"`
	State string `bson:"s" json:"state"`
	Zipcode string `bson:"z" json:"zipcode"`
}


const collection = "restaurant"

func NewRestaurant() *restaurant{
	return &restaurant{}
}

func (this *restaurant) NewBranch() *branch{
	return &branch{}
}

func (this *restaurant) AddBranch(b *branch){
	this.Branches = append(this.Branches, *b)
}

func GetRestaurantById(id string) (*restaurant, error){
	s := db.NewMongoSession()
	defer s.Close()
	
	c := s.DB(db.DbName()).C(collection)
	var r restaurant
	err := c.FindId(bson.ObjectIdHex(id)).One(&r)
	return &r, err
}

func (this *restaurant) Save() error{
	s := db.NewMongoSession()
	defer s.Close()
	
	var ins bool
	now := time.Now()
	if this.Id == ""{
		this.Id = bson.NewObjectId()
		this.Key = bson.NewObjectId().Hex()
		this.CreateTime = now
		ins = true
	}
	this.UpdateTime = now
	c := s.DB(db.DbName()).C(collection)
	c.EnsureIndex(mgo.Index{Key: []string{"key"},Unique: true})
	
	if ins{ return c.Insert(&this)} else{ return c.UpdateId(this.Id, &this)}
}
