package main 

import (
	"net/http"
	"fmt"
	"log"
	"env"
	"db/session"
	"os"
	"runtime"
)


func home(w http.ResponseWriter, r *http.Request){
	s := session.NewSession()
	fmt.Println("blah", s)
	fmt.Fprintf(w, "Hello Kranthi!!!")
}

func root(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello Kranthi!!!")
}



func main() {
	runtime.GOMAXPROCS(8)
	dir, _ := os.Getwd()
	env.Init(dir + "/config.json", "dev")
	
//	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir(env.Config.StaticDir))))
	
	http.HandleFunc("/hello", home)
	http.HandleFunc("/", root)
	
//	loginHandler := &handlers.LoginHandler{}
//	loginHandler.RegisterHandlers()
	log.Println(http.ListenAndServe(fmt.Sprintf(":%s", env.Config.Port), nil))

}