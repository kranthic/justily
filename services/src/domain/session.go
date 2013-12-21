package domain

import (
	"env"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"time"
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
	return s.DB(env.Config.Mongo.Db).C(collName)
}

func (session *Session) Save() error{
	s := env.NewMongoSession()
	defer s.Close()

	session.Id = bson.NewObjectId()
	coll := getCollection(s)
	return coll.Insert(session)
}

func GetSessionById(sid string) (*Session, error){
	s := env.NewMongoSession()
	defer s.Close()
	
	id := bson.ObjectIdHex(sid)
	coll := getCollection(s)
	
	session := &Session{}
	err := coll.Find(bson.M{"_id": id}).One(session)
	
	return session, err
}