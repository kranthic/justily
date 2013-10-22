package main 

import (
	"fmt"
	"net/http"
	"github.com/restly/db"	
	"github.com/restly/service/consts"
	"github.com/restly/service/handlers"
	"github.com/restly/service/helpers"
	"encoding/json"
	"log"
)

func user(w http.ResponseWriter, r *http.Request){
	email := r.URL.Path[len("/user/"):]
	u, err := db.UserByEmail(email)
	if(err != nil){
		w.WriteHeader(http.StatusInternalServerError)
	} else{
		buf, _ := json.Marshal(u)
		w.Write(buf)
	}
}

func hello(w http.ResponseWriter, r *http.Request){
	u,err := db.UserByEmail("kranthi.chalasani@gmail.com")
	if(err != nil){
		log.Fatal(err)
	}
	buf,_ := json.Marshal(u)
	w.Write(buf)
}

func hello2(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello %s", r.URL.Path[1:])
}

func main() {

	http.HandleFunc(consts.MY_RESTAURANT_PATH, helpers.ErrorHandler(helpers.GetOnly(handlers.Get)))
	http.HandleFunc("/user/", user)
	http.HandleFunc("/hello/", hello2)
	http.HandleFunc("/", hello)
	log.Println(http.ListenAndServe(":8080", nil))

}

