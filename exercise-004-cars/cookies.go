package main

import (
  "net/http"
  "time"
)

func inOneYear() time.Time {
  return time.Now().AddDate(1, 0, 0)
}

func ClearUsername(w http.ResponseWriter) {
  cookie := http.Cookie{
    Name: "username",
    MaxAge: -1,
  }
  http.SetCookie(w, &cookie)
}

func GetUsername(r *http.Request) (string, error) {
  username, err := r.Cookie("username")
  if(err != nil) {
    return "", err
  }else {
    return username.Value, nil
  }
}

func SetUsername(username string, w http.ResponseWriter) bool {
  if len(username) == 0 {
    return false
  }

  cookie := http.Cookie{
    Name: "username",
    Value: username,
    Expires: inOneYear(),
  }
  http.SetCookie(w, &cookie)
  return true
}
