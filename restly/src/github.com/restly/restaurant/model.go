package restaurant

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"github.com/restly/db"
	"github.com/restly/user"
	
	"time"
)

type Restaurant struct{
	Id bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Name string `bson:"n" json:"name"`
	About string `bson:"d,omitempty" json:"about"`
	Branches []Branch `bson:"b" json:"branches"`
	AdminIds []bson.ObjectId `bson:"a" json:"-"`
	Admins []user.User	`bson:"-" json:"admins"`
	Key string `bson:"key" json:"key"` //Need a unique key that can be used to retrieve this restaurant. Name will not be unique
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

func ById(id string) (*Restaurant, error){
	session := db.Session()
	defer session.Close()
	
	r := &Restaurant{}
	c := session.DB(db.Name).C(collection)
	err := c.Find(bson.M{"_id": id}).One(r)
	
	return r, err
}

func ByKey(key string) (*Restaurant, error){
	session := db.Session()
	defer session.Close()
	
	r := &Restaurant{}
	c := session.DB(db.Name).C(collection)
	err := c.Find(bson.M{"key": key}).One(r)
	
	return r, err
}



func (r *Restaurant) Save() error{
	session := db.Session()
	defer session.Close()
	
	c := session.DB(db.Name).C(collection)
	c.EnsureIndex(mgo.Index{Key: []string{"key"}, Unique: true})
	return c.Insert(&r)
}
