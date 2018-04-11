package main

// import fmt, net/http, time
import (
	"fmt"
	"net/http"
	"time"
)

// return a time.Time one year from now
func inOneYear() time.Time {

	// add 1 year to today's date
	return time.Now().AddDate(1, 0, 0)
}

// handler for the '/poke' path
func poke(w http.ResponseWriter, r *http.Request) {

	// create (without setting) a Cookie with the given Name and Value
	// with an expiration date of one year from the time of creation
	cookie := http.Cookie{Name: "username", Value: "gopher", Expires: inOneYear()}

	// Set the Cookie that was just created.
	http.SetCookie(w, &cookie)

	// Print the details of the cookie that was just set
	fmt.Fprintf(w, "Just set cookie named 'username' set to 'gopher'")
}

// handler for the '/peek' path
func peek(w http.ResponseWriter, r *http.Request) {

	// fetch the value for the 'username' cookie on the request
	username, err := r.Cookie("username")
	if err != nil {

		// the 'username' cookie was never set, so print the err.
		fmt.Fprintf(w, "Could not find cookie named 'username'")
		return
	}

	// print the value of the 'username' cookie
	fmt.Fprintf(w, "You have a cookie named 'username' set to '%s'", username)
}

// handler for the '/hide' path
func hide(w http.ResponseWriter, r *http.Request) {

	// create (without setting) a Cookie with the given Name
	// and a MaxAge of -1, which determines when the cookie expires.
	// This means the cookie has already expired.
	cookie := http.Cookie{Name: "username", MaxAge: -1}

	// Set the Cookie that was just created
	http.SetCookie(w, &cookie)

	// Fprint what was just performed
	fmt.Fprintf(w, "The cookie named 'username' should be gone!")
}

// main function
func main() {

	// set a handler function for '/poke'
	http.HandleFunc("/poke", poke)

	// set a handler function for '/peek'
	http.HandleFunc("/peek", peek)

	// set a handler function for '/hide'
	http.HandleFunc("/hide", hide)

	// spin up the server on Addr=":8080" with a nil Handler.
	http.ListenAndServe(":8080", nil)
}
