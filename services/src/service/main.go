package main 

import (
	"net/http"
	"log"
	"service/handlers"
	"service/user"
	"os"
	"fmt"
	"config"
)

func main() {
	currentDir, err := os.Getwd()
	if err != nil{
		log.Fatal(err)
	}
	configJson := fmt.Sprintf("%s/config.json", currentDir)
	config.Init(configJson, "dev")
	
	http.Handle("/user/login", handlers.JsonHandler(user.LoginUser))
	log.Print(http.ListenAndServe(":8080", nil))
}

