package user

import (
	"fmt"
	"net/http"
	"net/url"
//	"db/session"
	"db/user"
	"html/template"
	"service"
	"io/ioutil"
)

const GOOGLE_AUTH_URI = "https://accounts.google.com/o/oauth2/auth"
const GOOGLE_REVOKE_URI = "https://accounts.google.com/o/oauth2/revoke"
const GOOGLE_TOKEN_URI = "https://accounts.google.com/o/oauth2/token"
const GOOGLE_USER_DETAILS = "https://www.googleapis.com/oauth2/v3/userinfo?access_token=ya29.AHES6ZSceSQbxRbC7iI5gippczAyvuq9zZH9wAd_ptYolog"

type LoginStruct struct{
	IsIdentified bool
	Name string
}

func LoginCheckHandler(w http.ResponseWriter, r *http.Request){
	js := service.ManageSession(w, r)
	t, _ := template.ParseFiles("login.html")
	
	data := &LoginStruct{IsIdentified: js.UserId != ""}
	if js.UserId != ""{
		u, err := user.ById(js.UserId.String())
		if err == nil{
			data.Name = fmt.Sprintf("%s %s", u.FirstName, u.LastName)
		} else{
			data.IsIdentified = false
		}
	}
	t.Execute(w, data)
}

func LoginHandler(w http.ResponseWriter, r *http.Request){
	auth_token := r.FormValue("auth_token")
	postForm := url.Values{
				"grant_type" : {"authorization_code"},
				"client_id" : {"551013514703.apps.googleusercontent.com"},
				"client_secret" : {"m8fzousA0bNiqRrACSFYriNy"},
				"code" : {auth_token},
				"redirect_uri" : {"postmessage"}}
//				"scope": {"openid email"}}

	response, err := http.PostForm(GOOGLE_TOKEN_URI, postForm)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(response)
	defer response.Body.Close()
	
	dataBytes, err := ioutil.ReadAll(response.Body)
	
}


