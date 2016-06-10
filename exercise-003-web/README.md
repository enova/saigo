## Description
Write a fully functional web server!

## Introduction
Check out this code:

```go
package main

import (
  "fmt"
  "net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Just wanted to say hi!")
}

func main() {
  http.HandleFunc("/", hello)
  http.ListenAndServe(":8080", nil)
}
```

This is a simple web server that routes all requests (i.e. any URI with prefix `"/"`)
to the function (or _action_) `hello`. The function `hello` writes a plain-text
message to the response handle `w`. This code is located in the sub-directory
`exhibit-a/server.go`. Give it a try:

```
exercise-003-web$ go run exhibit-a/server.go
```
Fire up a web-browser and enter the URL `localhost:8080`.

## Comprehension Tasks
Within this exercise there are five directories:

1. `exhibit-a`: Basic Server
1. `exhibit-b`: HTML Template
1. `exhibit-c`: HTML Template With A View
1. `exhibit-d`: Parsing Forms
1. `exhibit-e`: Testing

Each directory contains code for a self-contained web-server.
Run each server from the command-line and use a browser to see them work.
Read through the implementation for each of these servers (in order) and explain what each line of code does.

## Engineering Task

Build a web-application that allows a user to enter their name.
The server should display (on the same page) a table of all users
that have entered their name along with the frequency of submissions.

This task does not require data to be persisted after shutdown. So
if your server is shut down and brought back up, it is okay if the
table is empty.

## Bonus (Optional)

Persist the list of users and their frequencies across shutdowns by
writing the requests to a file and reading that file on start-up!
