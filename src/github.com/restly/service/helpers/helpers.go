package helpers

import (
	"net/http"
	"github.com/restly/service/consts"
	"errors"
//	"fmt"
	"encoding/json"
)

type Response struct{
	Status int `json:status`
	Data interface{} `json:data`
	Error string `json:error`
	Diag string	`json:diag`
}
func ConsumerKey(r *http.Request)(string, error){
	ckey := r.FormValue(consts.CKEY_URL)
	if ckey == ""{
		ckey = r.Header.Get(consts.CKEY_HEADER)
	}
	
	if ckey == ""{
		return "", errors.New("Consumer Key not available")
	}
	return ckey, nil

}

func GetOnly(f http.HandlerFunc)(http.HandlerFunc){
	return func(w http.ResponseWriter, r *http.Request){
		if r.Method == "GET"{
			f(w, r)
		} else{
			panic("Supports only GET requests")
		}
	}
}

func GetnPut(f http.HandlerFunc)(http.HandlerFunc){
	return func(w http.ResponseWriter, r *http.Request){
		if r.Method == "GET" || r.Method == "PUT"{
			f(w, r)
		} else{
			panic("Supports only GET or PUT requests")
		}
	}
}

func ErrorHandler(fn http.HandlerFunc) http.HandlerFunc {
        return func(w http.ResponseWriter, r *http.Request) {
                defer func() {
                        if err, ok := recover().(error); ok {
                        	response := Response{Status:http.StatusInternalServerError, Error: err.Error()}
                        	responseJson,_ := json.Marshal(response)
                        	http.Error(w, string(responseJson), http.StatusForbidden)
                        }
                }()
                fn(w, r)
        }
}


func PanicOnError(err error){
	if err != nil{
		panic(err)
	}
}
