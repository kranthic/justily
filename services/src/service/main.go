package main 

import (
	"net/http"
	"log"
	"service/handlers"
	"service/user"
	"service/restaurant"
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
	http.Handle("/restaurant/add", handlers.JsonHandler(restaurant.AddNewRestaurant))
	http.Handle("/restaurant/branch/add", handlers.JsonHandler(restaurant.AddNewBranch))
	http.Handle("/restaurant/branch/edit", handlers.JsonHandler(restaurant.EditBranch))
	http.Handle("/restaurant", handlers.JsonHandler(restaurant.Restaurant))
	log.Print(http.ListenAndServe(":8080", nil))
}

