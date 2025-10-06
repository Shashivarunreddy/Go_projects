package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Convert map[int]User to []User
	userList := []User{}
	for _, u := range users {
		userList = append(userList, u)
	}

	json.NewEncoder(w).Encode(userList) // returns array instead of map
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var newUser User
	json.NewDecoder(r.Body).Decode(&newUser)

	newID := len(users) + 1
	newUser.ID = newID
	users[newID] = newUser // ✅ store in map

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newUser)
}

func getUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := strings.TrimPrefix(r.URL.Path, "/users/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	user, exists := users[id]
	if !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/users/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var updatedUser User
	json.NewDecoder(r.Body).Decode(&updatedUser)

	if _, exists := users[id]; !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	updatedUser.ID = id
	users[id] = updatedUser // ✅ overwrite directly

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedUser)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/users/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if _, exists := users[id]; !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	delete(users, id) // ✅ remove from map
	w.WriteHeader(http.StatusNoContent)
}
