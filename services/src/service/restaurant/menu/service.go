package menu

import (
	dbr "db/restaurant"
	dbm "db/restaurant/menu"
	dbi "db/restaurant/item"
	"net/http"
	"encoding/json"
	"service/handlers"
	"log"
	"strconv"
)

func AddMenu(r *handlers.HttpRequest) *handlers.JsonResponder{
	req := r.Request
	
	rkey := req.FormValue("rkey")
	if !isValidKey(rkey){
		return handlers.ErrorJsonResponder("Valid RestaurantKey is required")
	}
	
	name := req.FormValue("name")
	if name == ""{
		return handlers.ErrorJsonResponder("Name is required")
	}
	
	menu := dbm.NewMenu()
	menu.Name = name
	menu.Save(rkey)
	
	return menuJsonResponder(menu)
}

func AddMenuCategory(r *handlers.HttpRequest) *handlers.JsonResponder{
	req := r.Request
	rkey := req.FormValue("rkey")
	if !isValidKey(rkey){
		return handlers.ErrorJsonResponder("Valid RestaurantKey is required")
	}
	
	name := req.FormValue("name")
	if name == ""{
		return handlers.ErrorJsonResponder("Name is required")
	}
	
	menuId := req.FormValue("menuId")
	if menuId == ""{
		return handlers.ErrorJsonResponder("MenuId is required")
	}
	
	menu, err := dbm.GetMenuById(menuId, rkey)
	if err != nil{
		return handlers.ErrorJsonResponder(err.Error())
	}
	
	category := menu.NewCategory()
	category.Name = name
	menu.AddCategory(category)
	
	menu.Save(rkey)
	
	return menuJsonResponder(menu)
}

func AddMenuCatItem(r *handlers.HttpRequest) *handlers.JsonResponder{
	req := r.Request
	req.ParseForm()
	rkey := req.FormValue("rkey")
	if !isValidKey(rkey){
		return handlers.ErrorJsonResponder("Valid RestaurantKey is required")
	}
	
	
	itemId := req.FormValue("itemId")
	if itemId == ""{
		return handlers.ErrorJsonResponder("ItemId is required")
	}
	
	item, err := dbi.GetByItemId(itemId, rkey)
	if err != nil{
		return handlers.ErrorJsonResponder(err.Error())
	}
	
	skus := req.Form["skus"]
	if skus == nil || len(skus) == 0{
		return handlers.ErrorJsonResponder("Skus are required")
	}
	
	menuId := req.FormValue("menuId")
	if menuId == ""{
		return handlers.ErrorJsonResponder("MenuId is required")
	}
	menu, err := dbm.GetMenuById(menuId, rkey)
	if err != nil{
		return handlers.ErrorJsonResponder(err.Error())
	}
	
	catId := req.FormValue("catId")
	if catId == ""{
		return handlers.ErrorJsonResponder("CategoryId is required")
	}
	catIdx, err := strconv.Atoi(catId)
	if err != nil{
		return handlers.ErrorJsonResponder(err.Error())
	}
	if catIdx >= len(menu.Categories){
		return handlers.ErrorJsonResponder("Invalid CategoryId")
	}
	category := menu.Categories[catIdx]
	
	catItem := category.NewCatItem()
	catItem.ItemId = item.Id
	
	ci := *item
	ci.Skus = make([]dbi.Sku, len(skus))
	for idx, sku := range skus{
		if skuInt, err := strconv.Atoi(sku); err != nil{
			return handlers.ErrorJsonResponder(err.Error())
		} else{
			if skuInt >= len(item.Skus){
				return handlers.ErrorJsonResponder("Invalid sku")
			}
			
			catItem.Sku = append(catItem.Sku, skuInt)		
			ci.Skus[idx] = item.Skus[skuInt]  
		}
	}
	catItem.Item = &ci
	
	category.AddItem(catItem)
	menu.Categories[catIdx] = category
	menu.Save(rkey)
	
	return menuJsonResponder(menu)
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

func menuJsonResponder(menu *dbm.Menu) *handlers.JsonResponder{
	bytes, err := json.Marshal(menu)
	if err != nil{
		return handlers.ErrorJsonResponder(err.Error())
	}
	
	responder := &handlers.JsonResponder{}
	responder.SetStatus(http.StatusOK)
	responder.Write(string(bytes))
	return responder
}
