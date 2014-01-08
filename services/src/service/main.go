package main 

import (
	"net/http"
	"log"
	"service/handlers"
	"service/user"
)

func main() {
//	currentDir, err := os.Getwd()
//	if err != nil{
//		log.Fatal(err)
//	}
//	configJson := fmt.Sprintf("%s/config.json", currentDir)
//	config.Init(configJson, "dev")
//	
	http.Handle("/user", handlers.JsonHandler(user.HelloUser))
	log.Print(http.ListenAndServe(":8080", nil))
}

