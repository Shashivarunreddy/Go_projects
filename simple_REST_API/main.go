package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			getUsers(w, r)
		case "POST":
			createUser(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			getUserByID(w, r)
		case "PUT":
			updateUser(w, r)
		case "DELETE":
			deleteUser(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.ListenAndServe(":8080", nil)
}
