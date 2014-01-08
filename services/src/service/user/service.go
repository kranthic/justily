package user

import (
	"errors"
	"fmt"
	dbu "db/user"
	"time"
	"service/session"
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


type UserService struct{
	Rs *session.RequestSession
}

func (us *UserService) authorizeToken(token, provider string) (*OAuth, error){
	var oauth OAuth
	var err error

	switch provider{
	case "google":
		oauth := &GoogleOAuth{}
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
	
	return &oauth, nil
}

func (us *UserService) saveUserIfNew(oauth OAuth) (string, error){
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
	
	if err != nil{
		return "", err
	}
	
	return user.Id.String(), nil
}

func (us *UserService) updateRequestSession(userId string) error{
	return us.Rs.UpdateUserId(userId)
}

func HelloUser(r handlers.HttpRequest) handlers.JsonResponder{
	fmt.Println("Running User Handler")
	cookie := &http.Cookie{Name:"blah", Value: "blah"}
	var responder handlers.JsonResponder
	responder.AddCookie(*cookie)
	u := &dbu.User{FirstName: "Kranthi"}
	bytes, _ := json.Marshal(u)
	responder.Write(string(bytes))
	responder.SetStatus(http.StatusOK)
	
	return responder	
}