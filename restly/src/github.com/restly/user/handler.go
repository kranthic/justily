package user

import (
	"net/http"
)

const basePath = "/user"
func AddHandlers(){
	http.HandleFunc(basePath + "/", nil)
}

func login(w http.ResponseWriter, r *http.Request){



}
