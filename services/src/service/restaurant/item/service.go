package item

import (
	dbi "db/restaurant/item"
	dbr "db/restaurant"
	"net/http"
	"encoding/json"
	"service/handlers"
	"log"
	"strconv"
)

func AddItem(r *handlers.HttpRequest) *handlers.JsonResponder{
	req := r.Request
	
	rkey := req.FormValue("rkey")
	if !isValidKey(rkey){
		return handlers.ErrorJsonResponder("Valid RestauranKey is required")
	}
	
	item := dbi.NewItem()
	item.Name = req.FormValue("name")
	item.Desc = req.FormValue("desc")
	
	if err := item.Save(rkey); err != nil{
		return handlers.ErrorJsonResponder(err.Error())
	}
	
	return itemJsonResponder(item)
}

func isValidKey(key string) bool{
	if key == ""{
		return false
	} 
	
	if _, err := dbr.GetRestaurantByKey(key); err != nil{
		log.Println("Tried to retrieve restaurant with key - ", key, err)
		return false
	}
	
	return true
}

func itemJsonResponder(item interface{}) *handlers.JsonResponder{
	bytes, err := json.Marshal(item)
	if err != nil{
		return handlers.ErrorJsonResponder(err.Error())
	}
	
	responder := &handlers.JsonResponder{}
	responder.SetStatus(http.StatusOK)
	responder.Write(string(bytes))
	return responder
}

func AddSku(r *handlers.HttpRequest) *handlers.JsonResponder{
	req := r.Request
	
	rkey := req.FormValue("rkey")
	if !isValidKey(rkey){
		return handlers.ErrorJsonResponder("Valid RestaurantKey is required")
	}
	
	itemId := req.FormValue("itemId")
	price := req.FormValue("price")
	pricef, err := strconv.ParseFloat(price, 32)
	if err != nil{
		return handlers.ErrorJsonResponder(err.Error())
	}
	
	item, err := dbi.GetByItemId(itemId, rkey)
	if err != nil{
		return handlers.ErrorJsonResponder(err.Error())
	}
	
	sku := item.NewSku()
	sku.Price = float32(pricef)
	item.AddSku(sku)
	
	if err := item.Save(rkey); err != nil{
		return handlers.ErrorJsonResponder(err.Error())
	}
	
	return itemJsonResponder(item)
	
}