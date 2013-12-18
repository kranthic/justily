package service

import (
	"net/http"
	"db/session"
)


const sessionCookieName = "sid"

func ManageSession(w http.ResponseWriter, r *http.Request) *session.JustilySession{
	if isSessionCookieAvailable(r){
		c, _ := r.Cookie(sessionCookieName)
		if isValidSessionCookie(c){
			return session.GetSession(c.Value)
		}
	}
	js := session.NewSession()
	cookie := &http.Cookie{Name: "sid", Value: js.Id.String()}
	http.SetCookie(w, cookie)
	return js
}

func isSessionCookieAvailable(r *http.Request) bool{
	_, err := r.Cookie(sessionCookieName)
	return err == nil
}

func isValidSessionCookie(cookie *http.Cookie) bool{
	return cookie.Value != ""
}

func IsUserLoggedIn(js *session.JustilySession) bool{
	return js != nil && js.UserId != ""
}

