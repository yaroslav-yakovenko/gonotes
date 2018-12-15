package main

import (
	"encoding/json"
	"gonotes/gonotesserver/pkg/model"
	"math/rand"
	"net/http"
	"time"
)

func authUserHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if r.Method == "OPTIONS" {
		return
	}

	var data model.User

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	if data.Email != "dtsp@yandex.ru" || data.Password != "password" {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("bad credentials"))
		return
	}

	var session model.Session
	session.Email = data.Email
	rand.Seed(time.Now().Unix())
	session.SessionID = rand.Int63n(10000000)

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(session)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
}
