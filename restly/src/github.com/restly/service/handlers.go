package service 

import (
	"net/http"
	"github.com/restly/db"
	"encoding/json"
	"strings"
	"fmt"
	"errors"
)

type Response struct{
	Status int `json:status`
	Data interface{} `json:data`
	Error string `json:error`
	Diag string	`json:diag`
}

type RestaurantHandler struct{
	basePath string
}

func (r *RestaurantHandler) RegisterHandlerFuncs(){
	r.basePath = "/restaurants"
	http.HandleFunc(r.basePath + "/", r.handle)
	http.HandleFunc(r.basePath, r.handle)
}

func (r *RestaurantHandler) handle(w http.ResponseWriter, req *http.Request){
	w.Header().Add("Content-Type", "application/json;charset=utf-8")
	
	var call = func(fn func(req *http.Request) *Response) error{
					response := fn(req)
					responseInBytes, err := json.Marshal(response)
					fmt.Println(string(responseInBytes))
					if err == nil{
						w.WriteHeader(http.StatusOK)
						w.Write(responseInBytes)
					}
					
					return err
				}
	
	var err error
	if req.Method == "GET"{
		if err = call(r.get); err == nil{
			return
		}
		
	}else if req.Method == "POST"{
		if err = call(r.post); err == nil{
			return
		}
	}else if req.Method == "PUT"{
		if err = call(r.put); err == nil{
			return
		}
	} else{
		err = errors.New("Unsupported HTTP method")
	}
	
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		response,_ := json.Marshal(&Response{Status: http.StatusBadRequest, Error: err.Error()})
		w.Write(response)
	}
}

func missingConsumerKeyResponse() *Response{
	return &Response{Status: http.StatusBadRequest, Error: "Consumer Key is missing"}
}

func errorResponse(err error) *Response{
	return &Response{Status: http.StatusInternalServerError, Error: err.Error()}
}

func (r *RestaurantHandler) get(req *http.Request) *Response{
	ckey, err := ConsumerKey(req)
	if err != nil{
		return missingConsumerKeyResponse()
	}
	
	var restaurantId string
	if !strings.Contains(req.URL.Path, "mine"){
		restaurantId = req.URL.Path[len(r.basePath):]
	}
	
	var restaurant *db.Restaurant
	if restaurantId != ""{
		restaurant, err = db.RestaurantById(restaurantId)
	} else{
		restaurant, err = db.RestaurantByKey(ckey)
	}
	
	if err != nil{
		return errorResponse(err)
	}
	
	return &Response{Status: http.StatusOK, Data: *restaurant}
	
}

func (r *RestaurantHandler) post(req *http.Request) *Response{
	return nil
}

func (r *RestaurantHandler) put(req *http.Request) *Response{
	return nil
}