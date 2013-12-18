package user

import (
	"net/http"
	"net/url"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"db/user"
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

func GoogleHandler(w http.ResponseWriter, r *http.Request){
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
	userDetails := getUserDetails(googleTokens)
	
	
	dataBytes, err = json.Marshal(userDetails)
	w.Header().Add("Content-Type", "text/html;charset=utf-8")
	cookie := &http.Cookie{Name: "at", Value: googleTokens.Access_token}
	http.SetCookie(w, cookie)
	
}

func getUserDetails(googleTokens *googleTokenDetails) *googleUserDetails{
	
	response, err := http.Get(GOOGLE_USER_DETAILS + googleTokens.Access_token)
	if err != nil{
		panic(err)
	}
	defer response.Body.Close()
	
	dataBytes, err := ioutil.ReadAll(response.Body)
	userDetails := &googleUserDetails{}
	err = json.Unmarshal(dataBytes, userDetails)
	go storeUserDetails(userDetails)
	return userDetails
}

func storeUserDetails(userDetails *googleUserDetails){
	u := &user.User{
		Email: userDetails.Email,
		FirstName: userDetails.Given_name,
		LastName: userDetails.Family_name,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	
	u.Save()
}