package restaurant

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"db"
	dbu "db/user"
	"time"
//	"fmt"
)

type Restaurant struct{
	Id bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Name string `bson:"n" json:"name"`
	About string `bson:"d,omitempty" json:"about"`
	Branches []Branch `bson:"b" json:"branches"`
	AdminIds []bson.ObjectId `bson:"a" json:"-"`
	Admins []dbu.User	`bson:"-" json:"admins"`
	Key string `bson:"key" json:"-"` //Need a unique key that can be used to retrieve this restaurant. Name will not be unique
	CreateTime time.Time `bson:"ct" json:"created"`
	UpdateTime time.Time `bson:"ut" json:"updated"`
}

type Branch struct{
	Name string `bson:"n" json:"name"`
	Address string `bson:"a" json:"address"`
	City string `bson:"c" json:"city"`
	State string `bson:"s" json:"state"`
	Zipcode string `bson:"z" json:"zipcode"`
}


const collection = "restaurant"

func NewRestaurant() *Restaurant{
	return &Restaurant{}
}

func (this *Restaurant) NewBranch() *Branch{
	return &Branch{}
}

func (this *Restaurant) AddBranch(b *Branch){
	this.Branches = append(this.Branches, *b)
}

func GetRestaurantByKey(key string) (*Restaurant, error){
	s := db.NewMongoSession()
	defer s.Close()
	
	c := s.DB(db.DbName()).C(collection)
	var r Restaurant
	err := c.Find(bson.M{"key": key}).One(&r)
	return &r, err
}

func GetRestaurantById(id string) (*Restaurant, error){
	s := db.NewMongoSession()
	defer s.Close()
	
	c := s.DB(db.DbName()).C(collection)
	var r Restaurant
	err := c.FindId(bson.ObjectIdHex(id)).One(&r)
	return &r, err
}

func (this *Restaurant) Save() error{
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
