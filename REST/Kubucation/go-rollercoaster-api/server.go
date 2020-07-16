package main

import (
	"os"
	"strings"
	"time"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"net/http"
	"sync"
)

type Coaster struct {
	Name         string `json:"name,omitempty"`
	Manufacturer string `json:"manufacturer,omitempty"`
	ID           string `json:"id,omitempty"`
	InPark       string `json:"in_park,omitempty"`
	Height       int    `json:"height,omitempty"`
}

type coasterHandlers struct {
	sync.Mutex // Just inherit
	store map[string]Coaster
}

func (h *coasterHandlers) coasters(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.get(w, r)
		return
	case "POST":
		h.post(w, r)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
		return
	}
}

func (h *coasterHandlers) post(w http.ResponseWriter, r *http.Request) {
	// We must read the request body
	bodyBytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	ct := r.Header.Get("content-type")
	if ct != "application/json" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Write([]byte(fmt.Sprintf("Need content type 'application/json', but got '%s'", ct)))
		return
	}


	// UnMarshall our body bytes (json) into a body object
	var coaster Coaster
	err = json.Unmarshal(bodyBytes, &coaster)

	if err != nil {
		// If we cant unmarshal a json the user probably sent bad data, bad request
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	coaster.ID = fmt.Sprintf("%d", time.Now().UnixNano())

	h.Lock()
	h.store[coaster.ID] = coaster
	defer h.Unlock()
}

func (h *coasterHandlers) get(w http.ResponseWriter, r *http.Request) {
	coasters := make([]Coaster, len(h.store))

	h.Lock()
	i := 0
	for _, coaster := range h.store {
		coasters[i] = coaster
		i++
	}
	h.Unlock()

	jsonBytes, err := json.Marshal(coasters)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}


	w.Header().Add("content-type","application/json") //pretty browser/postman
	w.WriteHeader(http.StatusOK)
	//curl -v localhost:8080
	w.Write(jsonBytes)
}

func (h *coasterHandlers) getCoaster(w http.ResponseWriter, r *http.Request) {
	// curl localhost:8080/coasters/foo

	// R.URL = /coasters/foo

	parts := strings.Split(r.URL.String(), "/")
	if len(parts) != 3 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	
	h.Lock()
	coaster, ok := h.store[parts[2]] 
	h.Unlock()
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}


	jsonBytes, err := json.Marshal(coaster)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}


	w.Header().Add("content-type","application/json") //pretty browser/postman
	w.WriteHeader(http.StatusOK)
	//curl -v localhost:8080
	w.Write(jsonBytes)
}

type adminPortal struct {
	password string
}

func newAdminPortal() *adminPortal {
	password := os.Getenv("ADMIN_PASSWORD")
	if password == "" {
		panic("required env var ADMIN_PASSWORD not set")
	}
	return &adminPortal{password: password}
}

func (a adminPortal) handler (w http.ResponseWriter, r *http.Request) {
	user, pass, ok := r.BasicAuth()
	if !ok || user != "admin" || pass != a.password {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("401 - unauthorized"))
		return 
	}
	w.Write([]byte("<html><h1>Super secret admin portal</h1><html>"))
}

func newCoasterHandlers() *coasterHandlers {
	return &coasterHandlers{
		store: map[string]Coaster{
			"id1": Coaster{
				Name: "Fury 325",
				Height: 99,
				ID: "id1",
				InPark: "Carowinds",
				Manufacturer: "B+M",
			},
		},
	}
}

func main (){
	admin := newAdminPortal()
	coasterHandlers := newCoasterHandlers()
	http.HandleFunc("/coasters", coasterHandlers.coasters)
	http.HandleFunc("/coasters/", coasterHandlers.getCoaster)
	http.HandleFunc("/admin/", admin.handler)
	http.ListenAndServe("localhost:8080", nil)
}