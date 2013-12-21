package user

import (
	"errors"
	"fmt"
	"domain"
	"time"
	"utils"
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


type UserService struct{
	Rs *utils.RequestSession
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
	user, err := domain.GetUserByOAuthId(oauth.ProviderName(), oauth.UserId())
	if err != nil{
		user = &domain.User{}
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