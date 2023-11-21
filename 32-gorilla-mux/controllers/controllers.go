package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	m "gorilla-mux/model"

	"github.com/gorilla/mux"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "home Page")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func ReturnAll(w http.ResponseWriter, r *http.Request) {
	// json.NewEncoder(w).Encode(m.Articles)

	res, _ := json.Marshal(m.Articles)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func ReturnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	i, _ := strconv.ParseInt(key, 0, 0)
	json.NewEncoder(w).Encode(m.Articles[i-1])

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
