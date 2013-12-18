package db

import (
	"labix.org/v2/mgo"
)

var session, connectErr = mgo.Dial("localhost")
const Name = "restly"

func Session()(*mgo.Session){
	if connectErr != nil{
		panic(connectErr)
	}
	
	return session.New()
}

