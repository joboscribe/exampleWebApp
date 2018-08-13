package main

import (
	"net/http"
	"github.com/gorilla/sessions"
	"log"
)

func sessionWrapper(f func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//check to see if the user is logged in, kick them back to login if they are not
		var s *sessions.Session

		if idString, ok := s.Values["userID"]; !ok {
			// this assumes we have a handler assigned for the "/login" route, which currently we do not
			http.Redirect(w, r, "/login", 302)
		} else {
			log.Print("logged in user ID: ", idString)
		}
		// if we have a userID in the session that means the user is logged in
		// so we let them proceed to their destination by calling the wrapped function
		f(w,r)

		// NOTE: we could do more stuff here
		// like add stuff to the session, look up the logged in user in the database, etc
	}
}

func clientHandler(w http.ResponseWriter, r *http.Request) {

}
