package utils 

import (
	"domain"
	"net/http"
//	"fmt"
	"labix.org/v2/mgo/bson"
	"log"
)

type RequestSession struct{
	Sid string
	session *domain.Session
}

func (rs *RequestSession) getSession(){
	session, err := domain.GetSessionById(rs.Sid)
	if err != nil{
		log.Print("Unable to retrieve session for ", rs.Sid, err)
		rs.session = &domain.Session{}
	}
	
	rs.session = session
}

func (rs *RequestSession) UpdateUserId(userId string) error{
	rs.session.UserId = bson.ObjectIdHex(userId)
	return rs.session.Save()
}

func (rs *RequestSession) IsValid() bool{
	if rs.session == nil{
		rs.getSession()
	}
	
	return rs.session.Id != ""
}

func (rs *RequestSession) IsUserLoggedIn() bool{
	if rs.session == nil{
		rs.getSession()
	}
	
	return rs.session.UserId != ""
}

func (rs *RequestSession) Cookie() *http.Cookie{
	return &http.Cookie{Name: "sid", Value: rs.Sid}
}

func (rs *RequestSession) UserIdCookie() *http.Cookie{
	if rs.session == nil || rs.session.UserId == ""{
		return nil
	}
	return &http.Cookie{Name: "uid", Value: rs.session.UserId.String()}
}