package config 

import (
	"os"
	"io/ioutil"
	"encoding/json"
	"db"
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
	db.InitMongo(Config.Mongo.Url, Config.Mongo.Db)
	inited = true
}
