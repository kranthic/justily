package domain

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"time"
	"db"
)

type User struct{
	Id bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Email string `bson:"e" json:"email"`
	FirstName string `bson:"fn" json:"firstName"`
	LastName string `bson:"ln" json:"lastName"`
	OAuthProvider string `bson:"oap"`
	OAuthUserId string `bson:"uid"`
	Phone	string	`bson:"phn,omitempty" json:"phone"`
	CreateTime time.Time `bson:"ct" json:"createTime"`
	UpdateTime time.Time `bson:"ut" json:"updateTime"`
}

const collection = "user"

func (user *User) Save() error{
	s := db.NewMongoSession()
	defer s.Close()
	
	var ins bool
	if user.Id == ""{
		user.Id = bson.NewObjectId()
		ins = true
	}
	c := s.DB(db.DbName()).C(collection)
	c.EnsureIndex(mgo.Index{Key: []string{"oap", "uid"},Unique: true})
	if ins{ return c.Insert(&user) }else {return c.UpdateId(user.Id, &user)}
}

func GetUserById(id string) (*User, error){
	s := db.NewMongoSession()
	defer s.Close()
	
	u := &User{}
	c := s.DB(db.DbName()).C(collection)
	err := c.Find(bson.M{"_id": id}).One(u)
	
	return u, err
}

func GetUserByOAuthId(provider, id string) (*User, error){
	s := db.NewMongoSession()
	defer s.Close()
	
	u := &User{}
	c := s.DB(db.DbName()).C(collection)
	err := c.Find(bson.M{"oap": provider, "uid": id}).One(u)
	
	return u, err
}