package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type UserName struct {
	Fname string `json:fname`
	Lname string `json:lname`
}

var user = UserName{}

func reqHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST":
		{
			err := json.NewDecoder(r.Body).Decode(&user)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(user)

		}
	case "GET":
		{
			w.Header().Set("Content-Type", "Application/json")
			err := json.NewEncoder(w).Encode(user)
			if err != nil {
				fmt.Println(err)
			}
		}
	}

}

func main() {
	fmt.Println("Server has started")
	http.HandleFunc("/home", reqHandler)
	http.ListenAndServe("127.0.0.1:8080", nil)
}
