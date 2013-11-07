package main 

import (
	"fmt"
)

type HttpHandler struct{
	Ckey string
	Uid string
}

type RestaurantHandler struct{
	HttpHandler
}

func main() {

	rh := RestaurantHandler{HttpHandler{Ckey: "a", Uid: "b"}}
	
	fmt.Println(rh.Ckey, rh.Uid)

}

