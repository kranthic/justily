package session 

import (
	dbSess "db/session"
	"net/http"
	"fmt"
	"labix.org/v2/mgo/bson"
	"log"
	"time"
	"errors"
)

type RequestSession struct{
	Sid string
	session *dbSess.Session
}

const sessionIdCookie = "sid"
const userIdCookie = "uid"

func NewSession() (*RequestSession, error){
	session := &dbSess.Session{}
	session.Id = bson.NewObjectId()
	session.LastAccessedTime = time.Now()
	err := session.Save()
	
	if err != nil{
		log.Println(err)
		return nil, err
	}
	return &RequestSession{Sid: session.Id.Hex(), session: session}, nil
}

func HasSessionCookie(r *http.Request) bool{
	_, err := r.Cookie(sessionIdCookie)
	if err != nil{
		return false
	}
	
	return true
}

func GetRequestSession(r *http.Request) (*RequestSession, error){
	if !HasSessionCookie(r){
		return nil, errors.New("No session cookie")
	}
	
	cookie, _ := r.Cookie(sessionIdCookie)
	return &RequestSession{Sid: cookie.Value}, nil
}

func (rs *RequestSession) getSession(){
	session, err := dbSess.GetSessionById(rs.Sid)
	if err != nil{
		log.Print("Unable to retrieve session for ", rs.Sid, err)
		rs.session = &dbSess.Session{}
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
	fmt.Println(rs.Sid)
	return &http.Cookie{Name: sessionIdCookie, Value: rs.Sid}
}

func (rs *RequestSession) UserIdCookie() *http.Cookie{
	if rs.session == nil || rs.session.UserId == ""{
		return nil
	}
	return &http.Cookie{Name: userIdCookie, Value: rs.session.UserId.String()}
}