package handlers 

import (
	"net/http"
	"fmt"
	dbsn "db/session"
	"time"
	"encoding/json"
)

type HttpRequest struct{
	Session *dbsn.Session
	Request *http.Request
}


const sCkName = "sid"
func getSession(r *http.Request) *dbsn.Session{
	ck, err := r.Cookie(sCkName)
	if err != nil{
		return createNewSession()
	}
	
	if s, err := dbsn.GetSessionById(ck.Value); err != nil{
		return createNewSession()
	} else{
		return s
	}
	
}

func createNewSession() *dbsn.Session{
	s := &dbsn.Session{}
	s.LastAccessedTime = time.Now()
	if err := s.Save(); err != nil{
		panic(err)
	}
	
	return s
}

type httpResponder interface{
	AddCookie(cookie http.Cookie)
	SetStatus(status int)
	GetCookies() *[]http.Cookie
	GetStatus() int
}

type stringResponder interface{
	httpResponder
	Write(data string)
	Output() string
}

type JsonHandler func(req *HttpRequest) *JsonResponder

func (fn JsonHandler) ServeHTTP(w http.ResponseWriter, req *http.Request){
	session := getSession(req)

	responder := fn(&HttpRequest{Request: req, Session: session})
	if responder.GetStatus() == http.StatusOK{
		setResponderData(w, responder, session)
	} else {
		w.WriteHeader(responder.GetStatus())
	}
	
	w.Header().Add("Content-Type", "application/json;charset=UTF-8")
	fmt.Fprintf(w, responder.Output())
}

func setResponderData(w http.ResponseWriter, responder httpResponder, session *dbsn.Session){
	if responder.GetCookies() != nil{
		for _, ck := range *(responder.GetCookies()){
			http.SetCookie(w, &ck)
		}
		http.SetCookie(w, &http.Cookie{Name: "sid", Value: session.Id.Hex()})
	}
	
}

func ErrorJsonResponder(err string) *JsonResponder{
	responder := &JsonResponder{}
	responder.SetStatus(http.StatusBadRequest)
	
	type output struct{
		Err string
	}
	bytes, _ := json.Marshal(&output{Err: err})
	responder.Write(string(bytes))
	return responder
}

type JsonResponder struct{
	status int
	cookies []http.Cookie
	data string
}

func (jr *JsonResponder) AddCookie(cookie http.Cookie){
	jr.cookies = append(jr.cookies, cookie)
}

func (jr *JsonResponder) SetStatus(status int){
	jr.status = status
}

func (jr *JsonResponder) GetStatus() int{
	return jr.status
}

func (jr *JsonResponder) GetCookies() *[]http.Cookie{
	return &jr.cookies
}

func (jr *JsonResponder) Write(data string){
	jr.data = data
}

func (jr *JsonResponder) Output() string{
	return jr.data
}