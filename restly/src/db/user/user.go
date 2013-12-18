package user

import (
	"labix.org/v2/mgo/bson"
	"time"
	"env"
)

type User struct{
	Id bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Email string `bson:"e" json:"e"`
	FirstName string `bson:"fn" json:"fn"`
	LastName string `bson:"ln" json:"ln"`
	Phone	string	`bson:"phn,omitempty" json:"phn"`
	CreateTime time.Time `bson:"ct" json:"created"`
	UpdateTime time.Time `bson:"ut" json:"updated"`
}

const collection = "user"

func ById(id string) (*User, error){
	session := env.NewMongoSession()
	defer session.Close()
	
	u := &User{}
	c := session.DB(env.Config.Mongo.Db).C(collection)
	err := c.Find(bson.M{"_id": id}).One(u)
	
	return u, err
}

func (u *User) Save() error{
	session := env.NewMongoSession()
	defer session.Close()
	
	c := session.DB(env.Config.Mongo.Db).C(collection)
	return c.Insert(&u)
}