package main

import (
	"net/http"
	"log"
)

func index(w http.ResponseWriter, r *http.Request){
}

func main() {
	log.Print(http.ListenAndServe(":8080", nil))
}

