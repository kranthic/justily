package restaurant

import (
	dbr "db/restaurant"
	"net/http"
	"encoding/json"
	"service/handlers"
	"strconv"
//	"fmt"
)

func EditBranch(r *handlers.HttpRequest) *handlers.JsonResponder{
	req := r.Request
	
	rid := req.FormValue("rid")
	bidx := req.FormValue("bidx")
	if rid == "" || bidx == ""{
		return handlers.ErrorJsonResponder("Restaurant Id (rid) and Branch Index (bidx) are required")
	}
	res, err := dbr.GetRestaurantById(rid)
	if err != nil{
		return handlers.ErrorJsonResponder(err.Error())
	}
	bidxInt, err := strconv.Atoi(bidx)
	if err != nil{
		return handlers.ErrorJsonResponder("Invalid Branch Idx - " + bidx)
	}
	if len(res.Branches) <= bidxInt{
		return handlers.ErrorJsonResponder("Invalid branch index")
	}
	b := res.Branches[bidxInt]
	if name := req.FormValue("bname"); name != ""{
		b.Name = name
	}
	if address := req.FormValue("address"); address != ""{
		b.Address = address
	}
	if city := req.FormValue("city"); city != ""{
		b.City = city
	}
	if state := req.FormValue("state"); state != ""{
		b.State = state
	}
	if zipcode := req.FormValue("zipcode"); zipcode != ""{
		b.Zipcode = zipcode
	}
	
	res.Branches[bidxInt] = b
	if err := res.Save(); err != nil{
		return handlers.ErrorJsonResponder(err.Error())
	}
	
	bytes, err := json.Marshal(res)
	if err != nil{
		return handlers.ErrorJsonResponder(err.Error())
	}
	responder := &handlers.JsonResponder{}
	responder.SetStatus(http.StatusOK)
	responder.Write(string(bytes))
	
	return responder
	
}


func AddNewBranch(r *handlers.HttpRequest) *handlers.JsonResponder{
	req := r.Request
	
	rid := req.FormValue("rid")
	if rid == ""{
		return handlers.ErrorJsonResponder("Restaurant Id (rid) required")
	}
	res, err := dbr.GetRestaurantById(rid)
	if err != nil{
		return handlers.ErrorJsonResponder(err.Error())
	}
	
	b := res.NewBranch()
	b.Name = req.FormValue("bname")
	b.Address = req.FormValue("address")
	b.City = req.FormValue("city")
	b.State = req.FormValue("state")
	b.Zipcode = req.FormValue("zipcode")
	
	res.AddBranch(b)
	err = res.Save()
	if err != nil{
		return handlers.ErrorJsonResponder(err.Error())
	}
	
	bytes, err := json.Marshal(res)
	if err != nil{
		return handlers.ErrorJsonResponder(err.Error())
	}
	responder := &handlers.JsonResponder{}
	responder.SetStatus(http.StatusOK)
	responder.Write(string(bytes))
	
	return responder
}


func AddNewRestaurant(r *handlers.HttpRequest) *handlers.JsonResponder{
	req := r.Request
	
	restaurant := dbr.NewRestaurant()
	restaurant.Name = req.FormValue("name")
	restaurant.About = req.FormValue("about")
	
	errorResponder := func(err error) *handlers.JsonResponder{
		return handlers.ErrorJsonResponder(err.Error())
	}
	
	if err := restaurant.Save(); err != nil{
		return errorResponder(err)
	}
	
	bytes, err := json.Marshal(restaurant)
	if err != nil{
		return errorResponder(err)
	}
	
	responder := &handlers.JsonResponder{}
	responder.SetStatus(http.StatusOK)
	responder.Write(string(bytes))
	
	return responder
}

func Restaurant(r *handlers.HttpRequest) *handlers.JsonResponder{
	req := r.Request
	
	rid := req.FormValue("rid")
	if rid == ""{
		return handlers.ErrorJsonResponder("Restaurant Id (rid) is required")
	}
	
	res, err := dbr.GetRestaurantById(rid)
	if err != nil{
		return handlers.ErrorJsonResponder(err.Error())
	}
	
	bytes, err := json.Marshal(res)
	if err != nil{
		return handlers.ErrorJsonResponder(err.Error())
	}
	
	responder := &handlers.JsonResponder{}
	responder.SetStatus(http.StatusOK)
	responder.Write(string(bytes))
	
	return responder
}