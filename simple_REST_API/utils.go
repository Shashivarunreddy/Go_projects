package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func respondJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(status)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func respondError(w http.ResponseWriter, status int, message string) {
	respondJSON(w, status, map[string]string{"erroe": message})
}

func parseIDFromPath(r *http.Request) (int, error) {
	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	idStr := parts[len(parts)-1]
	return strconv.Atoi(idStr)
}
