package env

import (
	"os"
	"io/ioutil"
	"encoding/json"
	"labix.org/v2/mgo"
//	"fmt"
)

type environments struct{
	Dev AppConfig
}

type AppConfig struct{
	Port string
	StaticDir string
	Mongo mongo
}

type mongo struct{
	Url string
	Db string
}

var Config *AppConfig
var session *mgo.Session
var inited = false

func Init(fileName string, environment string){
	f, err := os.Open(fileName)
	if err != nil{
		panic(err)
	}
	defer f.Close()
	
	dataInBytes, err := ioutil.ReadAll(f)
	if err != nil{
		panic(err)
	}
	
	env := &environments{}
	err = json.Unmarshal(dataInBytes, env)
	if err != nil{
		panic(err)
	}
	
	switch environment{
		default:
			Config = &env.Dev
	}
	
	initMongo()
	inited = true
}

func initMongo(){
	s, err := mgo.Dial(Config.Mongo.Url)
	if err != nil{
		panic(err)
	}
	
	session = s
}

func NewMongoSession() *mgo.Session{
	return session.New()
}