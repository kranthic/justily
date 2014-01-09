package user

import (
	"errors"
	"fmt"
	dbu "db/user"
	"time"
	"service/handlers"
	"net/http"
	"encoding/json"
)

type OAuth interface{
	ProviderName() string
	UserId() string
	UserName() string
	FirstName() string
	LastName() string
	UserEmail() string
	GetUserDetails(token string) error
} 


func authorizeToken(token, provider string) (OAuth, error){
	var oauth OAuth
	var err error

	switch provider{
	case "google":
		oauth = &GoogleOAuth{}
		err = oauth.GetUserDetails(token)
	default :
		err = errors.New(fmt.Sprintf("Unrecognized provider %s", provider))
	}
	
	if err != nil{
		return nil, err
	}
	
	if oauth.UserId() == ""{
		return nil, errors.New(fmt.Sprintf("Unable to retrieve details based on token %s", token))
	}
	
	return oauth, nil
}

func saveUserIfNew(oauth OAuth) (*dbu.User, error){

	user, err := dbu.GetUserByOAuthId(oauth.ProviderName(), oauth.UserId())
	if err != nil{
		user = &dbu.User{}
		user.Email = oauth.UserEmail()
		user.FirstName = oauth.FirstName()
		user.LastName = oauth.LastName()
		user.OAuthProvider = oauth.ProviderName()
		user.OAuthUserId = oauth.UserId()
		
		now := time.Now()
		user.UpdateTime = now
		user.CreateTime = now
		
		err = user.Save()
	}
	
	return user, err 
}

func errorJsonResponse(err string) *handlers.JsonResponder{
	responder := &handlers.JsonResponder{}
	responder.SetStatus(http.StatusBadRequest)
	
	type output struct{
		Err string
	}
	bytes, _ := json.Marshal(&output{Err: err})
	responder.Write(string(bytes))
	return responder 
}

func LoginUser(r *handlers.HttpRequest) *handlers.JsonResponder{
	authToken := r.Request.FormValue("at")
	provider := r.Request.FormValue("p")
	if authToken == "" || provider == ""{
		return errorJsonResponse("authToken and provider parameters are required")
	}
	
	oauth, err := authorizeToken(authToken, provider)
	if err != nil{
		return errorJsonResponse(err.Error())
	}
	
	user, err := saveUserIfNew(oauth)
	if err != nil{
		return errorJsonResponse(err.Error())
	}
	
	bytes, err := json.Marshal(user)
	if err != nil{
		return errorJsonResponse(err.Error())
	}
	
	cookie := &http.Cookie{Name: "uid", Value: user.Id.Hex()}
	responder := &handlers.JsonResponder{}
	responder.AddCookie(*cookie)
	responder.SetStatus(http.StatusOK)
	responder.Write(string(bytes))
	
	return responder
}