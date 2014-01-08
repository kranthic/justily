package handlers 

import (
	"net/http"
	"fmt"
	"service/session"
)

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

type JsonHandler func(req HttpRequest) JsonResponder

func (fn JsonHandler) ServeHTTP(w http.ResponseWriter, req *http.Request){
	responder := fn(HttpRequest{Request: req})
	setResponderData(w, &responder)
	
	w.Header().Add("Content-Type", "application/json;charset=UTF-8")
	fmt.Fprintf(w, responder.Output())
}

func setResponderData(w http.ResponseWriter, responder httpResponder){
	if responder.GetCookies() != nil{
		for _, ck := range *(responder.GetCookies()){
			http.SetCookie(w, &ck)
		}
	}
	
	if responder.GetStatus() != http.StatusOK{
		w.WriteHeader(responder.GetStatus())
	}
}

type HttpRequest struct{
	Session *session.RequestSession
	Request *http.Request
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