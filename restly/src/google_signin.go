package main 

import (
//	"net/http"
//	"net/url"
//	"encoding/json"
	"encoding/base64"
	"fmt"
//	"io/ioutil"
	"strings"
//	"log"
)

const GOOGLE_AUTH_URI = "https://accounts.google.com/o/oauth2/auth"
const GOOGLE_REVOKE_URI = "https://accounts.google.com/o/oauth2/revoke"
const GOOGLE_TOKEN_URI = "https://accounts.google.com/o/oauth2/token"
const GOOGLE_USER_DETAILS = "https://www.googleapis.com/oauth2/v3/userinfo?access_token=ya29.AHES6ZSceSQbxRbC7iI5gippczAyvuq9zZH9wAd_ptYolog"

func main(){
//	postForm := url.Values{
//					"grant_type" : {"authorization_code"},
//					"client_id" : {"551013514703.apps.googleusercontent.com"},
//					"client_secret" : {"m8fzousA0bNiqRrACSFYriNy"},
//					"code" : {"4/dIF_2hjwXjEOkTqn7-bOCrVCcifW.Yqa0SZrzmF8TMqTmHjyTFGN8-vWVhAI"},
//					"redirect_uri" : {"postmessage"}}
////					"scope": {"openid email"}}
//
//	response, err := http.PostForm(GOOGLE_TOKEN_URI, postForm)
//	if err != nil{
//		fmt.Println(err)
//	}
//	fmt.Println(response)
//	defer response.Body.Close()
//	
//	dataBytes, err := ioutil.ReadAll(response.Body)
//	fmt.Println(string(dataBytes))

	id_token := "eyJhbGciOiJSUzI1NiIsImtpZCI6IjViYzA4MWRjOGY2M2IyODk2ZjYyMWIyNWNlYWIwNjc1NmY5YjYzODkifQ.eyJpc3MiOiJhY2NvdW50cy5nb29nbGUuY29tIiwiYXRfaGFzaCI6InFyMGVfblpENFZqcEUtcDlIYXM2eUEiLCJhdWQiOiI1NTEwMTM1MTQ3MDMuYXBwcy5nb29nbGV1c2VyY29udGVudC5jb20iLCJlbWFpbCI6ImtyYW50aGkuY2hhbGFzYW5pQGdtYWlsLmNvbSIsImF6cCI6IjU1MTAxMzUxNDcwMy5hcHBzLmdvb2dsZXVzZXJjb250ZW50LmNvbSIsImVtYWlsX3ZlcmlmaWVkIjoidHJ1ZSIsInN1YiI6IjEwNDg3ODAwNDU1MjYyMzA5MzY5NSIsImlhdCI6MTM4NDIzNzAzMCwiZXhwIjoxMzg0MjQwOTMwfQ.mHmgU6sCySweFECkJTX_q8rcFASNKWjEqDmvwpEhP9fs5_TybkHv-JIDd9tHw4oitteTgrs9xM9u_-6Q2ajA__xgYcggmx3AqnnbFsYZw_fjUSZCl98DPoHAp133lczGrTGmPZNIek6gDDuVYlhjbG4ujv1PyMgdPVyJRGTpzdk"
	
	for idx,tok := range(strings.Split(id_token, ".")){
		if idx < 2{
			fmt.Println(strings.Trim(tok, "\n"))
			data, _:= base64.StdEncoding.DecodeString(strings.Trim(tok, "\n") + ".")
//			if err != nil{
//				log.Fatal(err)
//			}
			fmt.Println(string(data))
		}
	}
}