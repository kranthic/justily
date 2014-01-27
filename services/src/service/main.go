package main 

import (
	"net/http"
	"log"
	"service/handlers"
	"service/user"
	"service/restaurant"
	"service/restaurant/item"
	"service/restaurant/menu"
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
	http.Handle("/restaurant/item/add", handlers.JsonHandler(item.AddItem))
	http.Handle("/restaurant/item/sku/add", handlers.JsonHandler(item.AddSku))
	http.Handle("/restaurant/menu/add", handlers.JsonHandler(menu.AddMenu))
	http.Handle("/restaurant/menu/category/add", handlers.JsonHandler(menu.AddMenuCategory))
	http.Handle("/restaurant/menu/category/item/add", handlers.JsonHandler(menu.AddMenuCatItem))
	
	log.Print(http.ListenAndServe(":8080", nil))
}

