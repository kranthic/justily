package domain

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"time"
	"db"
)

type Session struct{
	Id bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Identified bool `bson:"i" json:"identified"`
	UserId bson.ObjectId `bson:"uid,omitempty" json:"userId"`
	LastAccessedTime time.Time `bson:"at" json:"lastAccessedTime"`
	Business string `bson:"b" json:"businessId"`
}

const collName = "session"

func getCollection(s *mgo.Session) *mgo.Collection{
	return s.DB(db.DbName()).C(collName)
}

func (session *Session) Save() error{
	s := db.NewMongoSession()
	defer s.Close()

	session.Id = bson.NewObjectId()
	coll := getCollection(s)
	return coll.Insert(session)
}

func (session *Session) UpdateUserId(userId string) error{
	s := db.NewMongoSession()
	defer s.Close()
	
	coll := getCollection(s)
	return coll.UpdateId(session.Id, session) 
}

func GetSessionById(sid string) (*Session, error){
	s := db.NewMongoSession()
	defer s.Close()
	
	id := bson.ObjectIdHex(sid)
	coll := getCollection(s)
	
	session := &Session{}
	err := coll.Find(bson.M{"_id": id}).One(session)
	
	return session, err
}