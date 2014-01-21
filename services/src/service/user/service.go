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
	
	fmt.Println(oauth)
	
	return oauth, nil
}

func saveUserIfNew(oauth OAuth) (*dbu.User, error){

	fmt.Println(oauth)
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

func LoginUser(r *handlers.HttpRequest) *handlers.JsonResponder{
	authToken := r.Request.FormValue("at")
	provider := r.Request.FormValue("p")
	if authToken == "" || provider == ""{
		return handlers.ErrorJsonResponder("authToken and provider parameters are required")
	}
	
	errorResponder := func(err error) *handlers.JsonResponder{
		return handlers.ErrorJsonResponder(err.Error())
	}
	
	oauth, err := authorizeToken(authToken, provider)
	if err != nil{
		return errorResponder(err)
	}
	
	fmt.Println(oauth)
	user, err := saveUserIfNew(oauth)
	if err != nil{
		return errorResponder(err)
	}
	
	bytes, err := json.Marshal(user)
	if err != nil{
		return errorResponder(err)
	}
	
	cookie := &http.Cookie{Name: "uid", Value: user.Id.Hex()}
	responder := &handlers.JsonResponder{}
	responder.AddCookie(*cookie)
	responder.SetStatus(http.StatusOK)
	responder.Write(string(bytes))
	
	return responder
}