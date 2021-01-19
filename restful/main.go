package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var users = []User{
	{ID: 1, Name: "a"},
	{ID: 2, Name: "b"},
	{ID: 3, Name: "c"},
}

// User 用户
type User struct {
	ID   int
	Name string
}

func main() {
	http.HandleFunc("/users", handleUsers)
	http.ListenAndServe(":8080", nil)
}

func handleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		users, err := json.Marshal(users)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "{\"message\":\""+err.Error()+"\"}")
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write(users)
		}
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "{\"message\":\"not found\"}")
	}
}
