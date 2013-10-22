package handlers

import (
	"net/http"
	"github.com/restly/db"
	"github.com/restly/service/helpers"
	"github.com/restly/service/consts"
	"encoding/json"
	"strings"
)




func Get(w http.ResponseWriter, r *http.Request){
	ckey,err := helpers.ConsumerKey(r)
	helpers.PanicOnError(err)
	
	var restaurantId string
	if !strings.Contains(r.URL.Path, "mine"){
		restaurantId = r.URL.Path[len(consts.RESTAURANT_GET_PATH):]
	}
	
	var restaurant *db.Restaurant
	if restaurantId != ""{
		restaurant,err = db.RestaurantById(restaurantId)
	} else{
		restaurant, err = db.RestaurantByKey(ckey)
	}
	helpers.PanicOnError(err)
	
	restaurantJson, err := json.Marshal(restaurant)
	helpers.PanicOnError(err)
	w.Write(restaurantJson)
}

func Add(w http.ResponseWriter, r *http.Request)(*db.Restaurant, error){
	return nil, nil
}

func Update(w http.ResponseWriter, r *http.Request)(*db.Restaurant, error){
	return nil, nil
}

