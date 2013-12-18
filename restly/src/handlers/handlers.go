package handlers

import (
	"net/http"
	"net/url"
	"fmt"
	"html/template"
	"io/ioutil"
	"encoding/json"
	"github.com/restly/user"
	"time"
)

const GOOGLE_TOKEN_URI = "https://accounts.google.com/o/oauth2/token"
const GOOGLE_USER_DETAILS = "https://www.googleapis.com/oauth2/v3/userinfo?access_token="

type googleTokenDetails struct{
	Access_token string
	Token_type string
	Expires_in int
	id_token string
}

type googleUserDetails struct{
	Sub string
	Name string
	Given_name string
	Family_name string
	Profile string
	Picture string
	Email string
	Email_verified bool
	Gender string
	Locale string
}

type LoginData struct{
	IsIdentified bool
	Name string
}

type LoginHandler struct{
	
}

func (lh *LoginHandler) RegisterHandlers(){
	http.HandleFunc("/login", lh.login)
}

func (lh *LoginHandler) login(w http.ResponseWriter, r *http.Request){
	if r.FormValue("provider") == "" || r.FormValue("okey") == ""{
		lh.showLoginForm(&w, &LoginData{})
	} else {
		switch r.FormValue("provider"){
			case "google": lh.googleHandler(w, r)
		}
	}
}

func (lh *LoginHandler) googleHandler(w http.ResponseWriter, r *http.Request){
	postForm := url.Values{
					"grant_type" : {"authorization_code"},
					"client_id" : {"551013514703.apps.googleusercontent.com"},
					"client_secret" : {"m8fzousA0bNiqRrACSFYriNy"},
					"code" : {r.FormValue("okey")},
					"redirect_uri" : {"postmessage"},
					"scope": {"openid email"}}

	response, err := http.PostForm(GOOGLE_TOKEN_URI, postForm)
	if err != nil{
		fmt.Println(err)
	}
	defer response.Body.Close()
	
	dataBytes, err := ioutil.ReadAll(response.Body)
	googleTokens := &googleTokenDetails{}
	err = json.Unmarshal(dataBytes, googleTokens)
	userDetails := lh.getUserDetails(googleTokens)
	
	
	dataBytes, err = json.Marshal(userDetails)
	w.Header().Add("Content-Type", "text/html;charset=utf-8")
	cookie := &http.Cookie{Name: "at", Value: googleTokens.Access_token}
	http.SetCookie(w, cookie)
	ld := &LoginData{Name: userDetails.Name, IsIdentified: true}
	lh.showLoginForm(&w, ld)
	
}

func (lh *LoginHandler) getUserDetails(googleTokens *googleTokenDetails) *googleUserDetails{
	
	response, err := http.Get(GOOGLE_USER_DETAILS + googleTokens.Access_token)
	if err != nil{
		panic(err)
	}
	defer response.Body.Close()
	
	dataBytes, err := ioutil.ReadAll(response.Body)
	userDetails := &googleUserDetails{}
	err = json.Unmarshal(dataBytes, userDetails)
	go lh.storeUserDetails(userDetails)
	return userDetails
}

func (lh *LoginHandler) storeUserDetails(userDetails *googleUserDetails){
	u := &user.User{
		Email: userDetails.Email,
		FirstName: userDetails.Given_name,
		LastName: userDetails.Family_name,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	
	u.Save()
}

func (lh *LoginHandler) showLoginForm(w *http.ResponseWriter, ld *LoginData){
	t, _ := template.ParseFiles("login.html")
	t.Execute(*w, ld)
}