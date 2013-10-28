package service 

import (
	"net/http"
	"errors"
	"encoding/json"
)

type JustilyHandlerFunc func(r *http.Request)(interface{}, error)

func ConsumerKey(r *http.Request)(string, error){
	ckey := r.FormValue(CKEY_URL)
	if ckey == ""{
		ckey = r.Header.Get(CKEY_HEADER)
	}
	
	if ckey == ""{
		return "", errors.New("Consumer Key not available")
	}
	return ckey, nil

}

func GetOnly(f JustilyHandlerFunc)(http.HandlerFunc){
	return func(w http.ResponseWriter, r *http.Request){
		if r.Method == "GET"{
			 data, err := f(r)
			 var response Response
			 if err == nil{
             	response = Response{Status:http.StatusOK, Data: data}
             } else{
             	response = Response{Status:http.StatusInternalServerError, Error: err.Error()}
             }
             responseJson,err := json.Marshal(response)
             w.Write(responseJson)
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
                        	PanicOnError(err)
                        	response := Response{Status:http.StatusInternalServerError, Error: err.Error()}
                        	responseJson,_ := json.Marshal(response)
                        	http.Error(w, string(responseJson), http.StatusForbidden)
                        }
                }()
                fn(w, r)
        }
}

//func JustilyHandler(fn JustilyHandlerFunc) http.HandlerFunc {
//        return func(w http.ResponseWriter, r *http.Request) {
//                defer func() {
//                        if err, ok := recover().(error); ok {
//                        	response := Response{Status:http.StatusInternalServerError, Error: err.Error()}
//                        	responseJson,_ := json.Marshal(response)
//                        	http.Error(w, string(responseJson), http.StatusForbidden)
//                        }
//                }()
//                data, err := fn(r)
//                response := Response{Status:http.StatusOK, Data: data, Error: err.Error()}
//                responseJson,_ := json.Marshal(response)
//                w.Write(responseJson)
//        }
//}


func PanicOnError(err error){
	if err != nil{
		panic(err)
	}
}
