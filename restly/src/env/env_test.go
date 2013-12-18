package env

import (
	"testing"
	"fmt"
)

func TestAppConfigMustBeInitialized(t *testing.T){
	defer func(){
		if r := recover(); r == nil{
			fmt.Println("We must panic that application is not configured but, nothing happened.")
			t.Fail()
		}
	}()
	fmt.Println(NewMongoSession())
}

func TestInitializingAppConfig(t *testing.T){
	Init("../../config.json", "dev")
	defer func(){
		if r := recover(); r != nil{
			fmt.Println("We must not panic. We initialized the application already.")
			t.Fail()
		}
	}()
	fmt.Println(NewMongoSession())
}

