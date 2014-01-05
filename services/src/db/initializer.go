package db

import (
	"labix.org/v2/mgo"
)

var session *mgo.Session
var db string

func InitMongo(mongoUrl, dbName string){
	s, err := mgo.Dial(mongoUrl)
	if err != nil{
		panic(err)
	}
	if dbName == ""{
		panic("Database name not provided")
	}
	
	db = dbName
	session = s
}

func NewMongoSession() *mgo.Session{
	if session == nil{
		panic("Mongo connection not initialized")
	}
	return session.New()
}

func DbName() string{
	return db
}