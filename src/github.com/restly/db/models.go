package db

import (
	"errors"
//	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"time"
)

var missing_req = errors.New("Required properties are missing")
var call_update = errors.New("Connot re-create object. Invoke Update() method to make updates")

type mongo struct{
	Session *mgo.Session
}

func (m *mongo) connect(){
	session, err := mgo.Dial("localhost")
	if err != nil{
		panic(err)
	}
	m.Session = session
}

func (m *mongo) session() *mgo.Session{
	if m.Session == nil{
		m.connect()
	}
	return m.Session.New()
}

var m = mongo{}
var db = "justily"

type User struct{
	Id bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Email string `bson:"e" json:"e"`
	Password string `bson:"p" json:"-"`
	FirstName string `bson:"fn" json:"fn"`
	LastName string `bson:"ln" json:"ln"`
	Phone	string	`bson:"phn,omitempty" json:"phn"`
	CreateTime time.Time `bson:"ct" json:"created"`
	UpdateTime time.Time `bson:"ut" json:"updated"`
}


type Restaurant struct{
	Id bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Name string `bson:"n" json:"name"`
	About string `bson:"d,omitempty" json:"about"`
	Branches []RestaurantBranch `bson:"b" json:"branches"`
	AdminIds []bson.ObjectId `bson:"a" json:"-"`
	Admins []User	`bson:"-" json:"admins"`
	Key string `bson:"key" json:"key"` //Need a unique key that can be used to retrieve this restaurant. Name will not be unique
	CreateTime time.Time `bson:"ct" json:"created"`
	UpdateTime time.Time `bson:"ut" json:"updated"`
}

type RestaurantBranch struct{
	Name string `bson:"n" json:"name"`
	Address string `bson:"a" json:"address"`
	City string `bson:"c" json:"city"`
	State string `bson:"s" json:"state"`
	Zipcode string `bson:"z" json:"zipcode"`
}

func (u *User) name() string{
	return "user"
}

func (u *User) valid() bool{
	return u.Email != "" && u.FirstName != "" && u.LastName != "" && u.Password != ""
}

func (u *User) Create() error{
	if !u.valid(){
		return missing_req
	}
	
	if bson.ObjectId.Valid(u.Id){
		return call_update
	}
	
	session := m.session()
	defer session.Close()
	
	u.CreateTime = time.Now()
	u.UpdateTime = time.Now()
	c := session.DB(db).C(u.name())
	c.EnsureIndex(mgo.Index{Key: []string{"e"},Unique: true})
	return c.Insert(&u)
}

func (u *User) Update() error{
	if !u.valid(){
		return missing_req
	}
	
	session := m.session()
	defer session.Close()
	
	u.UpdateTime = time.Now()
	c := session.DB(db).C(u.name())
	return c.Update(bson.M{"_id":u.Id}, u)
}

func (r *Restaurant) name() string{
	return "restaurant"
}

func (r *Restaurant) valid() bool{
	if r.Name == "" || r.AdminIds == nil || r.Branches == nil || len(r.Branches) == 0{
		return false
	} 
	
	for _,b := range(r.Branches){
		if b.Name == "" || b.Address == "" || b.City == "" || b.State == "" || b.Zipcode == ""{
			return false
		}	
	}
	return true 
}

func (r *Restaurant) Create() error{
	if !r.valid(){
		return missing_req
	}
	
	if bson.ObjectId.Valid(r.Id){
		return call_update
	}
	
	session := m.session()
	defer session.Close()
	
	r.CreateTime = time.Now()
	r.UpdateTime = time.Now()
	c := session.DB(db).C(r.name())
	c.EnsureIndex(mgo.Index{Key: []string{"key"}, Unique: true})
	return c.Insert(&r)
	
}

func (r *Restaurant) Update() error{
	if !r.valid(){
		return missing_req
	}
	
	session := m.session()
	defer session.Close()
	
	r.UpdateTime = time.Now()
	c := session.DB(db).C(r.name())
	return c.Update(bson.M{"_id":r.Id}, r)
}


func UserByEmail(email string) (user *User, err error){
	session := m.session()
	defer session.Close()
	c := session.DB(db).C("user")
	user = &User{}
	err = c.Find(bson.M{"e": email}).One(user)
	
	return user, err
}

func RestaurantByKey(key string) (*Restaurant, error){
	session := m.session()
	defer session.Close()
	
	r := &Restaurant{}
	c := session.DB(db).C(r.name())
	err := c.Find(bson.M{"key": key}).One(r)
	
	return r, err

}







