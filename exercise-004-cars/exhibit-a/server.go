package main

import (
	"fmt"
	"net/http"
	"time"
)

func inOneYear() time.Time {
	return time.Now().AddDate(1, 0, 0) // func to return time of one year from now
}

func poke(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{Name: "username", Value: "gopher", Expires: inOneYear()} // init cookie (https://golang.org/pkg/net/http/#Cookie)
	http.SetCookie(w, &cookie)                                                     // adds a Set-Cookie header to the provided ResponseWriter's headers
	fmt.Fprintf(w, "Just set cookie named 'username' set to 'gopher'")
}

func peek(w http.ResponseWriter, r *http.Request) {
	username, err := r.Cookie("username") // grab cookie named "username"
	if err != nil {
		fmt.Fprintf(w, "Could not find cookie named 'username'") // if there is no cookie named "username", print error and return
		return
	}
	fmt.Fprintf(w, "You have a cookie named 'username' set to '%s'", username) // print username to browser
}

func hide(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{Name: "username", MaxAge: -1} // init cookie, "MaxAge<0 means delete cookie now", same Name as previous cookie
	http.SetCookie(w, &cookie)                          // writes a cookie w/ valid name, so "username" cookie will now overwrite and delete
	fmt.Fprintf(w, "The cookie named 'username' should be gone!")
}

func main() {
	http.HandleFunc("/poke", poke) // create a cookie named "username" with value "gopher" that lasts 1 year
	http.HandleFunc("/peek", peek) // check/read cookie
	http.HandleFunc("/hide", hide) // delete cookie
	http.ListenAndServe(":8080", nil)
}
