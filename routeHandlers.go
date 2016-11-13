package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// HOME
func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!")
}

// USERS

func CreateUser(rw http.ResponseWriter, req *http.Request) {
	body, readErr := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if readErr != nil {
		panic(readErr)
	}

	var request CreateUserRequest
	parseErr := json.Unmarshal(body, &request)
	if parseErr != nil {
		panic(parseErr)
	}

	response := CreateUserResponse{Response{http.StatusCreated}}

	rw.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(rw).Encode(response); err != nil {
		panic(err)
	}
}
