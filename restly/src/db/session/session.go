package session

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"time"
	"env"
)


type JustilySession struct{
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

func NewSession() (*JustilySession){
	s := env.NewMongoSession()
	defer s.Close()
	
	id := bson.NewObjectId()
	js := &JustilySession{Id: id, LastAccessedTime: time.Now(), Business: "0"}
	coll := getCollection(s)
	coll.Insert(js)
	
	return js
}

func IsUserRecoginzed(sid string) bool{
	if sid == ""{
		return false
	}
	
	return getSessionById(sid).UserId != ""
}

func GetSession(sid string) *JustilySession{
	if sid == ""{
		return &JustilySession{}
	}
	
	return getSessionById(sid)
}

func getSessionById(sid string) *JustilySession{
	s := env.NewMongoSession()
	defer s.Close()
	
	id := bson.ObjectIdHex(sid)
	coll := getCollection(s)
	js := &JustilySession{}
	coll.Find(bson.M{"_id": id}).One(js)
	return js
}

