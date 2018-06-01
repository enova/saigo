## Description
You have been paired up with a UI engineer to build out a racing-simulation web app.
The UI is done, but the back-end is non-existent. Your job is to build it!

This exercise will teach you how to manage cookies and sessions.

## Comprehension Task (Cookies)

Under the directory `exhibit-a` you will find the implementation of a server
that manipulates cookies. Explain what each line of code does.

Also fire up a browser and try hitting the different routes. See what happens
when you hit `peek` before `poke`.

```
exhibit-a/server.go
-------------------

package main

import (
	"fmt"
	"net/http"
	"time"
)

func inOneYear() time.Time {
	return time.Now().AddDate(1, 0, 0)
}

func poke(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{Name: "username", Value: "gopher", Expires: inOneYear()}
	http.SetCookie(w, &cookie)
	fmt.Fprintf(w, "Just set cookie named 'username' set to 'gopher'")
}

func peek(w http.ResponseWriter, r *http.Request) {
	username, err := r.Cookie("username")
	if err != nil {
		fmt.Fprintf(w, "Could not find cookie named 'username'")
		return
	}
	fmt.Fprintf(w, "You have a cookie named 'username' set to '%s'", username)
}

func hide(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{Name: "username", MaxAge: -1}
	http.SetCookie(w, &cookie)
	fmt.Fprintf(w, "The cookie named 'username' should be gone!")
}

func main() {
	http.HandleFunc("/poke", poke)
	http.HandleFunc("/peek", peek)
	http.HandleFunc("/hide", hide)
	http.ListenAndServe(":8080", nil)
}
```

## Engineering Task

In the `templates` directory you will find two views `join.html` and `play.html`.
The `join` view asks the user to enter a _username_. Your application should store
this value in a cookie named `username` in the user's browser. The user should then
be directed to `play` where they can add vehicles to their simulation.

On the simulation page (i.e. `play.html`) the user can add different types of vehicles
to their racing simulation. 

One of the challenges here is that your application will have to keep track of
the user's list of vehicles. Each time the user clicks `add`, the application must
add the vehicle to the list and then present that same list back to the view.

The template `play.html` assumes that the following data parameter `View` is presented:

```
type Vehicle struct {
	Name  string
	Count int
}

type Vehicles struct {
	List []*Vehicle
}

type View struct {
	Username string
	Vehicles Vehicles
}
```

Your server should therefore make sure an instance of `View` is populated and then
executed by the template.

As in the previous exercise, this task does not require data to be persisted
after shutdown. So if your server is shut down and brought back up, it is okay if the
table is empty.
