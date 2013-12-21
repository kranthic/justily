package user

import (
	"net/http"
	"net/url"
	"io/ioutil"
	"encoding/json"
)

const GOOGLE_TOKEN_URI = "https://accounts.google.com/o/oauth2/token"
const GOOGLE_USER_DETAILS = "https://www.googleapis.com/oauth2/v3/userinfo?access_token="
const GOOGLE_CLIENT_SECRET = "m8fzousA0bNiqRrACSFYriNy"
const GOOGLE_CLIENT_ID = "551013514703.apps.googleusercontent.com"
const GOOGLE_GRANT_TYPE = "authorization_code"
const GOOGLE_SCOPE = "openid email"

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

type GoogleOAuth struct{
	userDetails *googleUserDetails
}

func (goa *GoogleOAuth) ProviderName() string{
	return "Google"
}

func (goa *GoogleOAuth) UserId() string{
	return goa.userDetails.Sub
}

func (goa *GoogleOAuth) UserName() string{
	return goa.userDetails.Name
}

func (goa *GoogleOAuth) FirstName() string{
	return goa.userDetails.Given_name
}

func (goa *GoogleOAuth) LastName() string{
	return goa.userDetails.Family_name
}

func (goa *GoogleOAuth) UserEmail() string{
	return goa.userDetails.Email
}

func (goa *GoogleOAuth) getGoogleTokenDetails(token string) (*googleTokenDetails, error){
	postForm := url.Values{
					"grant_type" : {GOOGLE_GRANT_TYPE},
					"client_id" : {GOOGLE_CLIENT_ID},
					"client_secret" : {GOOGLE_CLIENT_SECRET},
					"code" : {token},
					"redirect_uri" : {"postmessage"},
					"scope": {GOOGLE_SCOPE}}

	response, err := http.PostForm(GOOGLE_TOKEN_URI, postForm)
	if err != nil{
		return nil, err
	}
	defer response.Body.Close()
	
	dataBytes, err := ioutil.ReadAll(response.Body)
	googleToken := &googleTokenDetails{}
	err = json.Unmarshal(dataBytes, googleToken)
	if err != nil{
		return nil, err
	}
	
	return googleToken, nil
}

func (goa *GoogleOAuth) GetUserDetails(token string) error{
	tokenDetails, err := goa.getGoogleTokenDetails(token)
	if err != nil{
		return err
	}
	
	response, err := http.Get(GOOGLE_USER_DETAILS + tokenDetails.Access_token)
	if err != nil{
		return err
	}
	defer response.Body.Close()
	
	dataBytes, err := ioutil.ReadAll(response.Body)
	userDetails := &googleUserDetails{}
	err = json.Unmarshal(dataBytes, userDetails)
	if err != nil{
		return err
	}
	
	goa.userDetails = userDetails
	return nil
}