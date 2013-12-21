package user

import (
	"errors"
	"fmt"
//	"net/http"
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


func authorizeToken(token, provider string) (*OAuth, error){
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
