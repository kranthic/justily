package main 

import (
	"fmt"
	"net/http"
	"html/template"
	"helpers"
)

func cookieHandler(w http.ResponseWriter, r *http.Request){
	t, _ := template.ParseFiles("cookies.html")
	
	uuid, _ := helpers.GenUUID()
	cookie := &http.Cookie{Name: "sid", Value: uuid}
	http.SetCookie(w, cookie)
	t.Execute(w, nil)
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("bootstrap"))))
	http.HandleFunc("/", cookieHandler)
	fmt.Println(http.ListenAndServe(":8080", nil))	
}

